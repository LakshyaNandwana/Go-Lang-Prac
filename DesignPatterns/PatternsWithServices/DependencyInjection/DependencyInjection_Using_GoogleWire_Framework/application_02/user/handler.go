package user

import (
	"DI-Wire/domain"
	"net/http"
)

type handler struct {
	svc domain.UserService
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	panic("implement me! :D")
}
