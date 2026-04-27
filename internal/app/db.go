package app

import (
	"github.com/tmazitov/tm_expenses_api/internal/domain/category"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
	"github.com/tmazitov/tm_expenses_api/internal/domain/user"
)

type DB interface {
	ExpenseRepo() expense.Repository
	CategoryRepo() category.Repository
	UserRepo() user.Repository
}
