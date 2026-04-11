package expense

type ExpenseStatIdentifier uint8

func NewExpenseStatIdentifier(value uint8, variant expenseStatVariant) (ExpenseStatIdentifier, error) {

	switch variant {
	case WeeklyStat:
		if value > 52 {
			return ExpenseStatIdentifier(0), ErrInvalidWeeklyID
		}
	case MonthlyStat:
		if value > 12 {
			return ExpenseStatIdentifier(0), ErrInvalidMonthlyID
		}
	default:
		return ExpenseStatIdentifier(0), ErrUnknownStatVariant
	}

	return ExpenseStatIdentifier(value), nil
}
