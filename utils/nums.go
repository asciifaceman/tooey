package utils

// SliceIntMax returns the highest integer within a []int
//
// an empty slice returns 0
func SliceIntMax(slice []int) int {
	m := 0
	for _, val := range slice {
		if val > m {
			m = val
		}
	}
	return m
}

// SliceIntIndexMin returns the index of a []int with the lowest integer value
func SliceIntIndexMin(slice []int) int {
	li := 0
	lc := SliceIntMax(slice)

	if lc == 0 {
		// figure out what to do here to tighten it up
		// I mean it works until we find a situation it doesn't
		return 0
	}

	for i, count := range slice {
		if count < lc {
			lc = count
			li = i
		}
	}

	return li
}
