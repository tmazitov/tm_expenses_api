package rest

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/api/rest/category"
	"github.com/tmazitov/ayda-order-service.git/api/rest/expense"
	"github.com/tmazitov/ayda-order-service.git/internal/app"
)

type RestAPI struct {
	categoryRouter *category.Router
	expenseRouter  *expense.Router
}

func NewRestAPI(a app.App) *RestAPI {
	return &RestAPI{
		categoryRouter: category.NewRouter(a.CategoryService()),
		expenseRouter:  expense.NewRouter(a.ExpenseService()),
	}
}

func (api *RestAPI) Register(app *fiber.App) {
	
	api.categoryRouter.Register(app)
	api.expenseRouter.Register(app)
}
