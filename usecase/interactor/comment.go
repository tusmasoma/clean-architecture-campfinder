package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Comment struct {
	CommentRepo port.CommentRepository
	UserRepo    port.UserRepository
}

func NewCommentInputPort(commentRepository port.CommentRepository, userRepository port.UserRepository) port.CommentInputPort {
	return &Comment{
		CommentRepo: commentRepository,
		UserRepo:    userRepository,
	}
}

func (c *Comment) GetCommentBySpotID(ctx context.Context, spotID string) []entity.Comment {
	comments, err := c.CommentRepo.GetCommentBySpotID(ctx, spotID)
	if err != nil {
		log.Printf("Failed to get comments by spot id: %v", err)
		return nil
	}
	return comments
}

func (c *Comment) CreateComment(ctx context.Context, spotID uuid.UUID, userID uuid.UUID, starRate float64, text string) error {
	comment := &entity.Comment{
		SpotID:   spotID,
		UserID:   userID,
		StarRate: starRate,
		Text:     text,
	}
	if err := c.CommentRepo.Create(ctx, comment); err != nil {
		log.Printf("Failed to create comment: %v", err)
		return err
	}
	return nil
}

func (c *Comment) UpdateComment(
	ctx context.Context,
	id uuid.UUID,
	spotID uuid.UUID,
	userID uuid.UUID,
	starRate float64,
	text string,
	ctxUserID uuid.UUID,
) error {
	user, err := c.UserRepo.GetUserByID(ctx, ctxUserID.String())
	if err != nil {
		log.Printf("Failed to get user by id: %v", err)
		return err
	}
	if !user.IsAdmin && user.ID != userID {
		log.Print("Don't have permission to update comment")
		return fmt.Errorf("don't have permission to update comment")
	}

	comment := entity.Comment{
		ID:       id,
		SpotID:   spotID,
		UserID:   userID,
		StarRate: starRate,
		Text:     text,
	}

	if err := c.CommentRepo.Update(ctx, comment); err != nil {
		log.Printf("Failed to update comment: %v", err)
		return err
	}
	return nil
}

func (c *Comment) DeleteComment(ctx context.Context, id string, userID string, ctxUserID uuid.UUID) error {
	user, err := c.UserRepo.GetUserByID(ctx, ctxUserID.String())
	if err != nil {
		log.Printf("Failed to get user by id: %v", err)
		return err
	}
	if !user.IsAdmin && user.ID.String() != userID {
		log.Print("Don't have permission to delete comment")
		return fmt.Errorf("don't have permission to delete comment")
	}

	if err = c.CommentRepo.Delete(ctx, id); err != nil {
		log.Printf("Failed to delete comment: %v", err)
		return err
	}
	return nil
}
