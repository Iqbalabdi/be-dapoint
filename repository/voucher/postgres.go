package voucher

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.VoucherRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo PostgresRepository) FindById(id uint64) (voucher entities.Voucher, err error) {
	//TODO implement me
	if err = repo.db.First(&voucher, id).Error; err != nil {
		return voucher, err
	}
	return voucher, nil
}

func (repo PostgresRepository) FindAll() (total int, vouchers []entities.Voucher, err error) {
	//TODO implement me
	if err = repo.db.Raw("SELECT COUNT(*) FROM users").Scan(&total).Error; err != nil {
		return total, nil, err
	}

	if err = repo.db.Find(&vouchers).Error; err != nil {
		return total, nil, err
	}
	return total, vouchers, nil
}

func (repo PostgresRepository) FindByQuery(key string, value interface{}) (voucher entities.Voucher, err error) {
	//TODO implement me

	err = repo.db.Where(key+" = ?", value).Find(&voucher).Error
	if err != nil {
		err = dapoint_api.ErrNotFound
		return
	}

	return voucher, nil
}

func (repo PostgresRepository) Insert(data entities.VoucherDTO) (id uint64, err error) {
	//TODO implement me
	//var voucherDetail entities.VoucherDetail
	//var newVoucher entities.Voucher
	voucherDetail := entities.ObjVoucher(data.Name, data.Stock, data.HargaPoint, data.Nominal)
	repo.db.Raw("SELECT id FROM voucher_details vd WHERE vd.name=?", data.TipeVoucher).Scan(&voucherDetail.VoucherDetailID)
	err = repo.db.Create(&voucherDetail).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return
	}
	return voucherDetail.ID, nil
}

func (repo PostgresRepository) Update(id int, data entities.Voucher) (res entities.Voucher, err error) {
	//TODO implement me
	var voucher entities.Voucher
	repo.db.First(&voucher, "id = ?", id)

	//repo.db.Raw("UPDATE vouchers SET "+key+" = ? "+"WHERE id = ?", value, id).Scan(&vouchers)
	if err = repo.db.Model(&voucher).Updates(map[string]interface{}{"name": data.Name, "stock": data.Stock, "harga_point": data.HargaPoint}).Error; err != nil {
		return voucher, err
	}
	return voucher, err
}

func (repo PostgresRepository) FindByType(value interface{}) (vouchers []entities.Voucher, err error) {

	repo.db.Raw("SELECT * FROM vouchers WHERE voucher_detail_id IN (SELECT id FROM voucher_details WHERE name = ?)", value).Scan(&vouchers)
	return vouchers, nil
}

func (repo PostgresRepository) GetTotal() (res interface{}, err error) {
	//TODO implement me
	var total uint
	if err = repo.db.Raw("SELECT COUNT(*) FROM users").Scan(&total).Error; err != nil {
		return nil, err
	}

	return total, nil
}

func (repo PostgresRepository) FindNominalByName(value interface{}) (res float64, err error) {
	//TODO implement me
	if err = repo.db.Raw("SELECT nominal FROM vouchers WHERE NAME = ?", value).Scan(&res).Error; err != nil {
		return
	}

	return res, nil
}
