package redeem_voucher

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.RedeemVoucherRepository {
	var redeemVoucherRepo entities.RedeemVoucherRepository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		redeemVoucherRepo = NewPostgresRepository(dbCon.Postgres)
	} else if dbCon.Driver == util.Mysql {
		// existing tetep jalan
		redeemVoucherRepo = NewPostgresRepository(dbCon.Mysql)
	}

	return redeemVoucherRepo
}
