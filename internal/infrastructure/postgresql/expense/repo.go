package expense

import (
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/uptrace/bun"
)

type expenseModel struct {
	bun.BaseModel `bun:"table:expense"`

	Id   string `bun:"id,pk"`
	Name string `bun:"name,notnull"`
}

func newExpenseModel(expense *expense.Expense) *expenseModel {
	return &expenseModel{
		Id:   expense.Id(),
		Name: expense.Name(),
	}
}
