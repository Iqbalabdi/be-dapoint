package voucher

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.VoucherRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo MysqlRepository) FindById(id uint64) (voucher entities.Voucher, err error) {
	//TODO implement me
	if err = repo.db.Find(&voucher, id).Error; err != nil {
		return
	}

	return voucher, nil
}

func (repo MysqlRepository) FindAll() (vouchers []entities.Voucher, err error) {
	//TODO implement me
	if err = repo.db.Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
}

func (repo MysqlRepository) FindByQuery(key string, value interface{}) (voucher entities.Voucher, err error) {
	//TODO implement me

	err = repo.db.Where(key+" = ?", value).Find(&voucher).Error
	if err != nil {
		err = dapoint_api.ErrNotFound
		return
	}

	return voucher, nil
}

func (repo MysqlRepository) Insert(data entities.Voucher) (id uint64, err error) {
	//TODO implement me

	err = repo.db.Create(&data).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return
	}
	return
}

func (repo MysqlRepository) Update(id int, data entities.Voucher) (res entities.Voucher, err error) {
	//TODO implement me
	var voucher entities.Voucher
	repo.db.First(&voucher, "id = ?", id)

	if err = repo.db.Model(&voucher).Updates(map[string]interface{}{"name": data.Name, "max_limit": data.MaxLimit}).Error; err != nil {
		return voucher, err
	}
	return voucher, err
}
