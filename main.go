package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"nochat/app/config"
	database2 "nochat/app/src/Database"
	"nochat/app/src/Routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&database2.Contact{})
	config.DB.AutoMigrate(&database2.Advertisement{})

	r := Routes.SetupRouter()

	err := r.Run()

	if err != nil {
		return
	}
}
