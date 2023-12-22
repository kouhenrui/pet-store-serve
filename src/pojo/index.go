package pojo

import (
	"gorm.io/gorm"
	"log"
)

func Repositoryinit(db *gorm.DB) {
	err := db.AutoMigrate(
		&Account{},  //账号表
		&Role{},     //角色表
		&Class{},    //身份类别表
		&Pet{},      //宠物表
		&PetClass{}, //宠物大类
		&PetType{},  //宠物小类
		&Menu{},
		&OperationLog{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("结构表创建成功")
}

type PermissiobType string

const (
	a PermissiobType = "admin"

	u PermissiobType = "user"
)
