package inter

import (
	"errors"
	"gorm.io/gorm"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/global"
	"pet-store-serve/src/pojo"
	"time"
)

var (
	db             = global.MysqlDClient
	a              = &pojo.Account{}
	resAccountList []resDto.AccountList
	count          int64
	reslist        resDto.CommonList
)

type AccountRepositoryImpl struct{}

type AccountRepositoryInter interface {
	AccountList(list reqDto.AccountList) (*resDto.CommonList, error)
	CheckByName(userName string) (*pojo.Account, error)
	CheckByPhone(phone string) (*pojo.Account, error)
	CheckByNickName(nickName string) (*pojo.Account, error)
	RemoveAccessToken(id uint) error
	UpdateToken(access_token string, id uint, ip string) error
	AddAccount(body pojo.Account) error
	AccountInfo(id uint) (*pojo.Account, error)
	ResetPwdBySelf(id uint, pwd string) error
}

// TODO 根据手机号查询
func (ap AccountRepositoryImpl) CheckByPhone(phone string) (*pojo.Account, error) {
	a.Phone = phone
	err := db.Preload("Roles").Find(&a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return a, nil
}

// TODO 列表
func (ap AccountRepositoryImpl) AccountList(list reqDto.AccountList) (*resDto.CommonList, error) {
	query := db.Model(a)
	if list.Name != "" {
		query.Where("user_name like ?", "%"+list.Name+"%")
	}
	err := query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAccountList).Count(&count).Error
	reslist.Count = uint(count)
	reslist.List = resAccountList
	if err != nil {
		return nil, err
	}
	return &reslist, nil
}

// TODO 查询账号
func (ap AccountRepositoryImpl) CheckByNickName(nickName string) (*pojo.Account, error) {
	a.NickName = nickName
	err := db.Preload("Roles").First(a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return a, nil
}

// TODO 查询昵称
func (ap AccountRepositoryImpl) CheckByName(userName string) (*pojo.Account, error) {
	a.UserName = userName
	err := db.Preload("Roles").First(&a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err

	}
	return a, nil
}

// TODO 详情数据
func (ap AccountRepositoryImpl) AccountInfo(id uint) (*pojo.Account, error) {
	a.ID = id
	err := db.Preload("Roles").First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// TODO 更新token数据
func (ap AccountRepositoryImpl) UpdateToken(access_token string, id uint, ip string) error {
	a.ID = id
	err := db.Model(a).Updates(pojo.Account{AccessToken: access_token, LoginAt: time.Now(), LoginIp: ip}).Error
	if err != nil {
		return err
	}
	return nil
}

// TODO 去除token
func (ap AccountRepositoryImpl) RemoveAccessToken(id uint) error {
	a.ID = id
	return db.Model(&a).Update("access_token", "").Error

}

// TODO 增加用户
func (ap AccountRepositoryImpl) AddAccount(account pojo.Account) error {
	return db.Create(&account).Error
}

// TODO 修改密码
func (ap AccountRepositoryImpl) ResetPwdBySelf(id uint, pwd string) error {
	a.ID = id
	return db.Model(&a).Update("password", pwd).Error
}
