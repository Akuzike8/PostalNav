package DB

import(
	"os"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB{
	dbuser := os.Getenv("DBUSER")
    dbpass := os.Getenv("DBPASS")
    dbhost := os.Getenv("DBHOST")
    dbport := os.Getenv("DBPORT")
    dbname := os.Getenv("DBNAME")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbuser,dbpass,dbhost,dbport,dbname)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
        log.Fatal(err)
    }
	return db
}