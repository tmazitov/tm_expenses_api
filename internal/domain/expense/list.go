package expense

type ListFilters struct {
	name  string
	page  int
	limit int
}

func validateListFilters(page, limit int) error {
	if page < 0 {
		return ErrPageIsNegative
	}
	if limit < 0 {
		return ErrLimitIsNegative
	}
	return nil
}

func NewListFilters(name string, page, limit int) (*ListFilters, error) {

	if err := validateListFilters(page, limit); err != nil {
		return nil, err
	}

	return &ListFilters{
		name:  name,
		limit: limit,
		page:  page,
	}, nil
}

func (f *ListFilters) Name() string { return f.name }
func (f *ListFilters) Limit() int   { return f.limit }
func (f *ListFilters) Page() int    { return f.page }
