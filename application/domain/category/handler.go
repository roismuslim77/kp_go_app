package category

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"inventory-simple-go/application/entity"
	"inventory-simple-go/pkg/response"
	"net/http"
)

type Service interface {
	GetAllCategory(ctx context.Context) ([]entity.Category, response.ErrorResponse)
	CreateCategory(ctx context.Context, name string) response.ErrorResponse
}

type handler struct {
	service Service
}

func NewHandler(svc Service) handler {
	return handler{
		service: svc,
	}
}

func (h handler) CreateCategory(ctx *gin.Context) {
	var req CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			resp.WithArgsMessage(ve[0].Field(), ve[0].Tag())
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	err := h.service.CreateCategory(ctx, req.Name)
	if !err.IsNoError {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	resp := response.Success("22151")
	ctx.JSON(resp.StatusCode, resp)
}

func (h handler) GetAllCategory(ctx *gin.Context) {
	cards, err := h.service.GetAllCategory(ctx)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22152").WithData(cards)
	ctx.JSON(resp.StatusCode, resp)
}
