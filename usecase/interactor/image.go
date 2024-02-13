package interactor

import (
	"context"
	"log"

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
