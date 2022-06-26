package modules

import (
	"dapoint-api/api"
	contentV1Controller "dapoint-api/api/v1/content"
	"dapoint-api/config"
	contentRepo "dapoint-api/repository/content"
	contentService "dapoint-api/service/content"

	authController "dapoint-api/api/v1/auth"
	authService "dapoint-api/service/auth"

	userController "dapoint-api/api/v1/user"
	userRepo "dapoint-api/repository/user"
	userService "dapoint-api/service/user"
	"dapoint-api/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	authPermitService := authService.NewService(config)
	authPermitController := authController.NewController(authPermitService)

	// user
	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository)
	userPermitController := userController.NewController(userPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		AuthController:      authPermitController,
		UserController:      userPermitController,
	}

	return controllers
}
