package redeem_voucher

import (
	"dapoint-api/entities"
)

type Controller struct {
	service entities.RedeemVoucherService
}

func NewController(service entities.RedeemVoucherService) *Controller {
	return &Controller{
		service: service,
	}
}

//func (controller *Controller) Redeem(c echo.Context) (err error) {
//
//	var newRedeemVoucher entities.RedeemVoucher
//	err = c.Bind(&newRedeemVoucher)
//
//	voucher, err := controller.service.Create(newRedeemVoucher)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, response.ApiResponse{
//			Status:  "error",
//			Message: err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, response.ApiResponseSuccess{
//		Status: "success",
//		Data:   voucher,
//	})
//}
