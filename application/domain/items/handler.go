package items

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"inventory-simple-go/application/entity"
	"inventory-simple-go/pkg/response"
	"net/http"
	"strconv"
)

type Service interface {
	GetAllItem(ctx context.Context, filter FilterListing) ([]entity.Item, PaginateListing, response.ErrorResponse)
	GetOneItem(ctx context.Context, id int) (entity.Item, response.ErrorResponse)
	CreateItem(ctx context.Context, req CreateDataRequest) response.ErrorResponse
	UpdateItem(ctx context.Context, req CreateDataRequest, id int) (entity.Item, response.ErrorResponse)
	DeleteItem(ctx context.Context, id int) response.ErrorResponse
}

type handler struct {
	service Service
}

func NewHandler(svc Service) handler {
	return handler{
		service: svc,
	}
}

func (h handler) DeleteItem(ctx *gin.Context) {
	ids := ctx.Param("id")
	if ids == "" {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	id, _ := strconv.Atoi(ids)
	err := h.service.DeleteItem(ctx, id)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22153")
	ctx.JSON(resp.StatusCode, resp)
}

func (h handler) UpdateItem(ctx *gin.Context) {
	var req CreateDataRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			resp.WithArgsMessage(ve[0].Field(), ve[0].Tag())
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	ids := ctx.Param("id")
	if ids == "" {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	id, _ := strconv.Atoi(ids)
	item, err := h.service.UpdateItem(ctx, req, id)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22154").WithData(item)
	ctx.JSON(resp.StatusCode, resp)
}

func (h handler) CreateItem(ctx *gin.Context) {
	var req CreateDataRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			resp.WithArgsMessage(ve[0].Field(), ve[0].Tag())
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	err := h.service.CreateItem(ctx, req)
	if !err.IsNoError {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	resp := response.Success("22151")
	ctx.JSON(resp.StatusCode, resp)
}

func (h handler) GetOneItem(ctx *gin.Context) {
	ids := ctx.Param("id")
	if ids == "" {
		resp := response.Error("22102").WithStatusCode(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	id, _ := strconv.Atoi(ids)
	agency, err := h.service.GetOneItem(ctx, id)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22152").WithData(agency)
	ctx.JSON(resp.StatusCode, resp)
}

func (h handler) GetAllItem(ctx *gin.Context) {
	var req FilterListing
	req.Page = ctx.Query("page")
	req.Size = ctx.Query("size")

	cards, paginate, err := h.service.GetAllItem(ctx, req)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22152").WithData(cards).
		WithTotalPage(int(paginate.TotalPage)).
		WithCount(int(paginate.TotalData))

	ctx.JSON(resp.StatusCode, resp)
}
