package packer

import (
	"math"
	"sort"
)

func numberOfPacks(items uint64, packSizes []uint64) map[uint64]uint64 {
	if len(packSizes) == 0 {
		return nil
	}
	if items == 0 {
		return nil
	}

	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	packSizeToNumberOfPacks := make(map[uint64]uint64, len(packSizes))
	for items != 0 {
		minExcess := uint64(math.MaxUint64)
		packSizeIndexWithMinExcess := 0
		for i := 0; i < len(packSizes); i++ {
			mod := items % packSizes[i]
			if mod == 0 {
				minExcess = mod
				packSizeIndexWithMinExcess = i
				break
			}

			excess := packSizes[i] - mod
			if excess < minExcess {
				minExcess = excess
				packSizeIndexWithMinExcess = i
			}
		}

		packSizeWithMinExcess := packSizes[packSizeIndexWithMinExcess]
		if packSizeIndexWithMinExcess == 0 && minExcess == 0 {
			packSizeToNumberOfPacks[packSizeWithMinExcess] += items / packSizeWithMinExcess
			break
		}

		packSizeToNumberOfPacks[packSizeWithMinExcess]++
		items -= packSizeWithMinExcess - minExcess
	}

	return packSizeToNumberOfPacks
}
