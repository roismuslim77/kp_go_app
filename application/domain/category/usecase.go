package category

import (
	"context"
	"inventory-simple-go/application/entity"
	"inventory-simple-go/pkg/response"
	"net/http"
	"time"
)

type Repository interface {
	GetAllCategory(ctx context.Context, filter FilterListing) ([]entity.Category, int64, float64, error)
	CreateCategory(ctx context.Context, req entity.Category) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) service {
	return service{
		repository: repo,
	}
}

func (s service) CreateCategory(ctx context.Context, name string) response.ErrorResponse {
	category := entity.Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repository.CreateCategory(ctx, category)
	if err != nil {
		return *response.Error("22101").WithStatusCode(http.StatusBadRequest).WithError(err.Error())
	}

	return *response.NotError()
}

func (s service) GetAllCategory(ctx context.Context, filter FilterListing) ([]entity.Category, PaginateListing, response.ErrorResponse) {
	paginate := PaginateListing{
		TotalData: 0,
		TotalPage: 1,
	}

	category, rows, page, err := s.repository.GetAllCategory(ctx, filter)
	if err != nil {
		return []entity.Category{}, paginate, *response.Error("22101").WithStatusCode(http.StatusBadRequest).WithError(err.Error())
	}

	paginate.TotalData = rows
	paginate.TotalPage = page
	return category, paginate, *response.NotError()
}
