package util

import (
	"dapoint-api/config"
	"dapoint-api/entities"
	"fmt"
	// "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	Postgres DatabaseDriver = "postgres"
	Mysql    DatabaseDriver = "mysql"
	Static   DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver   DatabaseDriver
	Mysql    *gorm.DB
	Postgres *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "postgres":
		db.Driver = Postgres
		db.Postgres = newPostgres(config)
	case "static":
		db.Driver = Static
	case "mysql":
		db.Driver = Mysql
		db.Mysql = newMysql(config)
	default:
		panic("unsupported driver")
	}

	return &db
}

func newPostgres(config *config.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.Database.Address,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Transaction{}, &entities.Voucher{}, &entities.UserVoucher{})
	if err != nil {
		return nil
	}

	return db

}

func newMysql(config *config.AppConfig) *gorm.DB {
	// root:@tcp(localhost:3306)/poseidon
	dbURL := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)
	// dbURL := config.Database.DBURL
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Transaction{}, &entities.Voucher{}, &entities.UserVoucher{})
	if err != nil {
		return nil
	}

	return db

}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
