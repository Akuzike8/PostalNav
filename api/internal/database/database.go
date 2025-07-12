package database

import (
	//"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	//"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var dbInstance *gorm.DB

func Connect() *gorm.DB {
	once.Do(func() {
	config := config.LoadConfig()
	dbuser := config.DBUSER
    dbpass := config.DBPASS
    dbhost := config.DBHOST
    dbport := config.DBPORT
    dbname := config.DBNAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",dbuser,dbpass,dbhost,dbport)
	serverdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)

	if err := serverdb.Exec(createDB).Error; err != nil {
		log.Fatal("failed to create database:", err)
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbuser,dbpass,dbhost,dbport,dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true,})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB,_ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db
	})

	return dbInstance

}

/*func toJSON(data []string) datatypes.JSON {
    jsonData, _ := json.Marshal(data)
    return datatypes.JSON(jsonData)
}*/

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{},&models.Permission{},&models.Audit{},&models.Bank{},&models.Merchant_Bank{},&models.Merchant_User{},&models.Merchant{},&models.Merchant_User_Role{},&models.Role{},&models.Role_Permission{},&models.Api_key{},&models.Country{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}


}
