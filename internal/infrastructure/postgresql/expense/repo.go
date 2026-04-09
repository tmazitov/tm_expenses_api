package expense

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/uptrace/bun"
)

type expenseModel struct {
	bun.BaseModel `bun:"table:expense"`

	Id        string          `bun:"id,pk"`
	Name      string          `bun:"name,notnull"`
	Price     decimal.Decimal `bun:"price,notnull"`
	CreatedAt time.Time       `bun:"created_at,notnull"`
}

func newExpenseModel(expense *expense.Expense) *expenseModel {
	return &expenseModel{
		Id:        expense.Id(),
		Name:      expense.Name(),
		CreatedAt: expense.CreatedAt(),
		Price:     expense.Price(),
	}
}
