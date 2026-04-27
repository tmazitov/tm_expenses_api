package expense

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql/category"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql/user"
	"github.com/uptrace/bun"
)

type expenseModel struct {
	bun.BaseModel `bun:"table:expense"`

	Id         string          `bun:"id,pk"`
	Name       string          `bun:"name,notnull"`
	UserId     int             `bun:"user_id,notnull"`
	CategoryId string          `bun:"category_id,default:null"`
	Price      decimal.Decimal `bun:"price,notnull"`
	CreatedAt  time.Time       `bun:"created_at,notnull"`

	Category *category.CategoryModel `bun:"rel:belongs-to,join:category_id=id"`
	User     *user.UserModel         `bun:"rel:belongs-to,join:user_id=id"`
}

func newExpenseModel(expense *expense.Expense) *expenseModel {
	return &expenseModel{
		Id:         expense.Id(),
		Name:       expense.Name(),
		CreatedAt:  expense.CreatedAt(),
		Price:      expense.Price(),
		CategoryId: expense.CategoryId(),
		UserId:     expense.UserId(),
	}
}

func (m *expenseModel) toExpenseParams() expense.ExpenseParams {
	return expense.ExpenseParams{
		Id:         m.Id,
		Name:       m.Name,
		Price:      m.Price,
		UserId:     m.UserId,
		CreatedAt:  m.CreatedAt,
		CategoryId: m.CategoryId,
	}
}
