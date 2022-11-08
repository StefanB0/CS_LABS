package rsa

func compareByteSlice(sliceA, sliceB []byte) bool {
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
