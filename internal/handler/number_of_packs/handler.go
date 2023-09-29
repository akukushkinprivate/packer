package number_of_packs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type inDto struct {
	Items uint64 `json:"items"`
}

type outDto struct {
	NumberOfPacks []string `json:"numberOfPacks"`
}

type packer interface {
	NumberOfPacks(items uint64) map[uint64]uint64
}

type Handler struct {
	packer packer
}

func New(packer packer) *Handler {
	return &Handler{packer: packer}
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

	if in.Items == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	numberOfPacks := h.packer.NumberOfPacks(in.Items)
	var out outDto
	out.NumberOfPacks = make([]string, 0, len(numberOfPacks))
	for packSize, packs := range numberOfPacks {
		out.NumberOfPacks = append(out.NumberOfPacks, fmt.Sprintf("%d x %d", packs, packSize))
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(responseWriter).Encode(out); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
