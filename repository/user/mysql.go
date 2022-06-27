package user

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) entities.UserRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo MysqlRepository) FindById(id uint64) (user entities.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo MysqlRepository) FindAll() (users []entities.User, err error) {
	//TODO implement me
	if err = repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo MysqlRepository) FindByQuery(key string, value interface{}) (user entities.User, err error) {
	//TODO implement me
	var userAuth entities.User
	if err = repo.db.Where("email = ?", user.Email).First(&userAuth).Error; err != nil {
		return
	}
	return userAuth, nil
}

func (repo MysqlRepository) Insert(data entities.User) (id uint64, err error) {
	//TODO implement me

	err = repo.db.Create(&data).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return
	}
	return
}

func (repo MysqlRepository) Update(data entities.User) (user entities.User, err error) {
	//TODO implement me
	panic("implement me")
}
