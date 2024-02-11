package controller

import (
	"database/sql"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type User struct {
	OutputFactory func(http.ResponseWriter) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	RepoFactory   func(c *sql.DB) port.UserRepository
	Conn          *sql.DB
}

func (u *User) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	email := "example@gmail.com"
	pasward := "hogehoge"
	outputport := u.OutputFactory(w)
	repo := u.RepoFactory(u.Conn)
	inputport := u.InputFactory(outputport, repo)
	inputport.CreateUser(r.Context(), email, pasward)
}
