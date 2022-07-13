package xendit

import (
	"dapoint-api/api/response"
	v1 "dapoint-api/api/v1"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	payload XenditCallbackPayload
}

func NewController(payload XenditCallbackPayload) *Controller {
	return &Controller{
		payload: payload,
	}
}

func (controller *Controller) AcceptCallback(c echo.Context) error {

	err := c.Bind(&controller.payload)
	if err != nil {
		return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(v1.GetErrorStatus(err), response.ApiResponse{
		Status:  "success",
		Message: "Payload accepted",
	})
}
