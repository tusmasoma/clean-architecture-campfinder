package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Image struct {
	w http.ResponseWriter
}

func NewImageOutputPort(w http.ResponseWriter) port.ImageOutputPort {
	return &Image{
		w: w,
	}
}

func (i *Image) Render() {
	i.w.WriteHeader(http.StatusOK)
}

func (i *Image) RenderError(err error) {
	i.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(i.w, err)
}

func (i *Image) RenderWithJson(response interface{}) {
	i.w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(i.w).Encode(response); err != nil {
		i.RenderError(err)
	}
}
