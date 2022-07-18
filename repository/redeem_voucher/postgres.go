package redeem_voucher

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.RedeemVoucherRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (m PostgresRepository) Insert(voucherID uint64, userID int) (entities.RedeemVoucher, error) {
	//TODO implement me
	var redeem entities.RedeemVoucher
	redeem.UserID = userID
	redeem.VoucherID = int(voucherID)

	err := m.db.Create(&redeem).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return entities.RedeemVoucher{}, nil
	}

	return redeem, nil
}
