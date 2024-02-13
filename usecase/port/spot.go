package port

import (
	"context"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
)

type SpotInputPort interface {
	CreateSpot(
		ctx context.Context,
		category string,
		name string,
		address string,
		lat float64,
		lng float64,
		period string,
		phone string,
		price string,
		description string,
		iconpath string,
	)
}

type SpotOutputPort interface {
	Render()
	RenderError(error)
	RenderWithJson(interface{})
}

type SpotRepository interface {
	CheckIfSpotExists(ctx context.Context, lat float64, lng float64, opts ...QueryOptions) (bool, error)
	GetSpotByID(ctx context.Context, id string, opts ...QueryOptions) (*entity.Spot, error)
	GetSpotByCategory(ctx context.Context, category string, opts ...QueryOptions) (spots []entity.Spot, err error)
	Create(ctx context.Context, spot *entity.Spot, opts ...QueryOptions) (err error)
	Update(ctx context.Context, spot entity.Spot, opts ...QueryOptions) (err error)
	Delete(ctx context.Context, id string, opts ...QueryOptions) (err error)
}
