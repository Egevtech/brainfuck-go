package util

func ForEach[T any](i []T, lb func(int, T)) {
	for index, curr := range i {
		lb(index, curr)
	}
}
