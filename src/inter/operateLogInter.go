package inter

import (
	"errors"
	"gorm.io/gorm"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/pojo"
)

type OperateLogInter interface {
	AddOperateLog(operation *pojo.OperationLog) error
}

type OperateResourceInter interface {
	OperateResourceDel(id uint) error
	OperateResourceAdd(add *pojo.OperateResource) error
	OperateResourceUpd(upd *reqDto.OperateResourceUpd) error
	OperateResourceInfo(id uint) (*resDto.OperateResourceInfo, error)
	OperateResourceList(list reqDto.OperateResourceList) (*resDto.CommonList, error)
}
type OperateLogImpl struct{}
type OperateResourceImpl struct{}

var (
	operateResource pojo.OperateResource
)

func (o OperateLogImpl) AddOperateLog(operation *pojo.OperationLog) error {
	return db.Create(&operation).Error
}

func (op OperateResourceImpl) OperateResourceDel(id uint) error {
	operateResource.ID = id
	return db.Delete(&operateResource).Error
}
func (op OperateResourceImpl) OperateResourceAdd(add *pojo.OperateResource) error {
	return db.Create(&add).Error
}

func (op OperateResourceImpl) OperateResourceUpd(upd *reqDto.OperateResourceUpd) error {
	operateResource.ID = upd.Id
	return db.Model(&operateResource).Updates(&upd).Error
}

func (op OperateResourceImpl) OperateResourceInfo(id uint) (*resDto.OperateResourceInfo, error) {
	var operateInfo resDto.OperateResourceInfo
	operateResource.ID = id
	err := db.Model(&operateResource).Find(&operateInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &operateInfo, nil
}
func (op OperateResourceImpl) OperateResourceList(list reqDto.OperateResourceList) (*resDto.CommonList, error) {
	var query = db.Model(&operateResource)
	if len(list.Name) > 0 {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	if len(list.ResourceMethod) > 0 {
		query.Where("method = ?", list.ResourceMethod)
	}
	var operateResourceInfo []resDto.OperateResourceInfo
	err := query.Limit(list.Take).Offset(list.Skip).Find(&operateResourceInfo).Count(&count).Error
	if err != nil {
		return nil, err
	}
	reslist.List = operateResourceInfo
	reslist.Count = uint(count)
	return &reslist, nil
}
