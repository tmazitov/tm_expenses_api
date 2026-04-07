package expense

import "github.com/google/uuid"

type Expense struct {
	id   string
	name string
}

func validationCheck(name string) error {
	if len(name) == 0 {
		return ErrEmptyName
	}

	if len(name) >= 256 {
		return ErrNameTooLong
	}

	return nil
}

func RestoreExpense(id, name string) (*Expense, error) {
	if err := validationCheck(name); err != nil {
		return nil, err
	}

	return &Expense{
		id:   id,
		name: name,
	}, nil
}

func NewExpense(name string) (*Expense, error) {

	if err := validationCheck(name); err != nil {
		return nil, err
	}

	return &Expense{
		id:   uuid.NewString(),
		name: name,
	}, nil
}

func (e *Expense) Id() string   { return e.id }
func (e *Expense) Name() string { return e.name }
