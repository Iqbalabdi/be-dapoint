package api

import (
	"dapoint-api/api/middleware"
	auth "dapoint-api/api/v1/auth"
	contentV1 "dapoint-api/api/v1/content"
	transactionController "dapoint-api/api/v1/transaction"
	"dapoint-api/api/v1/user"
	userVoucherController "dapoint-api/api/v1/user_voucher"
	voucherController "dapoint-api/api/v1/voucher"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentV1Controller   *contentV1.Controller
	UserController        *user.Controller
	AuthController        *auth.Controller
	MiddlewareJwt         middleware.JWTService
	VoucherController     *voucherController.Controller
	UserVoucherController *userVoucherController.Controller
	TransactionController *transactionController.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	// contentV1.Use(middleware.JWTMiddleware())
	contentV1.GET("", controller.ContentV1Controller.GetAll)
	contentV1.POST("", controller.ContentV1Controller.Create)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)

	user := e.Group("/user")
	user.POST("/login", controller.UserController.Login)     // login
	user.GET("/getall", controller.UserController.GetAll)    // get all users
	user.POST("/register", controller.UserController.Create) // register users
	user.GET("/:id", controller.UserController.GetByID)      // get users by id
	user.PUT("/:id", controller.UserController.Modify)       // update users
	// TODO Delete users
	user.DELETE("/:id", controller.UserController.Delete)
	user.POST("/user_voucher", controller.UserVoucherController.Redeem)
	user.GET("/user_transaction/:userid", controller.TransactionController.GetByUserID)

	// TODO Create vouchers
	voucher := e.Group("/vouchers")
	voucher.POST("/create", controller.VoucherController.Create)
	voucher.GET("/getall", controller.VoucherController.GetAll)
	// TODO Update Vouchers by id
	voucher.PUT("/update/:id", controller.VoucherController.Modify)
	voucher.GET("/getbyid/:id", controller.VoucherController.GetByID)
	voucher.GET("/getbytype/:tipe", controller.VoucherController.GetByParams)
	// TODO Delete vouchers

	// TODO Create user transactions
	//transaction := e.Group("/transaction")
	// TODO GET All user transactions

	// Admin
	admin := e.Group("/admin")
	// TODO Admin manage point customer
	admin.POST("/login", controller.UserController.Login)
	admin.PUT("/user_point/:id", controller.UserController.PointModify)
	admin.PUT("/user/:id", controller.UserController.Modify)
	admin.GET("/user/getall", controller.UserController.GetAll)
	admin.GET("/user/:id", controller.UserController.GetByID)
	admin.DELETE("/user/:id", controller.UserController.Delete)
	admin.GET("/voucher/getall", controller.VoucherController.GetAll)
	admin.GET("/voucher/:id", controller.VoucherController.GetByID)
	admin.POST("/voucher/create", controller.VoucherController.Create)
	admin.DELETE("/voucher/:id", controller.VoucherController.Delete)
	// transaction
	admin.POST("/user_transaction/create", controller.TransactionController.Create)
	admin.GET("/user_transaction/getall", controller.TransactionController.GetAll)
	admin.GET("/user_transaction/getbyuserid/:userid", controller.TransactionController.GetByUserID)
}
