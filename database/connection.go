package database

import (
	"todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB{
	dbConnectionDetails := "host=localhost user=postgres password=admin dbname=todoGolang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection,err:=gorm.Open(postgres.Open(dbConnectionDetails),&gorm.Config{})
	if(err!=nil){
		panic("Sorry! could not connect to database, checck your connection")
	}
	DB=connection
	connection.AutoMigrate(&models.User{},&models.Todo{},&models.Tag{})
	return DB
}