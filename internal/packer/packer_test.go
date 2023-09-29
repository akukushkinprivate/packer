package packer

import (
	"reflect"
	"testing"
)

func TestService_GetPackSizes(t *testing.T) {
	service := New()

	got := service.GetPackSizes()
	want := []uint64{250, 500, 1000, 2000, 5000}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPackSizes() = %v, want %v", got, want)
	}

	service.SetPackSizes([]uint64{5})
	got = service.GetPackSizes()
	want = []uint64{5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPackSizes() = %v, want %v", got, want)
	}

	service.SetPackSizes([]uint64{5, 5, 250, 500, 1000, 2000, 5000, 5000000})
	got = service.GetPackSizes()
	want = []uint64{5, 250, 500, 1000, 2000, 5000, 5000000}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPackSizes() = %v, want %v", got, want)
	}
}
