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
	return &Comment{
		w: w,
	}
}

func (c *Comment) Render() {
	c.w.WriteHeader(http.StatusOK)
}

func (c *Comment) RenderError(err error) {
	c.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(c.w, err)
}

func (c *Comment) RenderWithJson(response interface{}) {
	c.w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(c.w).Encode(response); err != nil {
		c.RenderError(err)
	}
}
