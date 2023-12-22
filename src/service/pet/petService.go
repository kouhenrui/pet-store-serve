package pet

import (
	"errors"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	inter2 "pet-store-serve/src/inter"
	"pet-store-serve/src/msg"
)

var (
	petRepository inter2.PetRepositoryInter = &inter2.PetRepositoryImpl{}
)

type PetService struct{}

type PetInter interface {
	PetDel(id uint) error
	PetAdd(add reqDto.PetAdd) error
	PetUpd(upd reqDto.PetUpd) error
	PetInfo(id uint) (*resDto.PetInfo, error)
	PetList(list reqDto.PetList) (*resDto.CommonList, error)

	PetTypeDel(id uint) error
	PetTypeAdd(add reqDto.PetTypeAdd) error
	PetTypeUpd(upd reqDto.PetTypeUpd) error
	PetTypeInfo(id uint) (*resDto.PetTypeInfo, error)
	PetTypeList(list reqDto.PetTypeList) (*resDto.CommonList, error)
}

// TODO
func (p PetService) PetDel(id uint) error {
	_, err := p.PetInfo(id)
	if err != nil {
		return err
	}
	return petRepository.PetTypeDel(id)
}

// TODO
func (p PetService) PetAdd(add reqDto.PetAdd) error {
	return petRepository.PetAdd(add)
}

// TODO
func (p PetService) PetUpd(upd reqDto.PetUpd) error {
	if _, err := p.PetInfo(upd.Id); err != nil {
		return err
	}
	return petRepository.PetUpdate(upd)
}

// TODO
func (p PetService) PetInfo(id uint) (*resDto.PetInfo, error) {
	petInfo, err := petRepository.PetInfo(id)
	if err != nil {
		return nil, err
	}
	if petInfo == nil && err == nil {
		return nil, errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petInfo, nil
}

// TODO
func (p PetService) PetTypeDel(id uint) error {
	info, err := petRepository.PetTypeInfo(id)
	if err != nil {
		return errors.New(msg.INTERNAL_ERROR)
	}
	if info == nil && err == nil {
		return errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petRepository.PetTypeDel(id)
}

// TODO
func (p PetService) PetTypeAdd(add reqDto.PetTypeAdd) error {
	return petRepository.PetTypeAdd(add)
}

// TODO
func (p PetService) PetTypeUpd(upd reqDto.PetTypeUpd) error {
	info, err := petRepository.PetTypeInfo(upd.Id)
	if err != nil {
		return errors.New(msg.INTERNAL_ERROR)
	}
	if info == nil && err == nil {
		return errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petRepository.PetTypeUpd(upd)
}

// TODO
func (p PetService) PetTypeInfo(id uint) (*resDto.PetTypeInfo, error) {
	petTypeInfo, err := petRepository.PetTypeInfo(id)
	if err != nil {
		return nil, err
	}
	if petTypeInfo == nil && err == nil {
		return nil, errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petTypeInfo, nil
}

// TODO pet列表
func (p PetService) PetList(list reqDto.PetList) (*resDto.CommonList, error) {
	resList, err := petRepository.PetList(list)
	if err != nil {
		return nil, err
	}
	return resList, nil
}

// TODO pettype 列表
func (p PetService) PetTypeList(list reqDto.PetTypeList) (*resDto.CommonList, error) {
	resList, err := petRepository.PetTypeList(list)
	if err != nil {
		return nil, err
	}
	return resList, nil
}

// TODO petclass列表
func (p PetService) PetClassList(list reqDto.PetClassList) (*resDto.CommonList, error) {
	resList, err := petRepository.PetClassList(list)
	if err != nil {
		return nil, err
	}
	return resList, nil
}

// TODO petclass详情
func (p PetService) PetClassInfo(id uint) (*resDto.PetClassInfo, error) {
	petTClassInfo, err := petRepository.PetClassInfo(id)
	if err != nil {
		return nil, err
	}
	if petTClassInfo == nil && err == nil {
		return nil, errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petTClassInfo, nil
}

// TODO petclass删除
func (p PetService) PetClassDel(id uint) error {
	return petRepository.PetClassDel(id)
}

// TODO petclass修改
func (p PetService) PetClassUpd(upd reqDto.PetClassUpd) error {
	info, err := petRepository.PetClassInfo(upd.Id)
	if err != nil {
		return errors.New(msg.INTERNAL_ERROR)
	}
	if info == nil && err == nil {
		return errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	return petRepository.PetClassUpd(upd)
}

// TODO petclass增加,需要验证名称是否存在
func (p PetService) PetClassAdd(add reqDto.PetClassAdd) error {
	return petRepository.PetClassAdd(add)
}
