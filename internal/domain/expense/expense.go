package expense

import "github.com/google/uuid"

type Expense struct {
	id   string
	name string
}

func NewExpense(id uuid.UUID, name string) (*Expense, error) {

	if len(name) == 0 {
		return nil, ErrEmptyName
	}

	if len(name) >= 256 {
		return nil, ErrNameTooLong
	}

	return &Expense{
		id:   id.String(),
		name: name,
	}, nil
}

func (e *Expense) Id() string   { return e.id }
func (e *Expense) Name() string { return e.name }
