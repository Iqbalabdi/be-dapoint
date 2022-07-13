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

func (repo PostgresRepository) FindAll() (vouchers []entities.Voucher, err error) {
	//TODO implement me
	if err = repo.db.Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
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
	voucherDetail := entities.ObjVoucher(data.Name, data.Stock, data.HargaPoint)
	repo.db.Raw("SELECT v.voucher_detail_id FROM voucher_details vd RIGHT JOIN vouchers v ON vd.id = v.voucher_detail_id WHERE vd.name = ?", data.TipeVoucher).Scan(&voucherDetail)

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

func (repo PostgresRepository) FindByParam(value interface{}) (vouchers []entities.Voucher, err error) {

	repo.db.Raw("SELECT * FROM vouchers v "+"INNER JOIN voucher_details vd ON v.voucher_detail_id = vd.id WHERE vd.name = ?", value).Scan(&vouchers)
	return vouchers, nil
}
