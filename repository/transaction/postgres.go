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

func (p PostgresRepository) FindByAny(value interface{}) (res interface{}, err error) {
	//TODO implement me
	res, err = p.GetAllUserPoint(value)
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return nil, err
	}
	return res, nil
}

func (p PostgresRepository) GetAllUserPoint(value interface{}) (res interface{}, err error) {

	type UserPointByMonth struct {
		TotalPoint int
		Januari    uint64 `json:"januari"`
		Februari   uint64 `json:"februari"`
		Maret      uint64 `json:"maret"`
		April      uint64 `json:"april"`
		Mei        uint64 `json:"mei"`
		Juni       uint64 `json:"juni"`
		Juli       uint64 `json:"juli"`
		Agustus    uint64 `json:"agustus"`
		September  uint64 `json:"september"`
		Oktober    uint64 `json:"oktober"`
		November   uint64 `json:"november"`
		Desember   uint64 `json:"desember"`
	}

	var userPoint UserPointByMonth
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=1").Scan(&userPoint.Januari)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=2").Scan(&userPoint.Februari)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=3").Scan(&userPoint.Maret)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=4").Scan(&userPoint.April)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=5").Scan(&userPoint.Mei)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=6").Scan(&userPoint.Juni)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=7").Scan(&userPoint.Juli)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=8").Scan(&userPoint.Agustus)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=9").Scan(&userPoint.September)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=10").Scan(&userPoint.Oktober)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=11").Scan(&userPoint.November)
	p.db.Raw("SELECT MontPoint.totalpoint FROM " +
		"(SELECT MONTH(CAST(created_at AS DATE)) AS Dateonly, SUM(point_earn) AS totalpoint FROM transactions GROUP BY Dateonly) " +
		"MontPoint WHERE MontPoint.Dateonly=12").Scan(&userPoint.Desember)
	p.db.Raw("SELECT SUM(point_earn) FROM transactions").Scan(&userPoint.TotalPoint)
	return userPoint, nil
}
