package port

import (
	"context"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
)

type ImageInputPort interface {
	GetSpotImgURLBySpotID(ctx context.Context, spotID string)
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
