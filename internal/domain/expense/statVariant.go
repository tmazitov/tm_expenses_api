package expense

type expenseStatVariant string

var (
	WeeklyStat  expenseStatVariant = "weekly"
	MonthlyStat expenseStatVariant = "monthly"
)

func NewExpenseStatVariant(value string) (expenseStatVariant, error) {
	switch v := expenseStatVariant(value); v {
	case MonthlyStat, WeeklyStat:
		return v, nil
	default:
		return expenseStatVariant(""), ErrUnknownStatVariant
	}
}
