package postgresql

import (
	domainCategory "github.com/tmazitov/tm_expenses_api/internal/domain/category"
	domainExpense "github.com/tmazitov/tm_expenses_api/internal/domain/expense"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql/category"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql/expense"
)

type Mock struct{}

func (m Mock) ExpenseRepo() domainExpense.Repository   { return expense.Mock{} }
func (m Mock) CategoryRepo() domainCategory.Repository { return category.Mock{} }
