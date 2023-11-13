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
	PetAdd(add reqDto.PetTypeAdd) error
	PetUpd(upd reqDto.PetTypeUpd) error
	PetInfo(id uint) (*resDto.PetTypeInfo, error)

	PetTypeDel(id uint) error
	PetTypeAdd(add reqDto.PetTypeAdd) error
	PetTypeUpd(upd reqDto.PetTypeUpd) error
	PetTypeInfo(id uint) (*resDto.PetTypeInfo, error)
	PetList(list reqDto.PetList) (*resDto.CommonList, error)
}

func (p PetService) PetDel(id uint) error {
	_, err := p.PetInfo(id)
	if err != nil {
		return err
	}
	return petRepository.PetTypeDel(id)
}

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

func (p PetService) PetTypeAdd(add reqDto.PetTypeAdd) error {
	return petRepository.PetTypeAdd(add)
}
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
func (p PetService) PetList(list reqDto.PetList) (*resDto.CommonList, error) {
	resList, err := petRepository.PetList(list)
	if err != nil {
		return nil, err
	}
	return resList, nil
}
