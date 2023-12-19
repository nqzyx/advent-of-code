package utils

func Reverse[T any](a []T) []T {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func Insert[T any](a []T, c T, i int) []T {
	return append(a[:i], append([]T{c}, a[i:]...)...)
}