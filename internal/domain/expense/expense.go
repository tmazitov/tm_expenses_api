package expense

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	id         string
	name       string
	categoryId string
	price      decimal.Decimal
	createdAt  time.Time
}

type ExpenseParams struct {
	Id         string
	Name       string
	CategoryId string
	Price      decimal.Decimal
	CreatedAt  time.Time
}

func (p ExpenseParams) validate() error {

	if p.Price.Cmp(decimal.NewFromInt(0)) != 1 {
		return ErrPriceIsNegative
	}

	if len(p.Name) == 0 {
		return ErrEmptyName
	}

	if len(p.Name) >= 256 {
		return ErrNameTooLong
	}

	return nil
}

// Creates new instance of Expense.
func NewExpense(params ExpenseParams) (*Expense, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &Expense{
		id:         params.Id,
		name:       params.Name,
		price:      params.Price,
		categoryId: params.CategoryId,
		createdAt:  params.CreatedAt,
	}, nil
}

func (e Expense) Id() string             { return e.id }
func (e Expense) Name() string           { return e.name }
func (e Expense) CategoryId() string     { return e.categoryId }
func (e Expense) Price() decimal.Decimal { return e.price }
func (e Expense) CreatedAt() time.Time   { return e.createdAt }
