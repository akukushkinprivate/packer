package packer

import (
	"reflect"
	"sort"
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
	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPackSizes() = %v, want %v", got, want)
	}
}
