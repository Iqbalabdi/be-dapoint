package user_voucher

import (
	"dapoint-api/entities"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.UserVoucherRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (m PostgresRepository) FindById(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m PostgresRepository) FindAll() (vouchers []entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m PostgresRepository) FindByQuery(key string, value interface{}) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m PostgresRepository) Insert(data entities.UserVoucher) (id uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (m PostgresRepository) Update(i int, voucher entities.UserVoucher) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m PostgresRepository) Redeem(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}
