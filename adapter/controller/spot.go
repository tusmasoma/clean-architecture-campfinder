package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Spot struct {
	OutputFactory func(http.ResponseWriter) port.SpotOutputPort
	InputFactory  func(o port.SpotOutputPort, u port.SpotRepository) port.SpotInputPort
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

func (s *Spot) HandleSpotCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := s.OutputFactory(w)
	repo := s.RepoFactory(s.Conn)
	inputport := s.InputFactory(outputport, repo)

	var requestBody SpotCreateRequest
	if ok := isValidateSpotCreateRequest(r.Body, &requestBody); !ok {
		outputport.RenderError(fmt.Errorf("Invalid spot create request: %d", http.StatusBadRequest))
	}

	inputport.CreateSpot(
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
	)
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
