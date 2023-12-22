package sys

import (
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/inter"
)

var (
	sysRepository inter.SysInter = &inter.SysRepositoryImpl{}
)

type SysService struct{}
type SysInter interface {
	MenuList(list reqDto.MenuList) (*resDto.CommonList, error)
	GetMenuList() (*resDto.CommonList, error)
}

func (s SysService) MenuList(list reqDto.MenuList) (*resDto.CommonList, error) {
	return sysRepository.MenuList(list)
}
func (s SysService) GetMenuList() (*resDto.CommonList, error) {
	return sysRepository.GetMenuList()
}
