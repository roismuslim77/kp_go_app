package category

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-simple-go/application/entity"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateCategory(ctx context.Context, req entity.Category) error {
	if err := r.db.Create(&req).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			err = errors.New("category is already exist")
			return err
		}
		return err
	}

	return nil
}

func (r repository) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	var data []entity.Category
	if err := r.db.Order("created_at DESC").Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
