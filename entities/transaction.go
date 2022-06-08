package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	UserID       uint
	TotalBelanja string `gorm:"unique"`
	PointEarn    uint
}

//func ObjTransaction(dataName string, dataEmail string, dataPassword string) (user *User) {
//	return &User{
//		Name:     dataName,
//		Email:    dataEmail,
//		Password: dataPassword,
//	}
//}
