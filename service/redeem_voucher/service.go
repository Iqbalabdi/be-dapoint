package redeem_voucher

import (
	"dapoint-api/entities"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.RedeemVoucherRepository
	validate   *validator.Validate
}

func NewService(repository entities.RedeemVoucherRepository) entities.RedeemVoucherRepository {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) Insert(voucherID uint64, userID int) (res entities.RedeemVoucher, err error) {
	//TODO implement me
	panic("implement me")
}
