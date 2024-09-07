package Database

import (
	"New_Book_Management/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:Mhdw2004@8313@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = Db.AutoMigrate(&Models.Book{})
	if err != nil {
		return
	}
}
