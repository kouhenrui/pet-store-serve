package account

import (
	"errors"
	"fmt"
	"log"
	"pet-store-serve/src/dto/comDto"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/global"
	inter2 "pet-store-serve/src/inter"
	"pet-store-serve/src/msg"
	"pet-store-serve/src/pojo"
	"pet-store-serve/src/service/common"
	util "pet-store-serve/src/utils"
	"time"
)

var (
	account           *pojo.Account
	err               error
	commonService     common.CommonServiceImpl
	accountRepository inter2.AccountRepositoryInter = &inter2.AccountRepositoryImpl{}
	sysRepository     inter2.SysInter               = &inter2.SysRepositoryImpl{}
	jwtService        util.JwtService
)

type AccountInter interface {
	Login(body reqDto.AccountLogin, ip string) (resDto.TokenAndExp, error)
	Register(body reqDto.AddAccount) error
	Info(id uint) (*resDto.AccountInfo, error)
	ResetPwdBySelf(body reqDto.UpdateAccount) error
	RefreshToken(oldToken, ip string) (resDto.TokenAndExp, error)
}
type AccountService struct{}

// TODO 登录方法
func (i AccountService) Login(body reqDto.AccountLogin, ip string) (t resDto.TokenAndExp, err error) {
	//captcha := &reqDto.Captcha{
	//	Id:      body.Uuid,
	//	Content: body.Code,
	//}
	//if err = global.ExistRedis(captcha.Id); err == nil {
	//	return t, errors.New(msg.CAPTCHA_ERROR)
	//}
	//if !commonService.VfCaptcha(*captcha) {
	//	return t, errors.New(msg.CAPTCHA_CREATE_ERROR)
	//}
	//判定登陆方式
	switch body.Method {
	case "name":
		account, err = accountRepository.CheckByName(body.UserName)
		break

	case "phone":
		account, err = accountRepository.CheckByPhone(body.Phone)
		break
	default:
		return t, errors.New(msg.AUTH_LOGIN_ERROR)
	}
	//将密码解密
	dePwd, err := util.DePwdCode(account.Password, account.Salt)
	if err != nil {
		return t, err
	}
	//解密密码和输入密码比对
	if dePwd != body.Password {
		return t, errors.New(msg.ACCOUNT_PWD_ERROR)
	}
	//查询redis缓存是否存在
	redisErr := global.ExistRedis(account.AccessToken)
	tokenDate := comDto.TokenClaims{
		Id:       account.ID,
		Phone:    account.Phone,
		Name:     account.UserName,
		NickName: account.NickName,
		Role:     account.Roles,
		Class:    account.Class,
	}

	if len(account.AccessToken) > 0 {
		if redisErr {
			tokenValue := global.GetRedis(account.AccessToken)
			util.UnMarshal([]byte(tokenValue), &t)
		} else {

			t, err = i.tokenRedis(tokenDate, ip)
			fmt.Println(t, "t1*****************")
		}
	} else {
		t, err = i.tokenRedis(tokenDate, ip)
	}

	return t, nil
}

// TODO info
func (i AccountService) Info(id uint) (*resDto.AccountInfo, error) {
	var accountInfo *pojo.Account
	accountInfo, err = accountRepository.AccountInfo(id)
	var info = &resDto.AccountInfo{
		Id:       accountInfo.ID,
		UserName: accountInfo.UserName,
		NickName: accountInfo.NickName,
		Phone:    accountInfo.Phone,
		Sex:      accountInfo.Sex,
		Avatar:   accountInfo.Avatar,
		Class:    accountInfo.Class,
		Email:    accountInfo.Email}
	for _, i := range accountInfo.Roles {
		r := i.Name
		info.Role = append(info.Role, r)
	}
	return info, err
}

