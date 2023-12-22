package inter

import (
	"errors"
	"gorm.io/gorm"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/pojo"
	util "pet-store-serve/src/utils"
	"strconv"
)

type PetRepositoryImpl struct{}
type PetRepositoryInter interface {
	PetAdd(add reqDto.PetAdd) error
	PetList(list reqDto.PetList) (*resDto.CommonList, error)
	PetInfo(id uint) (*resDto.PetInfo, error)
	PetUpdate(upd reqDto.PetUpd) error
	PetDelete(id uint) error
	PetTypeDel(id uint) error
	PetTypeAdd(add reqDto.PetTypeAdd) error
	PetTypeUpd(upd reqDto.PetTypeUpd) error
	PetTypeInfo(id uint) (*resDto.PetTypeInfo, error)
	PetTypeList(list reqDto.PetTypeList) (*resDto.CommonList, error)

	PetClassDel(id uint) error
	PetClassAdd(add reqDto.PetClassAdd) error
	PetClassUpd(upd reqDto.PetClassUpd) error
	PetClassInfo(id uint) (*resDto.PetClassInfo, error)
	PetClassList(list reqDto.PetClassList) (*resDto.CommonList, error)

	PetClassFindByTypeIdName(name string, typeId uint) (*resDto.PetClassInfo, error)
}

var (
	pet      pojo.Pet
	petType  pojo.PetType
	petClass pojo.PetClass

	petList      = []resDto.PetInfo{}
	petTypeList  = []resDto.PetTypeInfo{}
	petClassList = []resDto.PetClassInfo{}
)

// TODO pet 增加
func (pl PetRepositoryImpl) PetAdd(add reqDto.PetAdd) error {
	return db.Create(&add).Error
}

// TODO pet列表
func (pl PetRepositoryImpl) PetList(list reqDto.PetList) (*resDto.CommonList, error) {
	query := db.Model(&pet)
	if list.Title != "" {
		query.Where("title like ?", "%"+list.Title+"%")
	}
	if list.TypeId != "" {
		query.Where("type_id", list.TypeId)
	}
	if list.ClassId != "" {
		query.Where("class_id", list.ClassId)
	}
	err := query.Limit(list.Take).Offset(int(list.Skip)).Find(&petList).Count(&count).Error
	reslist.Count = uint(count)
	reslist.List = petList
	if err != nil {
		return nil, err
	}
	return &reslist, nil
}

// TODO pet详情
func (pl PetRepositoryImpl) PetInfo(id uint) (*resDto.PetInfo, error) {
	var petInfo resDto.PetInfo
	pet.ID = id
	err := db.Model(&pet).Find(&petInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if len(strconv.Itoa(int(petInfo.TypeId))) > 0 {
		if typeName, _ := pl.PetTypeInfo(petInfo.TypeId); typeName != nil {
			petInfo.TypeName = typeName.Name
		}
	}
	if len(strconv.Itoa(int(petInfo.ClassId))) > 0 {
		if className, _ := pl.PetClassInfo(petInfo.ClassId); className != nil {
			petInfo.ClassName = className.Name
		}
	}

	return &petInfo, nil
}

// TODO petType详情
func (pl PetRepositoryImpl) PetTypeInfo(id uint) (*resDto.PetTypeInfo, error) {
	var petTypeInfo resDto.PetTypeInfo
	petType.ID = id
	err := db.Model(&petType).Find(&petTypeInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &petTypeInfo, nil
}

// TODO petclass 详情
func (pl PetRepositoryImpl) PetClassInfo(id uint) (*resDto.PetClassInfo, error) {
	var petClassInfo resDto.PetClassInfo
	petClass.ID = id
	err := db.Model(&petClass).Preload("PetTypes").Find(&petClassInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &petClassInfo, nil
}

// TODO pet修改
func (pl PetRepositoryImpl) PetUpdate(upd reqDto.PetUpd) error {
	pet.ID = upd.Id
	return db.Model(&pet).Updates(&upd).Error
}

// TODO pet根据id删除
func (pl PetRepositoryImpl) PetDelete(id uint) error {
	pet.ID = id
	return db.Delete(&pet).Error
}

// TODO petType add
func (pl PetRepositoryImpl) PetTypeAdd(add reqDto.PetTypeAdd) error {
	return db.Create(&add).Error
}

// TODO petType list
func (pl PetRepositoryImpl) PetTypeList(list reqDto.PetTypeList) (*resDto.CommonList, error) {
	query := db.Model(&petType)
	if list.Name != "" {
		query.Where("name", list.Name)
	}
	err := query.Limit(list.Take).Offset(list.Skip).Find(&petTypeList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	reslist.List = petTypeList
	reslist.Count = uint(count)
	return &reslist, nil
}

// TODO petType修改
func (pl PetRepositoryImpl) PetTypeUpd(upd reqDto.PetTypeUpd) error {
	return db.Model(&petType).Updates(&upd).Error
}

// TODO petType删除
func (pl PetRepositoryImpl) PetTypeDel(id uint) error {
	petType.ID = id
	return db.Delete(&petType).Error
}

// TODO petclass add
func (pl PetRepositoryImpl) PetClassAdd(add reqDto.PetClassAdd) error {
	return db.Model(&petClass).Create(&add).Error
}

// TODO petClass列表
func (pl PetRepositoryImpl) PetClassList(list reqDto.PetClassList) (*resDto.CommonList, error) {
	query := db.Model(&petClass)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}

	if util.IsFieldEmpty(list, strconv.Itoa(int(list.TypeId))) {
		query.Where("type_id = ?", list.TypeId)
	}
	err := query.Limit(list.Take).Offset(list.Skip).Find(&petClassList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	reslist.List = petClassList
	reslist.Count = uint(count)
	return &reslist, nil
}

// TODO petClass删除
func (pl PetRepositoryImpl) PetClassDel(id uint) error {
	petClass.ID = id
	return db.Delete(&petClass).Error
}

func (pl PetRepositoryImpl) PetClassUpd(upd reqDto.PetClassUpd) error {
	petClass.ID = upd.Id
	return db.Model(&petClass).Updates(&upd).Error
}

func (pl PetRepositoryImpl) PetClassFindByTypeIdName(name string, typeId uint) (*resDto.PetClassInfo, error) {
	var petClassInfo resDto.PetClassInfo
	petClass.TypeId = typeId
	petClass.Name = name
	err := db.Model(&petClass).Preload("PetTypes").Find(&petClassInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &petClassInfo, nil
}
