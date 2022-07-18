package xendit

import "github.com/labstack/echo/v4"

type XenditService interface {
	CreateCharge(c echo.Context, param string) (interface{}, error)
	PaymentStatusCallback(userID uint64, param string) (res interface{}, err error)
}
