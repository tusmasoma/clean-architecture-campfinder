package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/entity"
)

type ImageInputPort interface {
	GetSpotImgURLBySpotID(ctx context.Context, spotID string) []entity.Image
	CreateImage(ctx context.Context, spotID uuid.UUID, userID uuid.UUID, url string) error
	DeleteImage(ctx context.Context, id string, userID string, ctxUserID uuid.UUID) error
}

type ImageOutputPort interface {
	Render()
	RenderError(error)
	RenderWithJson(interface{})
}

type ImageRepository interface {
	GetSpotImgURLBySpotID(ctx context.Context, spotID string, opts ...QueryOptions) (imgs []entity.Image, err error)
	Create(ctx context.Context, img *entity.Image, opts ...QueryOptions) (err error)
	Delete(ctx context.Context, id string, opts ...QueryOptions) (err error)
}
