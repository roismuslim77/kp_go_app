package items

import (
	"context"
	"inventory-simple-go/application/entity"
	"inventory-simple-go/pkg/response"
	"net/http"
	"time"
)

type Repository interface {
	GetAllItem(ctx context.Context, filter FilterListing) ([]entity.Item, int64, float64, error)
	GetItemByID(ctx context.Context, id int) (entity.Item, error)
	CreateItem(ctx context.Context, req entity.Item) error
	UpdateItem(ctx context.Context, req entity.Item, id int) (entity.Item, error)
	DeleteItem(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) service {
	return service{
		repository: repo,
	}
}

func (s service) UpdateItem(ctx context.Context, req CreateDataRequest, id int) (entity.Item, response.ErrorResponse) {
	item := entity.Item{
		CategoryId:  req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
		Price:       &req.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	item, err := s.repository.UpdateItem(ctx, item, id)
	if err != nil {
		resp := response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
		return entity.Item{}, *resp
	}

	return item, *response.NotError()
}

func (s service) DeleteItem(ctx context.Context, id int) response.ErrorResponse {
	err := s.repository.DeleteItem(ctx, id)
	if err != nil {
		resp := response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
		return *resp
	}

	return *response.NotError()
}

func (s service) CreateItem(ctx context.Context, req CreateDataRequest) response.ErrorResponse {
	item := entity.Item{
		CategoryId:  req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
		Price:       &req.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.repository.CreateItem(ctx, item)
	if err != nil {
		return *response.Error("22101").WithStatusCode(http.StatusBadRequest).WithError(err.Error())
	}

	return *response.NotError()
}

func (s service) GetOneItem(ctx context.Context, id int) (entity.Item, response.ErrorResponse) {
	item, err := s.repository.GetItemByID(ctx, id)
	if err != nil {
		resp := response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
		return entity.Item{}, *resp
	}

	return item, *response.NotError()
}

func (s service) GetAllItem(ctx context.Context, filter FilterListing) ([]entity.Item, PaginateListing, response.ErrorResponse) {
	paginate := PaginateListing{TotalPage: 1, TotalData: 0}

	category, rows, totalPage, err := s.repository.GetAllItem(ctx, filter)
	if err != nil {
		return []entity.Item{}, paginate, *response.Error("22101").WithStatusCode(http.StatusBadRequest).WithError(err.Error())
	}

	paginate.TotalData = rows
	paginate.TotalPage = totalPage
	return category, paginate, *response.NotError()
}
