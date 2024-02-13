package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/tusmasoma/clean-architecture-campfinder/config"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Comment struct {
	OutputFactory   func(http.ResponseWriter) port.CommentOutputPort
	InputFactory    func(o port.CommentOutputPort, c port.CommentRepository, u port.UserRepository) port.CommentInputPort
	RepoFactory     func(c *sql.DB) port.CommentRepository
	UserRepoFactory func(c *sql.DB) port.UserRepository
	Conn            *sql.DB
}

type CommentCreateRequest struct {
	SpotID   uuid.UUID `json:"spotID"`
	StarRate float64   `json:"starRate"`
	Text     string    `json:"text"`
}

type CommentUpdateRequest struct {
	ID       uuid.UUID `json:"id"`
	SpotID   uuid.UUID `json:"spotID"`
	UserID   uuid.UUID `json:"userID"`
	StarRate float64   `json:"starRate"`
	Text     string    `json:"text"`
}

func (c *Comment) HandleCommentGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spotID := r.URL.Query().Get("spot_id")

	outputport := c.OutputFactory(w)
	repo := c.RepoFactory(c.Conn)
	inputport := c.InputFactory(outputport, repo, nil)

	inputport.GetCommentBySpotID(ctx, spotID)
}

func (c *Comment) HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := c.OutputFactory(w)
	repo := c.RepoFactory(c.Conn)
	inputport := c.InputFactory(outputport, repo, nil)

	var ok bool
	userID := ctx.Value(config.ContextUserIDKey)
	userID, ok = userID.(uuid.UUID)
	if !ok {
		log.Printf("Failed to retrieve userId from context")
		outputport.RenderError(fmt.Errorf("user name not found in request context"))
	}

	var requestBody CommentCreateRequest
	if ok := isValidateCommentCreateRequest(r.Body, &requestBody); !ok {
		log.Printf("Invalid comment create request")
		outputport.RenderError(fmt.Errorf("Invalid comment create request: %v", http.StatusBadRequest))
	}
	defer r.Body.Close()

	inputport.CreateComment(requestBody.SpotID, userID, requestBody.StarRate, requestBody.Text)
}

func isValidateCommentCreateRequest(body io.ReadCloser, requestBody *CommentCreateRequest) bool {
	if err := json.NewDecoder(body).Decode(requestBody); err != nil {
		log.Printf("Invalid request body: %v", err)
		return false
	}
	if requestBody.SpotID.String() == config.DefaultUUID || requestBody.StarRate == 0 || requestBody.Text == "" {
		log.Printf("Missing required fields")
		return false
	}
	return true
}
