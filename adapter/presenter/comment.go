package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Comment struct {
	w http.ResponseWriter
}

func NewCommentOutputPort(w http.ResponseWriter) port.CommentOutputPort {
	return &Spot{
		w: w,
	}
}

func (s *Comment) Render() {
	s.w.WriteHeader(http.StatusOK)
}

func (s *Comment) RenderError(err error) {
	s.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(s.w, err)
}

func (s *Comment) RenderWithJson(response interface{}) {
	s.w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(s.w).Encode(response); err != nil {
		s.RenderError(err)
	}
}
