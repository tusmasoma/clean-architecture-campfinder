package port

import (
	"context"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
)

type UserInputPort interface {
	CreateUser(ctx context.Context, email string, passward string) (string, error)
}

type UserOutputPort interface {
	Render()
	RenderWithToken(string)
	RenderError(error)
}

type UserRepository interface {
	CheckIfUserExists(ctx context.Context, email string, opts ...QueryOptions) (bool, error)
	GetUserByID(ctx context.Context, id string, opts ...QueryOptions) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string, opts ...QueryOptions) (*entity.User, error)
	Create(ctx context.Context, user *entity.User, opts ...QueryOptions) error
	Update(ctx context.Context, user entity.User, opts ...QueryOptions) error
}
