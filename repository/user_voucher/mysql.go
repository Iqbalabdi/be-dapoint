package user_voucher

import (
	"dapoint-api/entities"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) entities.UserVoucherRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (m MysqlRepository) FindById(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepository) FindAll() (vouchers []entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepository) FindByQuery(key string, value interface{}) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepository) Insert(data entities.UserVoucher) (id uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepository) Update(i int, voucher entities.UserVoucher) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepository) Redeem(id uint64) (res entities.UserVoucher, err error) {
	//TODO implement me
	panic("implement me")
}
