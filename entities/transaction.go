package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	UserID       uint   `json:"user_id" form:"user_id"`
	TotalBelanja uint   `json:"total_belanja" form:"total_belanja"`
	PointEarn    uint
}

//func ObjTransaction(dataName string, dataEmail string, dataPassword string) (user *User) {
//	return &User{
//		Name:     dataName,
//		Email:    dataEmail,
//		Password: dataPassword,
//	}
//}

type TransactionRepository interface {
	FindById(id uint64) (transaction Transaction, err error)
	FindAll() (transactions []Transaction, err error)
	FindByQuery(key string, value interface{}) (transaction Transaction, err error)
	Insert(Transaction) (transaction Transaction, err error)
	Update(id int, data Transaction) (transaction Transaction, err error)
	FindByAny(value interface{}) (res interface{}, err error)
}

type TransactionService interface {
	GetById(id uint64) (transaction Transaction, err error)
	GetAll() (transactions []Transaction, err error)
	GetByQuery(key string, value interface{}) (transaction Transaction, err error)
	Create(Transaction) (transaction Transaction, err error)
	Modify(id int, data Transaction) (transaction Transaction, err error)
	GetByAny(value interface{}) (res interface{}, err error)
}
