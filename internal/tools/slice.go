package tools

func MakeSlice[T any](n int, fn func() T) []T {
	s := make([]T, n)
	for i := range n {
		s[i] = fn()
	}
	return s
}
