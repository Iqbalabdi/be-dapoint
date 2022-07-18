package xendit

import (
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"dapoint-api/service/xendit"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Controller struct {
	callbackPayload XenditCallbackPayload
	service         xendit.XenditService
}

type Payload struct {
	VoucherName string `json:"voucher_name" form:"voucher_name"`
}

func NewController(payload XenditCallbackPayload, service xendit.XenditService) *Controller {
	return &Controller{
		callbackPayload: payload,
		service:         service,
	}
}

func (controller *Controller) AcceptCallback(c echo.Context) error {
	var iface interface{}
	err := c.Bind(&iface)
	asByteJson, _ := json.Marshal(iface)
	//fmt.Println("masuk : ", string(asByteJson))
	userID := c.Get("userID")

	userIdconv, _ := strconv.Atoi(userID.(string))
	res, err := controller.service.PaymentStatusCallback(uint64(userIdconv), string(asByteJson))
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

func (controller *Controller) CreateCharge(c echo.Context) error {

	param := c.Param("name")
	res, err := controller.service.CreateCharge(c, param)
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
