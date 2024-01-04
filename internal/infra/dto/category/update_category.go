package categorydto

import (
	productentity "github.com/willjrcom/sales-backend-go/internal/domain/product"
)

var ()

type UpdateCategoryInput struct {
	productentity.PatchCategory
}

func (c *UpdateCategoryInput) validate() error {
	if c.Name == nil {
		return ErrNameIsEmpty
	}

	return nil
}

func (c *UpdateCategoryInput) UpdateModel(category *productentity.Category) (err error) {
	if err = c.validate(); err != nil {
		return err
	}

	if c.Name != nil {
		category.Name = *c.Name
	}

	return nil
}