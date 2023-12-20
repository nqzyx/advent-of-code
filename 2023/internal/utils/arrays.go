package utils

func Reverse[T any](a []T) []T {
	newA := make([]T, 0, len(a))
	for i, j := 0, len(a)-1; i < len(a)-1; i, j = i+1, j-1 {
		newA[i] = a[j]
	}
	return newA
}

func Insert[T any](a []T, c T, i int) []T {
	return append(a[:i], append([]T{c}, a[i:]...)...)
}