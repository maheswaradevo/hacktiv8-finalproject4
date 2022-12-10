package dto

import "github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"

type CreateCategoriesRequest struct {
	Type string `json:"type" validate:"required"`
}

func (dto *CreateCategoriesRequest) ToCategoriesEntity() (ctg *model.Category) {
	ctg = &model.Category{
		Type: dto.Type,
	}
	return
}

type EditCategoriesRequest struct {
	Type string `json:"type" validate:"required"`
}

func (dto *EditCategoriesRequest) ToCategoriesEntity() (ctg *model.Category) {
	ctg = &model.Category{
		Type: dto.Type,
	}
	return
}
