package util

func GetPtr[T any](s T) *T {
	return &s
}
