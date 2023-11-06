package pojo

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	UserName    string    `json:"user_name" gorm:"not null ;comment:'账号'" `
	NickName    string    `json:"nick_name" gorm:"comment:'昵称'"`
	Password    string    `json:"password" gorm:"comment:'密码'"`
	Salt        string    `json:"salt" gorm:"comment:'加密盐'"`
	Phone       string    `json:"phone" gorm:"unique:true;comment:'手机号'"`
	AccessToken string    `json:"access_token" gorm:"comment:'登录状态'"`
	Sex         int       `json:"sex" gorm:"comment:'性别'"`
	Avatar      string    `json:"avatar" gorm:"comment:'头像'"`
	Email       string    `json:"email" gorm:"comment:'邮箱'"`
	Roles       []Role    `gorm:"many2many:account_roles;"`
	LoginIp     string    `json:"login_ip" gorm:"comment:'登录ip';default:null"`
	LoginAt     time.Time `json:"login_at" gorm:"default:null;comment:'登陆时间'"`
	Class       string    `json:"class" gorm:"comment:'身份类别'"`
}

//type AdminRepositoryInter interface {
//	AdminList(list *reqDto.AdminList) (*resDto.CommonList, error)
//	CheckByName(userName string) (*Admin, error)
//	CheckByPhone(phone string) (*Admin, error)
//	CheckByNickName(nickName string) (*Admin, error)
//	RemoveAccessToken(id uint) error
//	UpdateToken(access_token string, id uint, ip string) error
//	AddAdmin(body Admin) error
//}
//
//var (
//	adminInfo    = resDto.AdminInfo{}
//	resAdminList = []resDto.AdminList{} //要查询的字段
//)
//
//// 分页,模糊查询用户
//func (a *Admin) AdminList(list *reqDto.AdminList) (*resDto.CommonList, error) {
//	query := db.Model(&a)
//	if list.Name != "" {
//		query.Where("user_name like ?", "%"+list.Name+"%")
//	}
//	err := query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAdminList).Count(&count).Error
//	reslist.Count = uint(count)
//	reslist.List = resAdminList
//	if err != nil {
//		return nil, err
//	}
//	return &reslist, nil
//}
//
//// 查询账号
//func (a *Admin) CheckByNickName(nickName string) (*Admin, error) {
//	a.NickName = nickName
//	err := db.Preload("Role").First(a).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//	}
//	return a, nil
//}
//
//// 查询手机号
//func (a *Admin) CheckByPhone(phone string) (*Admin, error) {
//	a.Phone = phone
//	err := db.Preload("Role").First(a).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//	}
//	return a, nil
//}
//
//// 查询名称
//func (a *Admin) CheckByName(userName string) (*Admin, error) {
//	fmt.Println("mysql", userName)
//	a.UserName = userName
//	err := db.First(a).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//
//	}
//	return a, nil
//}
//
//// 详情数据
//func (a *Admin) AdminInfo(id uint) (*resDto.AdminInfo, error) {
//	a.ID = id
//	err := db.Model(&a).Select("inter.*,r.name as role_name").Joins("left join rule as r on r.id = inter.role").Scan(&adminInfo).Error
//	if err != nil {
//		return nil, err
//	}
//	return &adminInfo, nil
//}
//
//// 更新token数据
//func (a *Admin) UpdateToken(access_token string, id uint, ip string) error {
//	a.ID = id
//	err := db.Model(&a).Updates(Admin{AccessToken: access_token, LoginAt: time.Now(), LoginIp: ip}).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (a *Admin) RemoveAccessToken(id uint) error {
//	a.ID = id
//	return db.Model(&a).Update("access_token", "").Error
//
//}
//
//// 增加用户
//func (a *Admin) AddAdmin(account Admin) error {
//	return db.Create(&account).Error
//}
