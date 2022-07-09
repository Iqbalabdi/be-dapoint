package user

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.UserRepository {
	var userRepo entities.UserRepository

	if dbCon.Driver == util.Mysql {
		// existing tetep jalan
		userRepo = NewPostgresRepository(dbCon.Mysql)
	}

	return userRepo
}
