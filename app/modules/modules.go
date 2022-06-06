package modules

import (
	"dapoint-api/api"
	contentV1Controller "dapoint-api/api/v1/content"
	contentV2Controller "dapoint-api/api/v2/content"
	contentService "dapoint-api/business/content"
	"dapoint-api/config"
	contentRepo "dapoint-api/repository/content"

	authController "dapoint-api/api/v1/auth"
	authService "dapoint-api/business/auth"
	"dapoint-api/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	contentV2PermitController := contentV2Controller.NewController(contentPermitService)

	authPermitService := authService.NewService(config)
	authPermitController := authController.NewController(authPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		ContentV2Controller: contentV2PermitController,
		AuthController:      authPermitController,
	}

	return controllers
}
