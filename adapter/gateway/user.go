package gateway

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CheckIfUserExists(
	ctx context.Context,
	email string,
	opts ...port.QueryOptions,
) (bool, error) {
	var exists bool
	var executor port.SQLExecutor = ur.db
	if len(opts) > 0 && opts[0].Executor != nil {
		executor = opts[0].Executor
	}

	query := `SELECT EXISTS(SELECT 1 FROM User WHERE email = ?)`
	err := executor.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *userRepository) GetUserByID(
	ctx context.Context,
	id string, opts ...port.QueryOptions,
) (*entity.User, error) {
	var executor port.SQLExecutor = ur.db
	if len(opts) > 0 && opts[0].Executor != nil {
		executor = opts[0].Executor
	}

	query := `
	SELECT *
	FROM User
	WHERE id = ?
	`

	var user entity.User
	err := executor.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
	)
	return &user, err
}

func (ur *userRepository) GetUserByEmail(
	ctx context.Context,
	email string,
	opts ...port.QueryOptions,
) (*entity.User, error) {
	var executor port.SQLExecutor = ur.db
	if len(opts) > 0 && opts[0].Executor != nil {
		executor = opts[0].Executor
	}

	query := `
	SELECT *
	FROM User
	WHERE email = ?
	`

	var user entity.User
	err := executor.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
	)
	return &user, err
}

func (ur *userRepository) Create(ctx context.Context, user *entity.User, opts ...port.QueryOptions) error {
	var executor port.SQLExecutor = ur.db
	if len(opts) > 0 && opts[0].Executor != nil {
		executor = opts[0].Executor
	}

	query := `
    INSERT INTO User (
        id, name, email, password
    )
    VALUES (?, ?, ?, ?)
    `

	user.ID = uuid.New()

	_, err := executor.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)

	return err
}

func (ur *userRepository) Update(ctx context.Context, user entity.User, opts ...port.QueryOptions) error {
	var executor port.SQLExecutor = ur.db
	if len(opts) > 0 && opts[0].Executor != nil {
		executor = opts[0].Executor
	}

	query := `
	UPDATE User SET
	name=?, email=?, password=?
	WHERE id = ?
	`
	_, err := executor.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID)
	return err
}
