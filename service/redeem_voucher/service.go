package redeem_voucher

import (
	"dapoint-api/entities"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.RedeemVoucherRepository
	validate   *validator.Validate
}

func NewService(repository entities.RedeemVoucherRepository) entities.RedeemVoucherService {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) GetAll() (total int, redeem []entities.RedeemVoucher, err error) {
	//TODO implement me
	total, redeem, err = s.repository.FindAll()
	if err != nil {
		return total, nil, err
	}
	return total, redeem, nil
}

func (s service) GetById(id uint64) (res interface{}, err error) {
	//TODO implement me
	res, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	return res, nil
}
