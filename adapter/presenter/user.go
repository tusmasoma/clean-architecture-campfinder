package presenter

import (
	"fmt"
	"net/http"

	"github.com/tusmasoma/clean-architecture-campfinder/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

func (u *User) Render() {
	u.w.WriteHeader(http.StatusOK)
}

func (u *User) RenderWithToken(jwt string) {
	u.w.Header().Set("Authorization", "Bearer "+jwt)
	u.w.WriteHeader(http.StatusOK)
}

func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
