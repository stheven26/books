package infra

import (
	"fmt"
	"test/internal/domain/entity"
	"test/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	config := config.LoadEnv()
	envData := map[string]string{
		"username": config.GetString("DB_USERNAME"),
		"password": config.GetString("DB_PASSWORD"),
		"host":     config.GetString("DB_HOST"),
		"port":     config.GetString("DB_PORT"),
		"name":     config.GetString("DB_NAME"),
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envData["username"], envData["password"], envData["host"], envData["port"], envData["name"],
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&entity.Book{})
	return db
}
