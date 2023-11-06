package inter

import (
	"errors"
	"gorm.io/gorm"
	"pet-store-serve/src/msg"
	"pet-store-serve/src/pojo"
)

var (
	r = &pojo.Role{}
)

type RoleInter interface {
	FindById(id int) (*pojo.Role, error)
	FindByName(name string) (*pojo.Role, error)
}
type RoleRepositoryImpl struct{}

// TODO 通过ID查找
func (rp RoleRepositoryImpl) FindById(id int) (*pojo.Role, error) {
	r.ID = uint(id)
	err := db.First(r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r, nil
}

// TODO 通过name查找
func (rp RoleRepositoryImpl) FindByName(name string) (*pojo.Role, error) {
	r.Name = name
	err := db.First(r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r, nil
}
