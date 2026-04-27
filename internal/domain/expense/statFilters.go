package expense

type ExpenseStatFilters struct {
	variant expenseStatVariant
	units   uint8
	page    int
	userId  int
}

type ExpenseStatFiltersParams struct {
	Variant expenseStatVariant
	Units   uint8
	UserId  int
	Page    int
}

func (p *ExpenseStatFiltersParams) validate() error {

	if p.Units > 6 {
		return ErrUnitsIsNegative
	}

	if p.Page < 0 {
		return ErrPageIsNegative
	}

	if p.UserId == 0 {
		return ErrUserIdZero
	}

	return nil
}

func NewExpenseStatFilters(params ExpenseStatFiltersParams) (*ExpenseStatFilters, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &ExpenseStatFilters{
		variant: params.Variant,
		units:   params.Units,
		page:    params.Page,
		userId:  params.UserId,
	}, nil
}

func (f *ExpenseStatFilters) Variant() expenseStatVariant { return f.variant }
func (f *ExpenseStatFilters) Units() uint8                { return f.units }
func (f *ExpenseStatFilters) Page() int                   { return f.page }
func (f *ExpenseStatFilters) UserId() int                 { return f.userId }
