package connection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB){
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/bankgateway?charset=utf8&parseTime=True&loc=Local"),&gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connection database ::: %s",err.Error())
	}

	return db
}
