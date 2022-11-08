package rsa

func compareI64Slice(sliceA, sliceB []int64) bool {
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
