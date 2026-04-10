package app

import (
	domainCategory "github.com/tmazitov/ayda-order-service.git/internal/domain/category"
	domainExpense "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type DB interface {
	ExpenseRepo() domainExpense.Repository
	CategoryRepo() domainCategory.Repository
}
