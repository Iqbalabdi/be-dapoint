package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Photo        string
	Role         string
	TotalPoint   int
	Transactions []Transaction
	UserVouchers []UserVoucher
}

//func ObjUser(dataName string, dataEmail string, dataPassword string) (user *User) {
//	return &User{
//		Name:     dataName,
//		Email:    dataEmail,
//		Password: dataPassword,
//	}
//}
