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

type Image struct {
	OutputFactory   func(http.ResponseWriter) port.ImageOutputPort
	InputFactory    func(o port.ImageOutputPort, i port.ImageRepository) port.ImageInputPort
	RepoFactory     func(c *sql.DB) port.ImageRepository
	UserRepoFactory func(c *sql.DB) port.UserRepository
	Conn            *sql.DB
}

type ImageCreateRequest struct {
	SpotID uuid.UUID `json:"spotID"`
	URL    string    `json:"url"`
}

func (i *Image) HandleImageGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spotID := r.URL.Query().Get("spot_id")

	outputport := i.OutputFactory(w)
	repo := i.RepoFactory(i.Conn)
	inputport := i.InputFactory(outputport, repo)

	inputport.GetSpotImgURLBySpotID(ctx, spotID)
}

func (i *Image) HandleImageCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := i.OutputFactory(w)
	repo := i.RepoFactory(i.Conn)
	inputport := i.InputFactory(outputport, repo)

	userIDValue := ctx.Value(config.ContextUserIDKey)
	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		log.Print("Failed to retrieve userId from context")
		outputport.RenderError(fmt.Errorf("user name not found in request context"))
		return
	}

	var requestBody ImageCreateRequest
	if ok := isValidateImageCreateRequest(r.Body, &requestBody); !ok {
		log.Print("Invalid image create request")
		outputport.RenderError(fmt.Errorf("Invalid image create request: %v", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()

	inputport.CreateImage(ctx, requestBody.SpotID, userID, requestBody.URL)
}

func isValidateImageCreateRequest(body io.ReadCloser, requestBody *ImageCreateRequest) bool {
	if err := json.NewDecoder(body).Decode(requestBody); err != nil {
		log.Printf("Invalid request body: %v", err)
		return false
	}
	if requestBody.SpotID.String() == config.DefaultUUID || requestBody.URL == "" {
		log.Printf("Missing required fields")
		return false
	}
	return true
}

func (i *Image) HandleImageDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := i.OutputFactory(w)
	repo := i.RepoFactory(i.Conn)
	inputport := i.InputFactory(outputport, repo)

	ctxUserIDValue := ctx.Value(config.ContextUserIDKey)
	ctxUserID, ok := ctxUserIDValue.(uuid.UUID)
	if !ok {
		log.Print("Failed to retrieve userId from context")
		outputport.RenderError(fmt.Errorf("user name not found in request context"))
		return
	}

	ok, id, userID := isValidateImageDeleteRequest(r)
	if !ok {
		log.Print("Invalid image delete request")
		outputport.RenderError(fmt.Errorf("Invalid image delete request"))
		return
	}

	inputport.DeleteImage(ctx, id, userID, ctxUserID)
}

func isValidateImageDeleteRequest(r *http.Request) (bool, string, string) {
	id := r.URL.Query().Get("id")
	userID := r.URL.Query().Get("user_id")

	if id == "" || userID == "" {
		log.Printf("Missing required fields")
		return false, "", ""
	}
	return true, id, userID
}
