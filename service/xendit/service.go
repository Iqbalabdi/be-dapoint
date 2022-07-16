package xendit

import (
	"dapoint-api/entities"
	"github.com/go-playground/validator/v10"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"log"
)

type service struct {
	voucherRepo entities.VoucherRepository
	validate    *validator.Validate
}

func NewService(repoVoucher entities.VoucherRepository) XenditService {
	return &service{
		voucherRepo: repoVoucher,
		validate:    validator.New(),
	}
}

func (s service) CreateCharge(param string) (interface{}, error) {
	//TODO implement me
	voucherNominal, err := s.voucherRepo.FindNominalByName(param)
	if err != nil {
		return nil, err
	}

	xendit.Opt.SecretKey = "xnd_development_G1TruBndRrzJWNiL9ybYCdqy6TN0FrFr5FFAZIS8KmwG909640zJf5Z0KlPPvB"

	data := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         voucherNominal,
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
