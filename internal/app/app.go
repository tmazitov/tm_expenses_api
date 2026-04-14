package app

import (
	"github.com/tmazitov/ayda-order-service.git/internal/app/category"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
)

type App struct {
	expenseService  *expense.Service
	categoryService *category.Service
}

func NewApp(db DB) *App {
	return &App{
		expenseService:  expense.NewService(db.ExpenseRepo()),
		categoryService: category.NewService(db.CategoryRepo()),
	}
}

func (a *App) ExpenseService() *expense.Service   { return a.expenseService }
func (a *App) CategoryService() *category.Service { return a.categoryService }
