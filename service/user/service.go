package user

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.UserRepository
	validate   *validator.Validate
}

func NewService(repository entities.UserRepository) entities.UserService {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) GetById(id uint64) (user entities.User, err error) {
	//TODO implement me
	user, err = s.repository.FindById(id)
	if err != nil {
		return user, err
	}
	return
}

func (s service) GetAll() (total int, users []entities.User, err error) {
	//TODO implement me
	total, users, err = s.repository.FindAll()
	if err != nil {
		return total, nil, err
	}
	return total, users, nil
}

func (s service) Create(data entities.User) (res entities.User, err error) {
	//TODO implement me
	err = s.validate.Struct(&data)
	if err != nil {
		err = dapoint_api.ErrBadRequest
		return
	}
	//var ok bool
	//if ok, err = s.validate(&data); !ok {
	//	return res, err
	//}
	//newUser := entities.ObjUser(data.Name, data.Email, data.Password)
	res, err = s.repository.Insert(data)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s service) Modify(id int, data entities.User) (user entities.User, err error) {
	//TODO implement
	res, err := s.repository.Update(id, data)
	if err != nil {
		return
	}

	return res, nil
}

func (s service) Login(data entities.UserLogin) (user entities.User, val bool, err error) {

	if err = s.validate.Struct(data); err != nil {
		return
	}

	res, err := s.repository.FindByQuery("email", data.Email)
	if err != nil || res.Password != data.Password {
		return res, false, err
	}

	return res, true, nil
}

func (s service) PointModify(id int, data entities.User) (ok bool, err error) {

	res, err := s.repository.PointUpdate(id, data)
	if err != nil {
		return
	}

	return res, nil
}

func (s service) GetTotal() (res interface{}, err error) {
	//TODO implement me
	res, err = s.repository.GetTotal()
	if err != nil {
		return
	}

	return res, nil
}
