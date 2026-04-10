package category

type Category struct {
	id    string
	name  string
	icon  string
	color *Color
}

type CategoryParams struct {
	Id    string
	Name  string
	Icon  string
	Color *Color
}

func (p CategoryParams) validate() error {
	if len(p.Id) == 0 {
		return ErrIdEmpty
	}

	if len(p.Name) == 0 {
		return ErrNameEmpty
	}

	if len(p.Name) > 100 {
		return ErrNameTooLong
	}

	return nil
}

func NewCategory(params CategoryParams) (*Category, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &Category{
		id:    params.Id,
		name:  params.Name,
		icon:  params.Icon,
		color: params.Color,
	}, nil
}

func (c Category) Id() string    { return c.id }
func (c Category) Name() string  { return c.name }
func (c Category) Icon() string  { return c.icon }
func (c Category) Color() *Color { return c.color }
