package get_pack_sizes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type packSizesGetterMock []uint64

func (p packSizesGetterMock) GetPackSizes() []uint64 {
	return p
}

func TestHandler_ServeHTTP_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	resp := httptest.NewRecorder()

	mock := packSizesGetterMock([]uint64{100, 200})
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("StatusCode = %v, want %v", result.StatusCode, http.StatusMethodNotAllowed)
	}
}

func TestHandler_ServeHTTP_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	mock := packSizesGetterMock([]uint64{100, 200})
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %v, want %v", result.StatusCode, http.StatusOK)
	}
}
