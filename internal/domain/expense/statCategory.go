package expense

type ExpenseStatCategory struct {
	id      string
	percent int
}

func (s ExpenseStatCategory) Id() string   { return s.id }
func (s ExpenseStatCategory) Percent() int { return s.percent }

func NewExpenseStatCategory(id string, percent int) (ExpenseStatCategory, error) {
	if percent < 0 || percent > 100 {
		return ExpenseStatCategory{}, ErrInvalidPercent
	}
	// if len(id) == 0 {
	// 	return ExpenseStatCategory{}, ErrInvalidId
	// }

	return ExpenseStatCategory{
		id:      id,
		percent: percent,
	}, nil
}
