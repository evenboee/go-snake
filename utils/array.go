package utils

func Shift[T any](arr []T) ([]T, T) {
	var first T

	if len(arr) == 0 {
		return arr, first
	}

	first = arr[0]

	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr = arr[:len(arr)-1]

	return arr, first
}

func ShiftIn[T any](arr []T, el T) ([]T, T) {
	arr, first := Shift(arr)
	arr = append(arr, el)

	return arr, first
}

func Any[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}
