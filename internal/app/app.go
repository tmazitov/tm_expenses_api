package app

import (
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/app/category"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
	"github.com/tmazitov/ayda-order-service.git/internal/app/user"
)

type App struct {
	expenseService  *expense.Service
	categoryService *category.Service
	userService     *user.Service
}

type Infrastructure struct {
	DB          DB
	Cache       Cache
	GoogleOAuth GoogleOAuthProvider
	Jwt         JwtProvider
}

func (i *Infrastructure) validate() error {
	if i.DB == nil || i.Cache == nil || i.GoogleOAuth == nil || i.Jwt == nil {
		return errors.New("app error : not enough components in infrastructure to set up app.")
	}
	return nil
}

func NewApp(infra Infrastructure) (*App, error) {

	if err := infra.validate(); err != nil {
		return nil, err
	}

	userService, err := user.NewService(user.ServiceParams{
		Jwt:                 infra.Jwt,
		GoogleOAuthProvider: infra.GoogleOAuth,
		Repo:                infra.DB.UserRepo(),
	})
	if err != nil {
		return nil, err
	}

	return &App{
		expenseService:  expense.NewService(infra.DB.ExpenseRepo()),
		categoryService: category.NewService(infra.DB.CategoryRepo()),
		userService:     userService,
	}, nil
}

func (a *App) UserService() *user.Service         { return a.userService }
func (a *App) ExpenseService() *expense.Service   { return a.expenseService }
func (a *App) CategoryService() *category.Service { return a.categoryService }
