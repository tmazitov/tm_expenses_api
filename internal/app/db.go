package app

import (
	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

type DB interface {
	ExpenseRepo() expense.Repository
	CategoryRepo() category.Repository
	UserRepo() user.Repository
}
