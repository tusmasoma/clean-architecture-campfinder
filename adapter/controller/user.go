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

type User struct {
	OutputFactory func(http.ResponseWriter) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	RepoFactory   func(c *sql.DB) port.UserRepository
	Conn          *sql.DB
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outputport := u.OutputFactory(w)
	repo := u.RepoFactory(u.Conn)
	inputport := u.InputFactory(outputport, repo)

	var requestBody UserCreateRequest
	if ok := isValidUserCreateRequest(r.Body, &requestBody); !ok {
		outputport.RenderError(fmt.Errorf("Invalid user create request: %d", http.StatusBadRequest))
	}
	defer r.Body.Close()

	inputport.CreateUser(ctx, requestBody.Email, requestBody.Password)
}

func isValidUserCreateRequest(body io.ReadCloser, requestBody *UserCreateRequest) bool {
	// リクエストボディのJSONを構造体にデコード
	if err := json.NewDecoder(body).Decode(requestBody); err != nil {
		log.Printf("Invalid request body: %v", err)
		return false
	}
	if requestBody.Email == "" || requestBody.Password == "" {
		log.Printf("Missing required fields: Name or Password")
		return false
	}
	return true
}
