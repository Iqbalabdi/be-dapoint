package redeem_voucher

import (
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"dapoint-api/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	service entities.RedeemVoucherService
}

func NewController(service entities.RedeemVoucherService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller Controller) GetAll(c echo.Context) error {
	total, res, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Count:  total,
		Data:   res,
	})
}

func (controller Controller) GetByUserID(c echo.Context) error {
	params := c.Param("userid")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	res, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   res,
	})
}
