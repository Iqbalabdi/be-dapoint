package transaction

import (
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"dapoint-api/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	service entities.TransactionService
}

func NewController(service entities.TransactionService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	listTransaction, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   listTransaction,
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

func (controller *Controller) Create(c echo.Context) (err error) {

	userId, _ := strconv.Atoi(c.Param("userid"))
	var newTransaction entities.Transaction
	err = c.Bind(&newTransaction)

	res, err := controller.service.Create(userId, newTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponseSuccess{
		Status: "success",
		Data:   res,
	})
}

func (controller *Controller) Modify(c echo.Context) (err error) {

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  "fail",
			Message: "put res id in endpoint",
		})
	}

	userParamsId, _ := strconv.Atoi(params)

	var data entities.Transaction
	err = c.Bind(&data)

	res, err := controller.service.Modify(userParamsId, data)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Status:  "success update transaction with id : " + strconv.Itoa(userParamsId),
		Message: res,
	})

}

func (controller *Controller) Delete(c echo.Context) (err error) {
	panic("")
}

func (controller *Controller) GetByParams(c echo.Context) (err error) {

	param := c.Param("tipe")
	listTransaction, err := controller.service.GetByParam(param)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   listTransaction,
	})
}
