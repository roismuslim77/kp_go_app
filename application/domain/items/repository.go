package items

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"inventory-simple-go/application/entity"
	"log"
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

func (r repository) UpdateItem(ctx context.Context, req entity.Item, id int) (entity.Item, error) {
	result := r.db.Clauses(&clause.Returning{}).Where("id = ?", id).Updates(&req)
	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			result.Error = errors.New("item is already exist")
			return req, result.Error
		}
		return req, result.Error
	}

	if result.RowsAffected < 1 {
		return req, errors.New("failed to update item")
	}

	return req, nil
}

func (r repository) DeleteItem(ctx context.Context, id int) error {
	result := r.db.Where("id = ?", id).Delete(&entity.Item{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("failed to delete item")
	}

	return nil
}

func (r repository) CreateItem(ctx context.Context, req entity.Item) error {
	if err := r.db.Create(&req).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			err = errors.New("items is already exist")
			return err
		}
		return err
	}

	return nil
}

func (r repository) GetItemByID(ctx context.Context, id int) (entity.Item, error) {
	var data entity.Item
	result := r.db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		if result.RowsAffected < 1 {
			data.IsEmpty = true
			return data, nil
		}
		return data, result.Error
	}

	return data, nil
}

func (r repository) GetAllItem(ctx context.Context, filter FilterListing) ([]entity.Item, int64, float64, error) {
	var data []entity.Item
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

	if err := r.db.Table("items").Offset(offset).Limit(size).Order("created_at DESC").Find(&data).Error; err != nil {
		return data, 0, 1, err
	}

	r.db.Table("items").Count(&totalData)
	if totalData > 0 {
		totalPage = float64(totalData) / float64(size)
		totalPage = math.Ceil(totalPage)
	} else {
		totalPage = 1
	}

	log.Println(totalPage)
	log.Println(totalData)
	return data, totalData, totalPage, nil
}
