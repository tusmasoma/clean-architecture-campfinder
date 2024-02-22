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
	ImageRepo port.ImageRepository
	UserRepo  port.UserRepository
}

func NewImageInputPort(imageRepository port.ImageRepository, userRepository port.UserRepository) port.ImageInputPort {
	return &Image{
		ImageRepo: imageRepository,
		UserRepo:  userRepository,
	}
}

func (i *Image) GetSpotImgURLBySpotID(ctx context.Context, spotID string) []entity.Image {
	imgs, err := i.ImageRepo.GetSpotImgURLBySpotID(ctx, spotID)
	if err != nil {
		log.Print("Failed to get images by spot id")
		return nil
	}
	return imgs
}

func (i *Image) CreateImage(ctx context.Context, spotID uuid.UUID, userID uuid.UUID, url string) error {
	img := &entity.Image{
		SpotID: spotID,
		UserID: userID,
		URL:    url,
	}
	if err := i.ImageRepo.Create(ctx, img); err != nil {
		log.Printf("Failed to create image: %v", err)
		return err
	}
	return nil
}

func (i *Image) DeleteImage(ctx context.Context, id string, userID string, ctxUserID uuid.UUID) error {
	user, err := i.UserRepo.GetUserByID(ctx, ctxUserID.String())
	if err != nil {
		log.Printf("Failed to get user by id: %v", err)
		return err
	}
	if !user.IsAdmin && user.ID.String() != userID {
		log.Print("Don't have permission to delete image")
		return fmt.Errorf("don't have permission to delete image")
	}

	if err = i.ImageRepo.Delete(ctx, id); err != nil {
		log.Printf("Failed to delete comment: %v", err)
		return err
	}
	return nil
}
