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

	mode establishMode
}

func NewRestAPI(a app.App, mode establishMode) *RestAPI {
	return &RestAPI{
		mode: mode,

		categoryRouter: category.NewRouter(a.CategoryService()),
		expenseRouter:  expense.NewRouter(a.ExpenseService()),
		userRouter:     user.NewRouter(a.UserService()),
		authMiddleware: middleware.AuthMiddleware(a.UserService(), mode == UnsafeMode),
	}
}

type Router interface {
	Register(a *fiber.App)
	Use(middleware fiber.Handler)
	ApplyRoutes()
}

func (api *RestAPI) Register(app *fiber.App) {

	routers := []Router{
		api.userRouter,
		api.expenseRouter,
		api.categoryRouter,
	}

	for _, router := range routers {
		router.Register(app)
		router.Use(api.authMiddleware)
		router.ApplyRoutes()
	}
}
