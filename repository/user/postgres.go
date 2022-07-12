package user

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) entities.UserRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo PostgresRepository) FindById(id uint64) (user entities.User, err error) {
	//TODO implement me
	if err = repo.db.Find(&user, id).Error; err != nil {
		return
	}

	return user, nil
}

func (repo PostgresRepository) FindAll() (users []entities.User, err error) {
	//TODO implement me
	if err = repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo PostgresRepository) FindByQuery(key string, value interface{}) (user entities.User, err error) {
	//TODO implement me

	err = repo.db.Where(key+" = ?", value).Find(&user).Error
	if err != nil {
		err = dapoint_api.ErrNotFound
		return
	}

	return user, nil
}

func (repo PostgresRepository) Insert(data entities.User) (res entities.User, err error) {
	//TODO implement me

	err = repo.db.Create(&data).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return res, err
	}

	return data, nil
}

func (repo PostgresRepository) Update(id int, data entities.User) (res entities.User, err error) {
	//TODO implement me
	var user entities.User

	repo.db.First(&user, "id = ?", id)

	if err = repo.db.Model(&user).Updates(map[string]interface{}{"name": data.Name, "email": data.Email, "password": data.Password, "photo": data.Photo}).Error; err != nil {
		return user, err
	}
	return user, err
}

func (repo PostgresRepository) PointUpdate(id int, data entities.User) (ok bool, err error) {
	//TODO implement me
	var userPoint entities.User
	repo.db.First(&userPoint, "id = ?", id)

	if err = repo.db.Model(&userPoint).Updates(map[string]interface{}{"total_point": data.TotalPoint}).Error; err != nil {
		return false, err
	}
	return true, err
}
