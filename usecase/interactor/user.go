package interactor

import (
	"context"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/internal/auth"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type User struct {
	UserRepo port.UserRepository
}

func NewUserInputPort(userRepository port.UserRepository) port.UserInputPort {
	return &User{
		UserRepo: userRepository,
	}
}

func (u *User) CreateUser(ctx context.Context, email string, passward string) (string, error) {
	exists, err := u.UserRepo.CheckIfUserExists(ctx, email)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		return "", err
	}
	if exists {
		log.Printf("User with this name already exists - status: %d", http.StatusConflict)
		return "", err
	}

	var user entity.User
	user.Email = email
	user.Name = auth.ExtractUsernameFromEmail(email)
	passward, err = auth.PasswordEncrypt(passward)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		return "", err
	}
	user.Password = passward

	if err = u.UserRepo.Create(ctx, &user); err != nil {
		log.Printf("Failed to create user: %v", err)
		return "", err
	}

	jwt, _ := auth.GenerateToken(user.ID.String(), user.Email)

	return jwt, nil
}
