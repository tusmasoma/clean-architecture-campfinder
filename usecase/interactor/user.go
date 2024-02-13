package interactor

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/internal/auth"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

func (u *User) CreateUser(ctx context.Context, email string, passward string) {
	exists, err := u.UserRepo.CheckIfUserExists(ctx, email)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		u.OutputPort.RenderError(err)
	}
	if exists {
		log.Printf("User with this name already exists - status: %d", http.StatusConflict)
		u.OutputPort.RenderError(fmt.Errorf("user with this email already exists"))
	}

	var user entity.User
	user.Email = email
	user.Name = ExtractUsernameFromEmail(email)
	passward, err = PasswordEncrypt(passward)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		u.OutputPort.RenderError(err)
	}
	user.Password = passward

	if err = u.UserRepo.Create(ctx, &user); err != nil {
		log.Printf("Failed to create user: %v", err)
		u.OutputPort.RenderError(err)
	}

	jwt, _ := auth.GenerateToken(user.ID.String(), user.Email)

	u.OutputPort.RenderWithToken(jwt)
}

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ExtractUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
