package get_pack_sizes

import (
	"encoding/json"
	"net/http"
)

type outDto struct {
	PackSizes []uint64 `json:"packSizes"`
}

type packSizesGetter interface {
	GetPackSizes() []uint64
}

type Handler struct {
	packSizesGetter packSizesGetter
}

func New(packSizesGetter packSizesGetter) *Handler {
	return &Handler{packSizesGetter: packSizesGetter}
}

func (h *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	packSizes := h.packSizesGetter.GetPackSizes()
	out := outDto{PackSizes: packSizes}
	responseWriter.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(responseWriter).Encode(out); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
