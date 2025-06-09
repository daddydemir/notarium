package entries

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http"
)

type Handler struct {
	service *Service
	r       *mux.Router
}

func NewHandler(service *Service, r *mux.Router) *Handler {
	return &Handler{service: service, r: r}
}

func (h *Handler) RegisterRoutes() {
	h.r.HandleFunc("/api/v1/entries/{date}", h.GetByDate).Methods("GET")
}

func (h *Handler) GetByDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	entry, err := h.service.GetByDate(r.Context(), date)
	if err != nil {
		http.Error(w, "Failed to fetch entry", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entry)
}
