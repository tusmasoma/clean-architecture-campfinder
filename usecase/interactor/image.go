package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Image struct {
	OutputPort port.ImageOutputPort
	ImageRepo  port.ImageRepository
	UserRepo   port.UserRepository
}

func NewImageInputPort(outputPort port.ImageOutputPort, imageRepository port.ImageRepository, userRepository port.UserRepository) port.ImageInputPort {
	return &Image{
		OutputPort: outputPort,
		ImageRepo:  imageRepository,
		UserRepo:   userRepository,
	}
}

type ImageGetResponse struct {
	Images []entity.Image `json:"images"`
}

func (i *Image) GetSpotImgURLBySpotID(ctx context.Context, spotID string) {
	imgs, err := i.ImageRepo.GetSpotImgURLBySpotID(ctx, spotID)
	if err != nil {
		log.Print("Failed to get images by spot id")
		i.OutputPort.RenderError(err)
		return
	}
	i.OutputPort.RenderWithJson(ImageGetResponse{Images: imgs})
}

func (i *Image) CreateImage(ctx context.Context, spotID uuid.UUID, userID uuid.UUID, url string) {
	img := &entity.Image{
		SpotID: spotID,
		UserID: userID,
		URL:    url,
	}
	if err := i.ImageRepo.Create(ctx, img); err != nil {
		log.Printf("Failed to create image: %v", err)
		i.OutputPort.RenderError(err)
		return
	}
	i.OutputPort.Render()
}

func (i *Image) DeleteImage(ctx context.Context, id string, userID string, ctxUserID uuid.UUID) {
	user, err := i.UserRepo.GetUserByID(ctx, ctxUserID.String())
	if err != nil {
		log.Printf("Failed to get user by id: %v", err)
		i.OutputPort.RenderError(err)
		return
	}
	if !user.IsAdmin && user.ID.String() != userID {
		log.Print("Don't have permission to delete image")
		i.OutputPort.RenderError(fmt.Errorf("don't have permission to delete image"))
		return
	}

	if err = i.ImageRepo.Delete(ctx, id); err != nil {
		log.Printf("Failed to delete comment: %v", err)
		i.OutputPort.RenderError(err)
		return
	}
	i.OutputPort.Render()
}
