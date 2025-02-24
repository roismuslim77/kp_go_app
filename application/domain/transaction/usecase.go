package transaction

import (
	"context"
	"math"
	"net/http"
	"simple-go/application/entity"
	"simple-go/pkg/response"
	"strconv"
	"time"
)

type Repository interface {
	GetCustomerTenor(ctx context.Context, customerId, tenor int) (entity.CustomerLimit, error)
	UpdateCustomerTenor(ctx context.Context, req entity.CustomerLimit, id int) (entity.CustomerLimit, error)

	CreateCustomerTransaction(ctx context.Context, req entity.Transaction) (entity.Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) service {
	return service{
		repository: repo,
	}
}

func (s service) CheckoutTransaction(ctx context.Context, customerId int, req CheckoutLoanReq) response.ErrorResponse {
	//check tenor customer, validate with otr_price
	limitTenor, err := s.repository.GetCustomerTenor(ctx, customerId, req.Tenor)
	if err != nil {
		return *response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
	}
	if limitTenor.IsEmpty {
		return *response.Error("22101").WithError("not found").WithStatusCode(http.StatusBadRequest)
	}
	if *limitTenor.RemainingLimit < req.OTRPrice {
		return *response.Error("22104").WithStatusCode(http.StatusBadRequest)
	}

	//create transaction
	unixNumber := int(math.Floor(float64(time.Now().UnixMicro() / 1000)))
	contractNumber := `INV/XYZ/` + strconv.Itoa(unixNumber)

	newTransaction := entity.Transaction{
		ContractNumber:  contractNumber,
		CustomerId:      customerId,
		CustomerLimitId: limitTenor.ID,
		OTRPrice:        req.OTRPrice,
		AdminFee:        req.AdminFee,
		InterestPrice:   req.InterestPrice,
		Status:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	_, err = s.repository.CreateCustomerTransaction(ctx, newTransaction)
	if err != nil {
		return *response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
	}

	//update tenor
	remaining := *limitTenor.RemainingLimit - req.OTRPrice
	updTenor := entity.CustomerLimit{
		RemainingLimit: &remaining,
		UpdatedAt:      time.Now(),
	}
	_, err = s.repository.UpdateCustomerTenor(ctx, updTenor, limitTenor.ID)
	if err != nil {
		return *response.Error("22101").WithError(err.Error()).WithStatusCode(http.StatusInternalServerError)
	}

	return *response.NotError()
}
