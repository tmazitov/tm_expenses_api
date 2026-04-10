package category

import "context"

type CategoryListItem struct {
	Id    string
	Name  string
	Icon  string
	Color string
}

func (s *Service) List(ctx context.Context) ([]CategoryListItem, error) {

	categories, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	list := make([]CategoryListItem, 0, len(categories))
	for _, c := range categories {

		item := CategoryListItem{
			Id:   c.Id(),
			Name: c.Name(),
			Icon: c.Icon(),
		}

		color := c.Color()
		if color != nil {
			item.Color = color.Hex()
		}

		list = append(list, item)
	}
	return list, nil
}
