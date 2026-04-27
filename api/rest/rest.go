package rest

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/tm_expenses_api/api/rest/category"
	"github.com/tmazitov/tm_expenses_api/api/rest/expense"
	"github.com/tmazitov/tm_expenses_api/api/rest/middleware"
	"github.com/tmazitov/tm_expenses_api/api/rest/user"
	"github.com/tmazitov/tm_expenses_api/internal/app"
)

type RestAPI struct {
	categoryRouter *category.Router
	expenseRouter  *expense.Router
	userRouter     *user.Router

	authMiddleware fiber.Handler
}

func NewRestAPI(a app.App) *RestAPI {
	return &RestAPI{
		categoryRouter: category.NewRouter(a.CategoryService()),
		expenseRouter:  expense.NewRouter(a.ExpenseService()),
		userRouter:     user.NewRouter(a.UserService()),
		authMiddleware: middleware.AuthMiddleware(a.UserService()),
	}
}

func (api *RestAPI) Register(app *fiber.App) {

	api.categoryRouter.Register(app).
		Use(api.authMiddleware).
		Routes()

	api.expenseRouter.Register(app).
		Use(api.authMiddleware).
		Routes()

	api.userRouter.Register(app).
		Use(api.authMiddleware).
		Routes()
}
