package entities

import "gorm.io/gorm"

type UserVoucher struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	VoucherID string
	UserID    string
	Quantity  string
}

type UserVoucherRepository interface {
	FindById(id uint64) (res UserVoucher, err error)
	FindAll() (vouchers []UserVoucher, err error)
	FindByQuery(key string, value interface{}) (res UserVoucher, err error)
	Insert(data UserVoucher) (id uint64, err error)
	Update(int, UserVoucher) (res UserVoucher, err error)
}

type UserVoucherService interface {
	GetById(id uint64) (res UserVoucher, err error)
	GetAll() (vouchers []UserVoucher, err error)
	Create(data UserVoucher) (id uint64, err error)
	Modify(int, UserVoucher) (res UserVoucher, err error)
}
