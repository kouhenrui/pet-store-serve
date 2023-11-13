package common

import (
	"pet-store-serve/src/dto/reqDto"
	util "pet-store-serve/src/utils"
)

var (
	err    error
	capcha = util.NewCaptchaService()
)

type CommonServiceImpl struct{}
type CommonServiceInter interface {
	GetCaptcha() (error, reqDto.Captcha)
	VfCaptcha(capt reqDto.Captcha) bool
}

/**
 * @Author Khr
 * @Description //TODO 调用方法生成验证码
 * @Date 9:32 2023/9/27
 * @Param
 * @return
 **/
func (c CommonServiceImpl) GetCaptcha() (error, reqDto.Captcha) {

	var newCaptcha reqDto.Captcha

	newCaptcha.Id, newCaptcha.Content, err = capcha.GenerateCaptcha()
	if err != nil {
		return err, newCaptcha
	}
	return nil, newCaptcha
}

/**
 * @Author Khr
 * @Description //TODO 验证输入码
 * @Date 9:33 2023/9/27
 * @Param
 * @return
 **/
func (c CommonServiceImpl) VfCaptcha(capt reqDto.Captcha) bool {
	return capcha.VerifyCaptcha(capt.Id, capt.Content)
}
