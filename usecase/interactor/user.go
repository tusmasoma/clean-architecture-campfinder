package interactor

import (
	"context"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
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
	user := entity.User{Email: email, Password: passward}
	err := u.UserRepo.Create(ctx, &user)
	if err != nil {
		u.OutputPort.RenderError(err)
	}
	u.OutputPort.Render()
}
