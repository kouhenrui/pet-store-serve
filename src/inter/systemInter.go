package inter

import (
	"errors"
	"gorm.io/gorm"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/msg"
	"pet-store-serve/src/pojo"
)

var (
	r = &pojo.Role{}
	m pojo.Menu
)

type SysInter interface {
	FindById(id int) (*pojo.Role, error)
	FindByName(name string) (*pojo.Role, error)

	MenuDel(id uint) error
	MenuAdd(add *pojo.Menu) error
	MenuUpd(upd reqDto.MenuUpd) error
	GetMenuList() (*resDto.CommonList, error)
	MenuInfo(id uint) (*resDto.MenuInfo, error)
	MenuList(list reqDto.MenuList) (*resDto.CommonList, error)
}
type SysRepositoryImpl struct{}

// TODO 通过ID查找
func (sp SysRepositoryImpl) FindById(id int) (*pojo.Role, error) {
	r.ID = uint(id)
	err := db.First(&r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r, nil
}

// TODO 通过name查找
func (sp SysRepositoryImpl) FindByName(name string) (*pojo.Role, error) {
	r.Name = name
	err := db.First(&r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(msg.NOT_FOUND_ERROR)
		}
		return nil, err
	}
	return r, nil
}

func (sp SysRepositoryImpl) MenuList(list reqDto.MenuList) (*resDto.CommonList, error) {
	var query = db.Model(&m)
	var menuList []resDto.MenuInfo
	if list.Name != "" {
		query.Where("menu_name like ?", "%"+list.Name+"%")
	}
	if list.ParentName != "" {
		query.Where("parent_name like ?", "%"+list.ParentName+"%")
	}
	err := query.Find(&menuList).Limit(list.Take).Offset(list.Skip).Count(&count).Error
	if err != nil {
		return nil, err
	}
	reslist.List = menuList
	reslist.Count = uint(count)
	return &reslist, nil
}
func (sp SysRepositoryImpl) GetMenuList() (*resDto.CommonList, error) {
	var query = db.Model(&m)
	var menuList []resDto.MenuInfo
	err := query.Find(&menuList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	reslist.List = menuList
	reslist.Count = uint(count)
	return &reslist, nil
}
func (sp SysRepositoryImpl) MenuInfo(id uint) (*resDto.MenuInfo, error) {
	m.ID = id
	var menuInfo resDto.MenuInfo
	err := db.Model(&m).First(&menuInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &menuInfo, nil
}

func (sp SysRepositoryImpl) MenuDel(id uint) error {
	m.ID = id
	return db.Delete(&m).Error
}

func (sp SysRepositoryImpl) MenuUpd(upd reqDto.MenuUpd) error {
	m.ID = upd.Id
	return db.Model(&m).Updates(&upd).Error
}
func (sp SysRepositoryImpl) MenuAdd(add *pojo.Menu) error {
	return db.Model(&m).Save(&add).Error
}
