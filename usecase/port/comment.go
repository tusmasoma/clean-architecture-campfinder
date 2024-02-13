package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/entity"
)

type CommentInputPort interface {
	GetCommentBySpotID(ctx context.Context, spotID string)
	CreateComment(spotID uuid.UUID, userID uuid.UUID, starRate float64, text string)
}

type CommentOutputPort interface {
	Render()
	RenderError(error)
	RenderWithJson(interface{})
}

type CommentRepository interface {
	GetCommentBySpotID(ctx context.Context, spotID string, opts ...QueryOptions) (comments []entity.Comment, err error)
	GetCommentByID(ctx context.Context, id string, opts ...QueryOptions) (comment *entity.Comment, err error)
	Create(ctx context.Context, comment *entity.Comment, opts ...QueryOptions) (err error)
	Update(ctx context.Context, comment entity.Comment, opts ...QueryOptions) (err error)
	Delete(ctx context.Context, id string, opts ...QueryOptions) (err error)
}
