package transaction

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type service struct {
	repository entities.TransactionRepository
	validate   *validator.Validate
}

func NewService(repository entities.TransactionRepository) entities.TransactionService {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) GetById(id uint64) (transaction entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetAll() (transactions []entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(id int, data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me
	err = s.validate.Struct(&data)
	if err != nil {
		err = dapoint_api.ErrBadRequest
		return
	}

	totalBelanja, _ := strconv.Atoi(data.TotalBelanja)
	var pointEarn uint
	if totalBelanja >= 1000 {
		pointEarn = totalBelanja / 1000
	}

	data.PointEarn = pointEarn
	id, err = s.repository.Insert()
	if err != nil {
		return
	}
	return
}

func (s service) Modify(id int, data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetByParam(value interface{}) (transactions []entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}
