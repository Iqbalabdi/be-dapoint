package xendit

import (
	"dapoint-api/entities"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"log"
)

var voucherID uint64

type service struct {
	voucherRepo entities.VoucherRepository
	validate    *validator.Validate
	redeemRepo  entities.RedeemVoucherRepository
	userRepo    entities.UserRepository
}

func NewService(repoVoucher entities.VoucherRepository, redeemVoucherRepo entities.RedeemVoucherRepository, userRepo entities.UserRepository) XenditService {
	return &service{
		voucherRepo: repoVoucher,
		validate:    validator.New(),
		redeemRepo:  redeemVoucherRepo,
		userRepo:    userRepo,
	}
}

func (s service) CreateCharge(c echo.Context, param string) (interface{}, error) {
	//TODO implement me
	res, err := s.voucherRepo.FindNominalByName(param)
	voucherID = res.ID
	//fmt.Println(res.ID, res.Nominal)
	if err != nil {
		return nil, err
	}

	xendit.Opt.SecretKey = "xnd_development_G1TruBndRrzJWNiL9ybYCdqy6TN0FrFr5FFAZIS8KmwG909640zJf5Z0KlPPvB"

	data := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         float64(res.Nominal),
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_DANA",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://redirect.me/payment",
		},
		Metadata: map[string]interface{}{
			"branch_code": "tree_branch",
		},
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	return charge, nil
}

func (s service) PaymentStatusCallback(userID uint64, param string) (res interface{}, err error) {
	//TODO implement me
	var data CallbackStatusData
	json.Unmarshal([]byte(param), &data)
	if data.Data.Status == "SUCCEEDED" {

		//get voucher point
		voucher, err := s.voucherRepo.FindById(voucherID)
		pointCharge := voucher.HargaPoint

		// update user TotalPoint
		user, err := s.userRepo.FindById(userID)
		user.TotalPoint = user.TotalPoint - uint64(pointCharge)
		_, err = s.userRepo.PointUpdate(int(userID), user)
		if err != nil {
			return nil, err
		}

		// Decrease voucher stock
		voucher.Stock -= 1
		_, err = s.voucherRepo.Update(int(voucherID), voucher)

		// insert redeem voucher
		res, err = s.redeemRepo.Insert(voucherID, int(userID))
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return
}

type CallbackStatusData struct {
	Data DataAssert
}

type DataAssert struct {
	Status       string  `json:"status"`
	ChargeAmount float64 `json:"charge_amount"`
}
