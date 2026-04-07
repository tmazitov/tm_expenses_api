package app

import "github.com/tmazitov/ayda-order-service.git/internal/app/expense"

type App struct {
	expenseService *expense.Service
}

func NewApp(db DB) *App {
	return &App{
		expenseService: expense.NewService(db.ExpenseRepo()),
	}
}

func (a *App) ExpenseService() *expense.Service { return a.expenseService }
