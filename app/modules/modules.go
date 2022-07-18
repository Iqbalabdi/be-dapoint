package modules

import (
	"dapoint-api/api"
	"dapoint-api/api/middleware"
	contentV1Controller "dapoint-api/api/v1/content"
	"dapoint-api/config"
	contentRepo "dapoint-api/repository/content"
	contentService "dapoint-api/service/content"

	authController "dapoint-api/api/v1/auth"
	authService "dapoint-api/service/auth"

	userController "dapoint-api/api/v1/user"
	userRepo "dapoint-api/repository/user"
	userService "dapoint-api/service/user"

	voucherController "dapoint-api/api/v1/voucher"
	voucherRepo "dapoint-api/repository/voucher"
	voucherService "dapoint-api/service/voucher"

	transactionController "dapoint-api/api/v1/transaction"
	transactionRepo "dapoint-api/repository/transaction"
	transactionService "dapoint-api/service/transaction"

	redeemController "dapoint-api/api/v1/redeem_voucher"
	redeemRepo "dapoint-api/repository/redeem_voucher"
	redeemService "dapoint-api/service/redeem_voucher"

	xenditController "dapoint-api/api/xendit"
	xenditPayload "dapoint-api/api/xendit"
	xenditService "dapoint-api/service/xendit"

	"dapoint-api/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	authPermitService := authService.NewService(config)
	authPermitController := authController.NewController(authPermitService)

	// jwt
	middlewarePermitJwt := middleware.NewJwtService(config.App.JWTKey)
	// user
	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository)
	userPermitController := userController.NewController(userPermitService, middlewarePermitJwt)

	//voucher
	voucherPermitRepository := voucherRepo.RepositoryFactory(dbCon)
	voucherPermitService := voucherService.NewService(voucherPermitRepository)
	voucherPermitController := voucherController.NewController(voucherPermitService)

	//transaction
	transactionPermitRepository := transactionRepo.RepositoryFactory(dbCon)
	transactionPermitService := transactionService.NewService(transactionPermitRepository, userPermitRepository)
	transactionPermitController := transactionController.NewController(transactionPermitService)

	redeemPermitRepository := redeemRepo.RepositoryFactory(dbCon)
	redeemPermitService := redeemService.NewService(redeemPermitRepository)
	redeemPermitController := redeemController.NewController(redeemPermitService)

	// Xendit
	xenditPermitService := xenditService.NewService(voucherPermitRepository, redeemPermitRepository, userPermitRepository)
	xenditPermitController := xenditController.NewController(xenditPayload.XenditCallbackPayload{}, xenditPermitService)

	controllers := api.Controller{
		ContentV1Controller:     contentV1PermitController,
		AuthController:          authPermitController,
		UserController:          userPermitController,
		MiddlewareJwt:           middlewarePermitJwt,
		VoucherController:       voucherPermitController,
		TransactionController:   transactionPermitController,
		XenditController:        xenditPermitController,
		RedeemVoucherController: redeemPermitController,
	}

	return controllers
}
