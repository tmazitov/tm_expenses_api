package utils

func TooLargeString(size int) string {
	bytes := make([]byte, size)

	for i := range size {
		bytes[i] = 'a'
	}

	return string(bytes)
}
