package slice

func Map[T any, R any](f func(T) R, s []T) []R {
	var r []R
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
