package category

import (
	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
	"github.com/uptrace/bun"
)

type CategoryModel struct {
	bun.BaseModel `bun:"table:category"`

	Id    string  `bun:"id,pk"`
	Name  string  `bun:"name,notnull"`
	Icon  *string `bun:"icon,default:null"`
	Color *uint32 `bun:"color,default:null"`
}

func NewCategoryModel(c *category.Category) *CategoryModel {

	model := CategoryModel{
		Id:   c.Id(),
		Name: c.Name(),
	}

	if c.Color() != nil {
		color := c.Color().Uint32()
		model.Color = &color
	}

	if len(c.Icon()) != 0 {
		icon := c.Icon()
		model.Icon = &icon
	}

	return &model
}
