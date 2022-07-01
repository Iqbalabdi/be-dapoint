package voucher

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.VoucherRepository {
	var voucherRepo entities.VoucherRepository

	if dbCon.Driver == util.Mysql {
		// existing tetep jalan
		voucherRepo = NewMysqlRepository(dbCon.Mysql)
	}

	return voucherRepo
}
