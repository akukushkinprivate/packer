package set_pack_sizes

import (
	"encoding/json"
	"net/http"
)

type inDto struct {
	PackSizes []uint64 `json:"packSizes"`
}

type packSizesSetter interface {
	SetPackSizes(packSizes []uint64)
}

type Handler struct {
	packSizesSetter packSizesSetter
}

func New(packSizesSetter packSizesSetter) *Handler {
	return &Handler{packSizesSetter: packSizesSetter}
}

func (h *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var in inDto
	if err := json.NewDecoder(request.Body).Decode(&in); err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(in.PackSizes) == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	h.packSizesSetter.SetPackSizes(in.PackSizes)
	responseWriter.WriteHeader(http.StatusOK)
}
