package entities

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey"`
	Name            string `json:"name" form:"name"`
	Stock           uint   `json:"stock" form:"stock"`
	HargaPoint      uint   `json:"harga_point" form:"harga_point"`
	Nominal         uint   `json:"nominal" form:"nominal"`
	VoucherDetailID uint   `json:"voucher_detail_id" form:"voucher_detail_id"`
	UserVouchers    []UserVoucher
}

type VoucherDTO struct {
	Name        string `json:"name" form:"name"`
	Stock       uint   `json:"stock" form:"stock"`
	HargaPoint  uint   `json:"harga_point" form:"harga_point"`
	Nominal     uint   `json:"nominal" form:"nominal"`
	TipeVoucher string `json:"tipe_voucher" form:"tipe_voucher"`
}

func ObjVoucher(dataName string, dataStock uint, dataHargaPoint uint, dataNominal uint) (voucher *Voucher) {
	return &Voucher{
		Name:       dataName,
		Stock:      dataStock,
		HargaPoint: dataHargaPoint,
		Nominal:    dataNominal,
	}
}

type VoucherRepository interface {
	FindById(id uint64) (voucher Voucher, err error)
	FindAll() (total int, vouchers []Voucher, err error)
	FindByQuery(key string, value interface{}) (voucher Voucher, err error)
	Insert(data VoucherDTO) (id uint64, err error)
	Update(id int, data Voucher) (voucher Voucher, err error)
	FindByType(value interface{}) (vouchers []Voucher, err error)
	GetTotal() (res interface{}, err error)
	FindNominalByName(value interface{}) (res float64, err error)
}

type VoucherService interface {
	GetById(id uint64) (voucher Voucher, err error)
	GetAll() (total int, vouchers []Voucher, err error)
	Create(data VoucherDTO) (id uint64, err error)
	Modify(id int, data Voucher) (voucher Voucher, err error)
	GetByType(value interface{}) (vouchers []Voucher, err error)
	GetTotal() (res interface{}, err error)
}
