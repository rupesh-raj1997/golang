package main

func getLast[T any](s []T) T {
	size := len(s)
	if size == 0 {
		return *new(T)
	}
	return s[size-1]
}
