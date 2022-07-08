package user_voucher

import (
	"dapoint-api/api/response"
	"dapoint-api/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service entities.UserVoucherService
}

func NewController(service entities.UserVoucherService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Redeem(c echo.Context) (err error) {

	var newUserVoucher entities.UserVoucher
	err = c.Bind(&newUserVoucher)

	voucher, err := controller.service.Create(newUserVoucher)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponseSuccess{
		Status: "success",
		Data:   voucher,
	})
}
