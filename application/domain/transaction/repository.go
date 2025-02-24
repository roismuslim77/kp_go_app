package transaction

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"simple-go/application/entity"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) GetCustomerTenor(ctx context.Context, customerId, tenor int) (entity.CustomerLimit, error) {
	var data entity.CustomerLimit
	result := r.db.Where("customer_id = ? and tenor = ?", customerId, tenor).First(&data)
	if result.Error != nil {
		if result.RowsAffected < 1 {
			data.IsEmpty = true
			return data, nil
		}
		return data, result.Error
	}

	return data, nil
}
func (r repository) UpdateCustomerTenor(ctx context.Context, req entity.CustomerLimit, id int) (entity.CustomerLimit, error) {
	log.Println(req)
	result := r.db.Clauses(&clause.Returning{}).Where("id = ?", id).Updates(&req)
	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			result.Error = errors.New("limit is already exist")
			return req, result.Error
		}
		return req, result.Error
	}

	if result.RowsAffected < 1 {
		return req, errors.New("failed to update limit")
	}

	return req, nil
}

func (r repository) CreateCustomerTransaction(ctx context.Context, req entity.Transaction) (entity.Transaction, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(&req).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			err = errors.New("transaction is already exist")
			return req, err
		}
		return req, err
	}

	return req, nil
}
