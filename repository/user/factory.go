package user

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.UserRepository {
	var userRepo entities.UserRepository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		userRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return userRepo
}
