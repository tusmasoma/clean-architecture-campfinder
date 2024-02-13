package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Spot struct {
	w http.ResponseWriter
}

func NewSpotOutputPort(w http.ResponseWriter) port.SpotOutputPort {
	return &Spot{
		w: w,
	}
}

func (s *Spot) Render() {
	s.w.WriteHeader(http.StatusOK)
}

func (s *Spot) RenderError(err error) {
	s.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(s.w, err)
}

func (s *Spot) RenderWithJson(response interface{}) {
	s.w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(s.w).Encode(response); err != nil {
		s.RenderError(err)
	}
}
