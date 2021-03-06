package validator

import (
	"myecho/config"
	"myecho/handler/errors"
	"myecho/handler/rtype"
	"myecho/model"
)

func ValidateCategoryID(categoryID uint) error {
	if categoryID == 0 {
		return nil
	}
	err := config.Database.Where("id = ?", categoryID).First(&model.Category{}).Error
	if err != nil {
		return errors.ErrCategoryNotFound
	}
	return nil

}
func ValidateCategoryCreate(req *rtype.CategoryCreateRequest) error {
	if req.Name == "" {
		return errors.ErrCategoryNameEmpty
	}
	if req.FatherID != nil {
		if err := ValidateCategoryID(*req.FatherID); err != nil {
			return err
		}
	}
	return nil
}

func ValidateCategoryUpdate(req *rtype.CategoryUpdateRequest) error {
	if req.Name != nil && *req.Name == "" {
		return errors.ErrCategoryNameEmpty
	}
	if req.FatherID != nil {
		if err := ValidateCategoryID(*req.FatherID); err != nil {
			return err
		}
	}
	return nil
}
