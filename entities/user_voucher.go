package entities

import "gorm.io/gorm"

type UserVoucher struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	VoucherID string
	UserID    string
	Quantity  string
}
