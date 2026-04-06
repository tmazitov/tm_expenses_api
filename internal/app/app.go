package app

import "github.com/tmazitov/ayda-order-service.git/internal/app/expense"

type App struct {
	expenseService *expense.Service
	db             DB
}

func NewApp(db DB) *App {
	return &App{
		expenseService: expense.NewService(db.ExpenseRepo()),
		db:             db,
	}
}

func (a *App) ExpenseService() *expense.Service { return a.expenseService }
