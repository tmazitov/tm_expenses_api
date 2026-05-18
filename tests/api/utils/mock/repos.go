package mock

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/tmazitov/tm_expenses_api/internal/domain/category"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
	"github.com/tmazitov/tm_expenses_api/internal/domain/user"
)

// CATEGORY REPO

type mockCategoryRepo struct{}

func (m mockCategoryRepo) Create(ctx context.Context, c *category.Category) error { return nil }
func (m mockCategoryRepo) List(ctx context.Context) ([]*category.Category, error) {
	return []*category.Category{}, nil
}

// EXPENSE REPO

type mockExpenseRepo struct{}

func (m mockExpenseRepo) Create(ctx context.Context, p *expense.Expense) error { return nil }
func (m mockExpenseRepo) GetById(ctx context.Context, id string) (*expense.Expense, error) {
	return &expense.Expense{}, nil
}
func (m mockExpenseRepo) List(ctx context.Context, filters expense.ListFilters) ([]*expense.Expense, error) {

	instance, _ := expense.NewExpense(expense.ExpenseParams{
		Id:         uuid.NewString(),
		UserId:     1,
		Name:       "expense",
		Price:      decimal.New(1, 10),
		CreatedAt:  time.Now(),
		CategoryId: uuid.NewString(),
	})

	return []*expense.Expense{instance}, nil
}
func (m mockExpenseRepo) WeeklyCategoryStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}
func (m mockExpenseRepo) WeeklyStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}
func (m mockExpenseRepo) MonthlyStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}
func (m mockExpenseRepo) MonthlyCategoryStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}

// USER REPO

type mockUserRepo struct{}

func (m mockUserRepo) GetBySub(ctx context.Context, method user.AuthMethod, sub string) (*user.User, error) {
	return &user.User{}, nil
}
func (m mockUserRepo) GetById(ctx context.Context, id int) (*user.User, error) {
	return &user.User{}, nil
}
func (m mockUserRepo) Create(ctx context.Context, u *user.User) (*user.User, error) {
	return &user.User{}, nil
}
func (m mockUserRepo) Update(ctx context.Context, u *user.User) error {
	return nil
}
