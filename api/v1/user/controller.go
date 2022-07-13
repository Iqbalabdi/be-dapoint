package user

import (
	"dapoint-api/api/middleware"
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"dapoint-api/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	service entities.UserService
	UJwt    middleware.JWTService
}

func NewController(service entities.UserService, jwt middleware.JWTService) *Controller {
	return &Controller{
		service: service,
		UJwt:    jwt,
	}

}

// GetAll godoc
// @Summary      Get all users
// @Description  Retrieve list of all users
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200	{object}	user.User
// @Failure      404	{object}	response.ApiResponse
// @Failure      403	{string}	string		"Unauthorized"
// @Failure      500	{object}	response.ApiResponse
// @Router       /users/getall [get]
func (controller *Controller) GetAll(c echo.Context) error {
	total, listUser, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Count:  total,
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

// Create godoc
// @Summary      Create user
// @Description  create user adn save to db
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponseSuccess
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/register [post]
func (controller *Controller) Create(c echo.Context) (err error) {

	var newUser entities.User
	err = c.Bind(&newUser)

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

func (controller *Controller) Modify(c echo.Context) (err error) {

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  "fail",
			Message: "put user id in endpoint",
		})
	}

	userParamsId, _ := strconv.Atoi(params)

	var data entities.User
	err = c.Bind(&data)
	res, err := controller.service.Modify(userParamsId, data)

	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Status:  "success update user with id : " + strconv.Itoa(userParamsId),
		Message: res,
	})

}

func (controller *Controller) Delete(c echo.Context) (err error) {
	panic("")
}

func (controller *Controller) Login(c echo.Context) (err error) {

	var userLogin entities.UserLogin
	//var data entities.User
	err = c.Bind(&userLogin)
	var ok bool
	data, ok, err := controller.service.Login(userLogin)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if ok == false {
		return c.JSON(http.StatusUnauthorized, response.ApiResponse{
			Status:  "Unauthorized",
			Message: "Email or Password is wrong",
		})
	}
	token, err := controller.UJwt.GenerateToken(data)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
			Status: "error",
			Data:   token,
		})
	}
	return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
		Status:  "success",
		Message: token,
	})
}

func (controller *Controller) PointModify(c echo.Context) (err error) {
	//var userPoint entities.User
	////var data entities.User
	//err = c.Bind(&userPoint)
	//var ok bool
	//
	//userParamsId, _ := strconv.Atoi(c.Param("id"))
	//fmt.Println(userPoint.TotalPoint)
	var ok bool

	params := c.Param("id")
	if params == "" {
		return c.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  "fail",
			Message: "put user id in endpoint",
		})
	}

	userParamsId, _ := strconv.Atoi(params)

	var data entities.User
	err = c.Bind(&data)

	ok, err = controller.service.PointModify(userParamsId, data)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if ok == false {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
		Status:  "success",
		Message: "Point Updated",
	})
}

func (controller *Controller) GetTotal(c echo.Context) (err error) {
	totalUser, err := controller.service.GetTotal()
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponseSuccess{
		Status: "success",
		Data:   totalUser,
	})
}
