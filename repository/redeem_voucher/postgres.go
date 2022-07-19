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

func (m PostgresRepository) FindAll() (total int, redeem []entities.RedeemVoucher, err error) {
	//TODO implement me
	if err = m.db.Raw("SELECT COUNT(*) FROM vouchers").Scan(&total).Error; err != nil {
		return total, nil, err
	}

	if err = m.db.Find(&redeem).Error; err != nil {
		return total, nil, err
	}
	return total, redeem, nil
}

func (m PostgresRepository) FindById(id uint64) (res interface{}, err error) {
	type RedeemDTO struct {
		Name       string `json:"name"`
		HargaPoint uint   `json:"harga_point"`
		Nominal    uint   `json:"nominal"`
	}
	var redeemDTO RedeemDTO
	m.db.Raw("SELECT v.name, v.harga_point, v.nominal FROM vouchers v "+
		"JOIN user_vouchers uv ON v.id = uv.voucher_id "+
		"JOIN users u ON uv.user_id = u.id "+
		"WHERE u.id = ?", id).Scan(&redeemDTO)

	return redeemDTO, nil
}
