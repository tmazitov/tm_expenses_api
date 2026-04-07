package expense

import "time"

type Expense struct {
	id        string
	name      string
	createdAt time.Time
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

// Creates new instance of Expense.
func NewExpense(id, name string) (*Expense, error) {

	if err := validationCheck(name); err != nil {
		return nil, err
	}

	return &Expense{
		id:        id,
		name:      name,
		createdAt: time.Now(),
	}, nil
}

// Allows to restore previously saved instance of Expense.
func RestoreExpense(id, name string, createdAt time.Time) (*Expense, error) {
	if err := validationCheck(name); err != nil {
		return nil, err
	}

	return &Expense{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}, nil
}

func (e *Expense) Id() string           { return e.id }
func (e *Expense) Name() string         { return e.name }
func (e *Expense) CreatedAt() time.Time { return e.createdAt }
