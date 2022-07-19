package entities

import "gorm.io/gorm"

type RedeemVoucher struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	VoucherID int
	UserID    int
}

type RedeemVoucherRepository interface {
	Insert(voucherID uint64, userID int) (RedeemVoucher, error)
	FindAll() (total int, redeem []RedeemVoucher, err error)
	FindById(id uint64) (res interface{}, err error)
}

type RedeemVoucherService interface {
	GetAll() (total int, redeem []RedeemVoucher, err error)
	GetById(id uint64) (res interface{}, err error)
}
