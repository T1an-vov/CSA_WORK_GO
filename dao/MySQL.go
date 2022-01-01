package dao

import (
	"CSA_Work_Go/module"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func SQLInit()(err error) {
	db, err := gorm.Open("mysql", "root:root@/csa?charset=utf8mb4")
	if err != nil {
		return err
	}
	err=db.DB().Ping()
	if err != nil {
		return err
	}
	db.AutoMigrate(&module.Student{})
	db.AutoMigrate(&module.Teacher{})
	db.AutoMigrate(&module.Record{})
	DB=db
	return nil
}
func Close() () {
	DB.Close()
}