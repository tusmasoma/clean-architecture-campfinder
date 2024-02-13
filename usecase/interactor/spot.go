package interactor

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Spot struct {
	OutputPort port.SpotOutputPort
	SpotRepo   port.SpotRepository
}

func NewSpotInputPort(outputPort port.SpotOutputPort, spotRepository port.SpotRepository) port.SpotInputPort {
	return &Spot{
		OutputPort: outputPort,
		SpotRepo:   spotRepository,
	}
}

func (s *Spot) CreateSpot(
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
) {
	exists, err := s.SpotRepo.CheckIfSpotExists(ctx, lat, lng)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		s.OutputPort.RenderError(err)
	}
	if !exists {
		log.Printf("Spot with this name already exists - status: %d", http.StatusConflict)
		s.OutputPort.RenderError(fmt.Errorf("user with this name already exists"))
	}

	if err = s.SpotRepo.Create(ctx, &entity.Spot{
		Category:    category,
		Name:        name,
		Address:     address,
		Lat:         lat,
		Lng:         lng,
		Period:      period,
		Phone:       phone,
		Price:       price,
		Description: description,
		IconPath:    iconpath,
	}); err != nil {
		log.Printf("Failed to create spot: %v", err)
		s.OutputPort.RenderError(err)
	}

	s.OutputPort.Render()
}