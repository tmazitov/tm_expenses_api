package expense

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	id        string
	name      string
	price     decimal.Decimal
	createdAt time.Time
}

func validationCheck(name string, price decimal.Decimal) error {

	if price.Cmp(decimal.NewFromInt(0)) != 1 {
		return ErrPriceIsNegative
	}

	if len(name) == 0 {
		return ErrEmptyName
	}

	if len(name) >= 256 {
		return ErrNameTooLong
	}

	return nil
}

// Creates new instance of Expense.
func NewExpense(id, name string, price decimal.Decimal) (*Expense, error) {

	if err := validationCheck(name, price); err != nil {
		return nil, err
	}

	return &Expense{
		id:        id,
		name:      name,
		price:     price,
		createdAt: time.Now(),
	}, nil
}

// Allows to restore previously saved instance of Expense.
func RestoreExpense(id, name string, price decimal.Decimal, createdAt time.Time) (*Expense, error) {
	if err := validationCheck(name, price); err != nil {
		return nil, err
	}

	return &Expense{
		id:        id,
		name:      name,
		price:     price,
		createdAt: createdAt,
	}, nil
}

func (e *Expense) Id() string             { return e.id }
func (e *Expense) Name() string           { return e.name }
func (e *Expense) Price() decimal.Decimal { return e.price }
func (e *Expense) CreatedAt() time.Time   { return e.createdAt }
