package fp

func Map[T any, U any](s []T, mapperFunc func(T) U) []U {
	results := make([]U, len(s))

	for i := range len(s) {
		results[i] = mapperFunc(s[i])
	}

	return results
}
