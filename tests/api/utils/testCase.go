package utils

type TestCase[I any] struct {
	Name  string
	Input I
}

func NewTestCase[T any](name string, input T) *TestCase[T] {

	return &TestCase[T]{
		Name:  name,
		Input: input,
	}
}
