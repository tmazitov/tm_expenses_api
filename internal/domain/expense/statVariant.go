package expense

type expenseStatVariant string

var (
	WeeklyStat          expenseStatVariant = "weekly"
	MonthlyStat         expenseStatVariant = "monthly"
	WeeklyCategoryStat  expenseStatVariant = "weekly-categories"
	MonthlyCategoryStat expenseStatVariant = "monthly-categories"
)

func NewExpenseStatVariant(value string) (expenseStatVariant, error) {
	switch v := expenseStatVariant(value); v {
	case MonthlyStat, WeeklyStat, WeeklyCategoryStat, MonthlyCategoryStat:
		return v, nil
	default:
		return expenseStatVariant(""), ErrUnknownStatVariant
	}
}
