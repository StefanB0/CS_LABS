package rsa

func compareU64Slice(sliceA, sliceB []uint64) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i, element := range sliceA {
		if element != sliceB[i] {
			return false
		}
	}
	return true
}
