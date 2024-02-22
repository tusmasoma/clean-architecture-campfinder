package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/entity"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Spot struct {
	OutputFactory func(http.ResponseWriter) port.SpotOutputPort
	InputFactory  func(u port.SpotRepository) port.SpotInputPort
	RepoFactory   func(c *sql.DB) port.SpotRepository
	Conn          *sql.DB
}

type SpotCreateRequest struct {
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Period      string  `json:"period"`
	Phone       string  `json:"phone"`
	Price       string  `json:"price"`
	Description string  `json:"description"`
	IconPath    string  `json:"iconpath"`
}

type GetResponse struct {
	Spots []entity.Spot `json:"spots"`
}

func (s *Spot) HandleSpotCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := s.OutputFactory(w)
	repo := s.RepoFactory(s.Conn)
	inputport := s.InputFactory(repo)

	var requestBody SpotCreateRequest
	if ok := isValidateSpotCreateRequest(r.Body, &requestBody); !ok {
		outputport.RenderError(fmt.Errorf("Invalid spot create request: %d", http.StatusBadRequest))
		return
	}

	if err := inputport.CreateSpot(
		ctx,
		requestBody.Category,
		requestBody.Name,
		requestBody.Address,
		requestBody.Lat,
		requestBody.Lng,
		requestBody.Period,
		requestBody.Phone,
		requestBody.Price,
		requestBody.Description,
		requestBody.IconPath,
	); err != nil {
		outputport.RenderError(err)
		return
	}

	outputport.Render()
}

func isValidateSpotCreateRequest(body io.ReadCloser, requestBody *SpotCreateRequest) bool {
	if err := json.NewDecoder(body).Decode(requestBody); err != nil {
		log.Printf("Invalid request body: %v", err)
		return false
	}
	if requestBody.Category == "" ||
		requestBody.Name == "" ||
		requestBody.Address == "" ||
		requestBody.Lat == 0 ||
		requestBody.Lng == 0 {
		log.Printf("Missing required fields")
		return false
	}
	return true
}

func (s *Spot) HandleSpotGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := s.OutputFactory(w)
	repo := s.RepoFactory(s.Conn)
	inputport := s.InputFactory(repo)

	categories := r.URL.Query()["category"]
	spotID := r.URL.Query().Get("spot_id")

	spots := inputport.GetSpot(ctx, categories, spotID)

	outputport.RenderWithJson(GetResponse{Spots: spots})
}
