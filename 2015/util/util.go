package util

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func Contains(list []string, valueToFind string) bool {
	for _, valueFromList := range list {
		if valueFromList == valueToFind {
			return true
		}
	}

	return false
}

// Perm calls f with each permutation of a.
func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
