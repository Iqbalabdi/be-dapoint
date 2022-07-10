package voucher

import (
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"dapoint-api/entities"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	service entities.VoucherService
}

func NewController(service entities.VoucherService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	listVoucher, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   listVoucher,
	})
}

func (controller *Controller) GetByID(c echo.Context) error {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	voucher, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   voucher,
	})
}

func (controller *Controller) Create(c echo.Context) (err error) {

	var newVoucher entities.Voucher
	err = c.Bind(&newVoucher)

	voucher, err := controller.service.Create(newVoucher)
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

func (controller *Controller) Modify(c echo.Context) (err error) {

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  "fail",
			Message: "put voucher id in endpoint",
		})
	}

	userParamsId, _ := strconv.Atoi(params)

	var data entities.Voucher
	err = c.Bind(&data)
	fmt.Println(data.MaxLimit)
	res, err := controller.service.Modify(userParamsId, data)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Status:  "success update voucher with id : " + strconv.Itoa(userParamsId),
		Message: res,
	})

}

func (controller *Controller) Delete(c echo.Context) (err error) {
	panic("")
}

func (controller *Controller) GetByParams(c echo.Context) (err error) {

	param := c.Param("tipe")
	listVoucher, err := controller.service.GetByParam(param)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   listVoucher,
	})
}
