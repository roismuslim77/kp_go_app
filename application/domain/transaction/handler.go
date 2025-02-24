package transaction

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"simple-go/pkg/response"
)

type Service interface {
	CheckoutTransaction(ctx context.Context, customerId int, payload CheckoutLoanReq) response.ErrorResponse
}

type handler struct {
	service Service
}

func NewHandler(svc Service) handler {
	return handler{
		service: svc,
	}
}

func (h handler) CheckoutTransaction(ctx *gin.Context) {
	customerId := ctx.GetInt("customerId")

	var req CheckoutLoanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			resp.WithArgsMessage(ve[0].Field(), ve[0].Tag())
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	err := h.service.CheckoutTransaction(ctx, customerId, req)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22151")
	ctx.JSON(resp.StatusCode, resp)
}
