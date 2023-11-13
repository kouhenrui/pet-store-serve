package pojo

import (
	"gorm.io/gorm"
)

// 宠物大类，猫狗等
type PetType struct {
	gorm.Model `json:"gorm.Model"`
	Name       string
	PetClass   []PetClass `gorm:"foreignkey:type_id"`
}

// 细分小类
type PetClass struct {
	gorm.Model `json:"gorm.Model"`
	Name       string `json:"name"`
	TypeId     uint   ` json:"type_id"`
}

// 宠物表
type Pet struct {
	gorm.Model  `json:"gorm.Model"`
	Title       string   `json:"title"`
	Age         string   `json:"age"`
	Description string   `json:"description"`
	Number      int      `json:"number"`
	Price       float32  `json:"price"`
	Picture     string   `json:"picture"` //轮播显示图片
	ShowPicture string   `json:"show_picture"`
	TypeId      uint     `json:"type_id"`  //绑定PetType
	ClassId     uint     `json:"class_id"` //绑定PetClass
	Type        PetType  `gorm:"foreignKey:type_id"`
	Class       PetClass `gorm:"foreignKey:class_id"`
}
