package user_voucher

import (
	"dapoint-api/entities"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.UserVoucherRepository
	validate   *validator.Validate
}

func NewService(repository entities.UserVoucherRepository) entities.UserVoucherRepository {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) FindById(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) FindAll() (vouchers []entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) FindByQuery(key string, value interface{}) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Insert(data entities.UserVoucher) (id uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(i int, voucher entities.UserVoucher) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Redeem(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}
