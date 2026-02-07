package products

import (
	"log/slog"
	"net/http"

	"github.com/nikhiln-code/ecom-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler{
	return &handler{
		service : service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		slog.Error("Got error while fetching the products on the service layer", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.service.ListProducts(r.Context())
    json.Write(w, http.StatusOK, nil)
}
