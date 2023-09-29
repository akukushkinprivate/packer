package packer

import (
	"sync"
)

const defaultPackSizesCapacity = 10000

var defaultPackSizes = []uint64{250, 500, 1000, 2000, 5000}

type Service struct {
	packSizes []uint64
	mu        sync.RWMutex
}

func New() *Service {
	return &Service{packSizes: defaultPackSizes, mu: sync.RWMutex{}}
}

func (s *Service) GetPackSizes() []uint64 {
	packSizes := make([]uint64, 0, defaultPackSizesCapacity)

	s.mu.RLock()
	defer s.mu.RUnlock()

	packSizes = copySlice(packSizes, s.packSizes)

	return packSizes
}

func (s *Service) SetPackSizes(packSizes []uint64) {
	packSizes = removeDuplicates(packSizes)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.packSizes = s.packSizes[:0]
	s.packSizes = copySlice(s.packSizes, packSizes)
}

func (s *Service) NumberOfPacks(items uint64) map[uint64]uint64 {
	packSizes := s.GetPackSizes()
	return numberOfPacks(items, packSizes)
}

func copySlice[T any](dst, src []T) []T {
	for _, s := range src {
		dst = append(dst, s)
	}

	return dst
}

func removeDuplicates[T comparable](src []T) []T {
	unique := make(map[T]struct{}, len(src))
	for _, s := range src {
		unique[s] = struct{}{}
	}

	src = src[:0]
	for u := range unique {
		src = append(src, u)
	}

	return src
}
