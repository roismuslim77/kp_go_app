package category

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-simple-go/application/entity"
	"math"
	"strconv"
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

func (r repository) GetAllCategory(ctx context.Context, filter FilterListing) ([]entity.Category, int64, float64, error) {
	var data []entity.Category
	var totalData int64
	var totalPage float64

	//pagination
	offset := 0
	size := 10
	if filter.Size != "" {
		size, _ = strconv.Atoi(filter.Size)
	}
	if filter.Page != "" {
		page, _ := strconv.Atoi(filter.Page)
		offset = (page - 1) * size
	}

	if err := r.db.Table("categories").Offset(offset).Limit(size).Order("created_at DESC").Find(&data).Error; err != nil {
		return data, 0, 1, err
	}

	r.db.Table("categories").Count(&totalData)
	if totalData > 0 {
		totalPage = float64(totalData) / float64(size)
		totalPage = math.Ceil(totalPage)
	} else {
		totalPage = 1
	}

	return data, totalData, totalPage, nil
}
