package controller

import (
	"database/sql"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Comment struct {
	OutputFactory func(http.ResponseWriter) port.CommentOutputPort
	InputFactory  func(o port.CommentOutputPort, u port.CommentRepository) port.CommentInputPort
	RepoFactory   func(c *sql.DB) port.CommentRepository
	Conn          *sql.DB
}

func (c *Comment) HandleCommentGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spotID := r.URL.Query().Get("spot_id")

	outputport := c.OutputFactory(w)
	repo := c.RepoFactory(c.Conn)
	inputport := c.InputFactory(outputport, repo)

	inputport.GetCommentBySpotID(ctx, spotID)
}
