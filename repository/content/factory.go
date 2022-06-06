package content

import (
	"dapoint-api/business/content"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
	var contentRepo content.Repository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		contentRepo = NewPostgresRepository(dbCon.Postgres)
	} else {
		// develop teknologi ini
		contentRepo = NewStaticRepository()
	}

	return contentRepo
}
