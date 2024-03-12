package utils

// Merge unique elements in two slices
func MergeSlice[T comparable](l1 []T, l2 []T) []T {
	var hashTarget []T
	var other []T
	if len(l1) > len(l2) {
		hashTarget = l2
		other = l1
	} else {
		hashTarget = l1
		other = l2
	}

	if len(l1) == 0 || len(l2) == 0 {
		return other
	}

	hm := make(map[T]bool, len(hashTarget))
	for _, t := range hashTarget {
		hm[t] = true
	}

	res := hashTarget
	for _, t := range other {
		if _, ok := hm[t]; !ok {
			res = append(res, t)
		}
	}

	return res
}
