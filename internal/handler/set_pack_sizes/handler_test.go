package set_pack_sizes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type packSizesSetterMock struct{}

func (p packSizesSetterMock) SetPackSizes(_ []uint64) {
}

func TestHandler_ServeHTTP_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	mock := packSizesSetterMock{}
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("StatusCode = %v, want %v", result.StatusCode, http.StatusMethodNotAllowed)
	}
}

func TestHandler_ServeHTTP_Empty_Request_Body(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	resp := httptest.NewRecorder()

	mock := packSizesSetterMock{}
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode = %v, want %v", result.StatusCode, http.StatusBadRequest)
	}
}

func TestHandler_ServeHTTP_Invalid_Request_Body(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"packSizes\":[]}"))
	resp := httptest.NewRecorder()

	mock := packSizesSetterMock{}
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode = %v, want %v", result.StatusCode, http.StatusBadRequest)
	}
}

func TestHandler_ServeHTTP_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"items\":1}"))
	resp := httptest.NewRecorder()

	mock := packSizesSetterMock{}
	handler := New(mock)
	handler.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != http.StatusOK {
		t.Fatalf("StatusCode = %v, want %v", result.StatusCode, http.StatusOK)
	}
	all, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(all)
	want := "{\"numberOfPacks\":[\"1 x 100\",\"3 x 200\"]}\n"
	wantAlt := "{\"numberOfPacks\":[\"3 x 200\",\"1 x 100\"]}\n"
	if got != want && got != wantAlt {
		t.Errorf("Body = %v, want %v", got, want)
	}
}
