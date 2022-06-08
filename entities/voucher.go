package entities

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	Name         string
	MaxLimit     uint
	HargaPoint   uint
	TipeVoucher  string
	UserVouchers []UserVoucher
}
