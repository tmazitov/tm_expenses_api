package expense

import "github.com/shopspring/decimal"

type ExpenseStat struct {
	key        ExpenseStatIdentifier
	value      decimal.Decimal
	categories []ExpenseStatCategory
}

func NewExpenseStat(key uint8, value decimal.Decimal, variant expenseStatVariant) (*ExpenseStat, error) {

	k, err := NewExpenseStatIdentifier(key, variant)
	if err != nil {
		return nil, err
	}

	return &ExpenseStat{
		key:   k,
		value: value,
	}, nil
}

func (s *ExpenseStat) WithCategories(categories []ExpenseStatCategory) *ExpenseStat {
	return &ExpenseStat{
		key:        s.key,
		value:      s.value,
		categories: categories,
	}
}

func (s *ExpenseStat) Key() uint8                        { return uint8(s.key) }
func (s *ExpenseStat) Value() decimal.Decimal            { return s.value }
func (s *ExpenseStat) Categories() []ExpenseStatCategory { return s.categories }