// TODO token生成，存入缓存
func (i AccountService) tokenRedis(tokenDate comDto.TokenClaims, ip string) (t resDto.TokenAndExp, err error) {
	var tt time.Duration
	switch tokenDate.Class {
	case "admin":
		tt = global.AdminExp
		break
	case "user":
		tt = global.UserExp
		break
	}
	t = jwtService.SignToken(tokenDate, tt)
	//生成access_token
	access_token := util.Rand6String()
	//redis缓存token
	if err = global.SetRedis(access_token, util.Marshal(t), tt); err != nil {
		return t, err
	}
	//数据库保存access_token
	if err = accountRepository.UpdateToken(access_token, account.ID, ip); err != nil {
		return t, err
	}
	return t, nil
}
func (i AccountService) refreshTokenRedis(tokenDate comDto.TokenClaims, ip, accessToken string) (t resDto.TokenAndExp, err error) {
	var tt time.Duration
	switch tokenDate.Class {
	case "admin":
		tt = global.AdminExp
		break
	case "user":
		tt = global.UserExp
		break
	}
	t = jwtService.SignToken(tokenDate, tt)
	//redis缓存token
	if err = global.SetRedis(accessToken, util.Marshal(t), tt); err != nil {
		return t, err
	}
	//数据库保存access_token
	if err = accountRepository.UpdateToken(accessToken, account.ID, ip); err != nil {
		return t, err
	}
	return t, nil
}

// TODO 添加数据
func (i AccountService) Register(body reqDto.AddAccount) error {
	if len(body.UserName) < 0 && len(body.Phone) < 0 {
		return errors.New(msg.ACCOUNT_PHONE_NOT_NULL)
	}
	if len(body.UserName) > 0 {
		fmt.Println(body.UserName)
		_, err = accountRepository.CheckByName(body.UserName)
		if err != nil {
			return err
		}
	}
	if len(body.Phone) > 0 {
		_, err = accountRepository.CheckByPhone(body.Phone)
		if err != nil {
			return err
		}
	}
	var salt = util.MakeSalt()
	var enpwd, _ = util.EnPwdCode(body.Password, salt)
	var roles = []pojo.Role{}
	if len(body.Role) > 0 {
		for _, id := range body.Role {
			var singleRole pojo.Role
			findRole, signErr := sysRepository.FindById(id)
			if signErr != nil {
				return err
			}
			singleRole.Model = findRole.Model
			singleRole.Name = findRole.Name
			roles = append(roles, singleRole)

		}
	}
	var newAccount = pojo.Account{
		UserName: body.UserName,
		NickName: body.NickName,
		Password: enpwd,
		Salt:     salt,
		Phone:    body.Phone,
		Sex:      body.Sex,
		Avatar:   body.Avatar,
		Email:    body.Email,
		Class:    body.Class,
		Roles:    roles,
	}
	log.Println("打印要插入的数据", newAccount)
	return accountRepository.AddAccount(newAccount)
}

// TODO 自己修改密码
func (i AccountService) ResetPwdBySelf(body reqDto.UpdateAccount) error {
	if body.Pwd2 != body.Password {
		return errors.New(msg.TWO_PWD_MATCH_ERROR)
	}
	info, err := accountRepository.AccountInfo(body.Id)
	if err != nil {
		return errors.New(msg.SQL_NOT_EXIT_ERROR)
	}
	dePd, _ := util.DePwdCode(info.Password, info.Salt)
	if dePd != body.Password {
		return errors.New(msg.AUTH_LOGIN_PASSWORD_ERROR)
	}
	var enpwd, _ = util.EnPwdCode(body.Password, info.Salt)
	return accountRepository.ResetPwdBySelf(body.Id, enpwd)
}

// TODO 刷新token
func (i AccountService) RefreshToken(oldToken, ip string) (resDto.TokenAndExp, error) {

	var userClaims comDto.TokenClaims
	userClaims = jwtService.ParseToken(oldToken)
	account, _ = accountRepository.AccountInfo(userClaims.Id)

	accessToken := account.AccessToken
	tokenExp := resDto.TokenAndExp{}
	tokenExp, err = i.refreshTokenRedis(userClaims, ip, accessToken)
	return tokenExp, nil
	//util.ParseToken()
}
