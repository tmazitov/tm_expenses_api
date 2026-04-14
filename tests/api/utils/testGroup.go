package utils

type TestGroup[I, O any] struct {
	Name     string
	Status   int
	Expected O
	Cases    []*TestCase[I]
}

func NewTestGroup[I, O any](name string, status int) *TestGroup[I, O] {
	return &TestGroup[I, O]{
		Name:   name,
		Status: status,
	}
}

func (g *TestGroup[I, O]) Output(output O) *TestGroup[I, O] {
	g.Expected = output
	return g
}

func (g *TestGroup[I, O]) Case(name string, input I) *TestGroup[I, O] {
	g.Cases = append(g.Cases, NewTestCase(name, input))
	return g
}
