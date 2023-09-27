package slice

func Contains[T any](source []T, target T, compare func(a T, b T) bool) bool {
	for _, item := range source {
		if compare(item, target) {
			return true
		}
	}
	return false
}
