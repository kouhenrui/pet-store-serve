package pojo

import (
	"gorm.io/gorm"
	"log"
)

func Repositoryinit(db *gorm.DB) {
	err := db.AutoMigrate(
		&Account{},      //管理员表
		&Role{},         //角色表
		&OperationLog{}, //操作记录表
		&App{},          //app记录表
	)
	if err != nil {
		panic(err)
	}
	log.Println("结构表创建成功")
}
