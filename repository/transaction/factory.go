package transaction

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.TransactionRepository {
	var transactionRepo entities.TransactionRepository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionRepo
}
