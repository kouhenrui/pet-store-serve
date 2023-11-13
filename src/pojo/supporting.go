package pojo

import "gorm.io/gorm"

type Supporting struct {
	gorm.Model
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	SupportingType SupportingType `json:"supporting_type"`
	Number         int            `json:"number"`
	Price          float32        `json:"price"`
	Picture        string         `json:"picture"`      //轮播显示图片
	ShowPicture    string         `json:"show_picture"` //详情页展示图片
}

type SupportingType struct {
	gorm.Model      `json:"gorm.Model"`
	Name            string
	SupportingClass []PetClass `gorm:"foreignkey:type_id"`
}

// 细分小类
type SupportingClass struct {
	gorm.Model `json:"gorm.Model"`
	Name       string `json:"name"`
	TypeId     uint   ` json:"type_id"`
}
