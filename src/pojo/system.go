package pojo

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}

//type RoleInterface interface {
//	FindById(id int) (*Role, error)
//}
//
//func (r *Role) FindById(id int) (*Role, error) {
//	r.ID = uint(id)
//	err := db.First(r).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errors.New(msg.NOT_FOUND_ERROR)
//		}
//		return nil, err
//	}
//	return r, nil
//}
