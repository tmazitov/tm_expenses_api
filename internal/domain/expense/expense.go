package expense

import "github.com/google/uuid"

type Expense struct {
	id   string `validate:"required,min=1"`
	name string `validate:"required,min=1"`
}

func NewExpense(name string) (*Expense, error) {

	if len(name) == 0 {
		return nil, ErrInvalidExpense
	}

	return &Expense{
		id:   uuid.NewString(),
		name: name,
	}, nil
}

func (e *Expense) Id() string   { return e.id }
func (e *Expense) Name() string { return e.name }
