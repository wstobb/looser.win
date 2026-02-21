package tools

func MakeSlice[T any](n int, fn func() T) []T {
	s := make([]T, n)
	for i := range n {
		s[i] = fn()
	}
	return s
}

func Pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}
