package interactor

import (
	"context"
	"log"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Comment struct {
	OutputPort  port.CommentOutputPort
	CommentRepo port.CommentRepository
}

func NewCommentInputPort(outputPort port.CommentOutputPort, commentRepository port.CommentRepository) port.CommentInputPort {
	return &Comment{
		OutputPort:  outputPort,
		CommentRepo: commentRepository,
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
	}
	c.OutputPort.RenderWithJson(CommentGetResponse{Comments: comments})
}
