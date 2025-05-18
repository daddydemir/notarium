package v1

import (
	"encoding/json"
	"net/http"

	"github.com/daddydemir/notarium/internal/services"
	"github.com/gorilla/mux"
)

type Handler[T any] struct {
	service services.Service[T]
}

func NewHandler[T any](service services.Service[T]) *Handler[T] {
	return &Handler[T]{service}
}

func (h *Handler[T]) List(w http.ResponseWriter, r *http.Request) {
	entities, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch entities", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entities)
}

func (h *Handler[T]) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	entity, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to fetch entity", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}

func (h *Handler[T]) Create(w http.ResponseWriter, r *http.Request) {
	var entity T
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.Create(r.Context(), entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler[T]) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var entity T
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.Update(r.Context(), id, entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler[T]) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete entity", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
