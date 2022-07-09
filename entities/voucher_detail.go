package entities

import "gorm.io/gorm"

type VoucherDetail struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Vouchers []Voucher
	Name     string `json:"name" form:"name"`
}
