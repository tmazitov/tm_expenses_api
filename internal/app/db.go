package app

import (
	domainExpense "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type DB interface {
	ExpenseRepo() domainExpense.Repository
}
