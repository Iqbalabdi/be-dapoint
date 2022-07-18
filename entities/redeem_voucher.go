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
}

type RedeemVoucherService interface {
}
