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

	userIDValue := ctx.Value(config.ContextUserIDKey)
	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		log.Printf("Failed to retrieve userId from context")
		outputport.RenderError(fmt.Errorf("user name not found in request context"))
		return
	}

	var requestBody CommentCreateRequest
	if ok := isValidateCommentCreateRequest(r.Body, &requestBody); !ok {
		log.Printf("Invalid comment create request")
		outputport.RenderError(fmt.Errorf("Invalid comment create request: %v", http.StatusBadRequest))
		return
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

func (c *Comment) HandleCommentUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := c.OutputFactory(w)
	repo := c.RepoFactory(c.Conn)
	inputport := c.InputFactory(outputport, repo, nil)

	ctxUserIDValue := ctx.Value(config.ContextUserIDKey)
	ctxUserID, ok := ctxUserIDValue.(uuid.UUID)
	if !ok {
		log.Printf("Failed to retrieve userId from context")
		outputport.RenderError(fmt.Errorf("user name not found in request context"))
		return
	}

	var requestBody CommentUpdateRequest
	if ok := isValidateCommentUpdateRequest(r.Body, &requestBody); !ok {
		log.Printf("Invalid comment update request")
		outputport.RenderError(fmt.Errorf("Invalid comment update request: %v", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()

	inputport.UpdateComment(ctx, requestBody.ID, requestBody.SpotID, requestBody.UserID, requestBody.StarRate, requestBody.Text, ctxUserID)
}

func isValidateCommentUpdateRequest(body io.ReadCloser, requestBody *CommentUpdateRequest) bool {
	if err := json.NewDecoder(body).Decode(requestBody); err != nil {
		log.Printf("Invalid request body: %v", err)
		return false
	}
	if requestBody.ID.String() == config.DefaultUUID ||
		requestBody.SpotID.String() == config.DefaultUUID ||
		requestBody.UserID.String() == config.DefaultUUID ||
		requestBody.StarRate == 0 ||
		requestBody.Text == "" {
		log.Printf("Missing required fields")
		return false
	}
	return true
}
