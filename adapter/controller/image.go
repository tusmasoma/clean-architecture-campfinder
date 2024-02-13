package controller

import (
	"database/sql"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type Image struct {
	OutputFactory   func(http.ResponseWriter) port.ImageOutputPort
	InputFactory    func(o port.ImageOutputPort, i port.ImageRepository) port.ImageInputPort
	RepoFactory     func(c *sql.DB) port.ImageRepository
	UserRepoFactory func(c *sql.DB) port.UserRepository
	Conn            *sql.DB
}

func (i *Image) HandleImageGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spotID := r.URL.Query().Get("spot_id")

	outputport := i.OutputFactory(w)
	repo := i.RepoFactory(i.Conn)
	inputport := i.InputFactory(outputport, repo)

	inputport.GetSpotImgURLBySpotID(ctx, spotID)
}
