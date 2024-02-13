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
	OutputPort  port.CommentOutputPort
	CommentRepo port.CommentRepository
	UserRepo    port.UserRepository
}

func NewCommentInputPort(outputPort port.CommentOutputPort, commentRepository port.CommentRepository, userRepository port.UserRepository) port.CommentInputPort {
	return &Comment{
		OutputPort:  outputPort,
		CommentRepo: commentRepository,
		UserRepo:    userRepository,
	}
}

type CommentGetResponse struct {
	Comments []entity.Comment `json:"comments"`
}

func (c *Comment) GetCommentBySpotID(ctx context.Context, spotID string) {
	comments, err := c.CommentRepo.GetCommentBySpotID(ctx, spotID)
	if err != nil {
		log.Printf("Failed to get comments by spot id: %v", err)
		c.OutputPort.RenderError(err)
		return
	}
	c.OutputPort.RenderWithJson(CommentGetResponse{Comments: comments})
}

func (c *Comment) CreateComment(spotID uuid.UUID, userID uuid.UUID, starRate float64, text string) {
	comment := &entity.Comment{
		SpotID:   spotID,
		UserID:   userID,
		StarRate: starRate,
		Text:     text,
	}
	if err := c.CommentRepo.Create(context.Background(), comment); err != nil {
		log.Printf("Failed to create comment: %v", err)
		c.OutputPort.RenderError(err)
		return
	}
	c.OutputPort.Render()
}

func (c *Comment) UpdateComment(
	ctx context.Context,
	id uuid.UUID,
	spotID uuid.UUID,
	userID uuid.UUID,
	starRate float64,
	text string,
	ctxUserID uuid.UUID,
) {
	user, err := c.UserRepo.GetUserByID(ctx, ctxUserID.String())
	if err != nil {
		log.Printf("Failed to get user by id: %v", err)
		c.OutputPort.RenderError(err)
		return
	}
	if !user.IsAdmin && user.ID != userID {
		log.Print("Don't have permission to update comment")
		c.OutputPort.RenderError(fmt.Errorf("don't have permission to update comment"))
		return
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
		c.OutputPort.RenderError(err)
	}
	c.OutputPort.Render()
}
