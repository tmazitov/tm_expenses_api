package expense

type ExpenseStat struct {
	key   ExpenseStatIdentifier
	value int
}

func NewExpenseStat(key uint8, value int, variant expenseStatVariant) (*ExpenseStat, error) {

	k, err := NewExpenseStatIdentifier(key, variant)
	if err != nil {
		return nil, err
	}

	return &ExpenseStat{
		key:   k,
		value: value,
	}, nil
}

func (s *ExpenseStat) Key() uint8 { return uint8(s.key) }
func (s *ExpenseStat) Value() int { return s.value }
