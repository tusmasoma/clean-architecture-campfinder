package interactor

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/internal/auth"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
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
		return
	}
	if exists {
		log.Printf("User with this name already exists - status: %d", http.StatusConflict)
		u.OutputPort.RenderError(fmt.Errorf("user with this email already exists"))
		return
	}

	var user entity.User
	user.Email = email
	user.Name = auth.ExtractUsernameFromEmail(email)
	passward, err = auth.PasswordEncrypt(passward)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		u.OutputPort.RenderError(err)
		return
	}
	user.Password = passward

	if err = u.UserRepo.Create(ctx, &user); err != nil {
		log.Printf("Failed to create user: %v", err)
		u.OutputPort.RenderError(err)
		return
	}

	jwt, _ := auth.GenerateToken(user.ID.String(), user.Email)

	u.OutputPort.RenderWithToken(jwt)
}
