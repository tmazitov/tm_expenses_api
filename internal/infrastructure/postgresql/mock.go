package postgresql

import (
	domainCategory "github.com/tmazitov/ayda-order-service.git/internal/domain/category"
	domainExpense "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql/category"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql/expense"
)

type Mock struct{}

func (m Mock) ExpenseRepo() domainExpense.Repository   { return expense.Mock{} }
func (m Mock) CategoryRepo() domainCategory.Repository { return category.Mock{} }
