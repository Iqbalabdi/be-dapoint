package transaction

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.TransactionRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p PostgresRepository) FindById(id uint64) (transaction entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) FindAll() (transactions []entities.Transaction, err error) {
	//TODO implement me
	if err = p.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p PostgresRepository) FindByQuery(key string, value interface{}) (transaction entities.Transaction, err error) {
	//TODO implement me
	err = p.db.Where(key+" = ?", value).Find(&transaction).Error
	if err != nil {
		err = dapoint_api.ErrNotFound
		return
	}

	return transaction, nil
}

func (p PostgresRepository) Insert(data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me
	err = p.db.Create(&data).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return transaction, err
	}
	return data, nil
}

func (p PostgresRepository) Update(id int, data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me
	//p.db.First(&transaction, "id = ?", id)
	//
	////repo.db.Raw("UPDATE vouchers SET "+key+" = ? "+"WHERE id = ?", value, id).Scan(&vouchers)
	//if err = p.db.Model(&transaction).Updates(map[string]interface{}{"name": data.Name, "stock": data.Stock, "harga_point": data.HargaPoint}).Error; err != nil {
	//	return transaction, err
	//}
	//return transaction, err
	panic("implement me")
}

func (p PostgresRepository) FindByParam(value interface{}) (transactions []entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}
