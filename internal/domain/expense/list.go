package expense

import "time"

type ListFilters struct {
	name  string
	page  int
	limit int
	date  time.Time
}
type ListFiltersParams struct {
	Name  string
	Date  time.Time
	Limit int
	Page  int
}

func (p *ListFiltersParams) validate() error {
	if p.Page < 0 {
		return ErrPageIsNegative
	}
	if p.Limit < 0 {
		return ErrLimitIsNegative
	}
	if p.Date.IsZero() {
		return ErrDateIsZero
	}

	return nil
}

func NewListFilters(params ListFiltersParams) (*ListFilters, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &ListFilters{
		name:  params.Name,
		limit: params.Limit,
		page:  params.Page,
		date:  params.Date,
	}, nil
}

func (f *ListFilters) Name() string { return f.name }
func (f *ListFilters) Date() string { return f.date.Format("02.01.2006") }
func (f *ListFilters) Limit() int   { return f.limit }
func (f *ListFilters) Page() int    { return f.page }
