package user

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
	service entities.UserService
}

func NewController(service entities.UserService) *Controller {
	return &Controller{
		service: service,
	}

}

func (controller *Controller) GetAll(c echo.Context) error {
	listUser, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   listUser,
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
	user, err := controller.service.GetById(uint64(id))
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   user,
	})
}

func (controller *Controller) Create(c echo.Context) (err error) {

	var newUser entities.User
	err = c.Bind(&newUser)
	fmt.Println("ini controller")
	fmt.Println(newUser)
	user, err := controller.service.Create(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponseSuccess{
		Status: "success",
		Data:   user,
	})
}
