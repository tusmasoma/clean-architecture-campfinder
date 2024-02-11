package driver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/adapter/controller"
	"github.com/tusmasoma/clean-architecture-campfinder/adapter/gateway"
	"github.com/tusmasoma/clean-architecture-campfinder/adapter/presenter"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/interactor"
)

func Serve(addr string) {
	dsn := "hoge"

	db, _ := sql.Open("mysql", dsn)

	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          db,
	}

	http.HandleFunc("/api/user/create", user.HandleUserCreate)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
