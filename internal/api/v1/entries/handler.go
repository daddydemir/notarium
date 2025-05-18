package entries

import (
	_ "encoding/json"
	_ "net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}
