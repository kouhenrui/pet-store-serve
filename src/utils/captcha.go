package util

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"pet-store-serve/src/global"
)

type CaptchaService struct {
	Store base64Captcha.Store
}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{Store: &RedisStore{}}
}

const (
	CaptchaHeight          = 80  //高度
	CaptchaWidth           = 240 //宽度
	CaptchaLength          = 4   //长度
	CaptchaNoiseCount      = 0   //干扰数
	CaptcahShowLineOptions = 6   //展示个数
	CaptcahSource          = "1234567890qwertyuioplkjhgfdsazxcvbnm"
)

func (s *CaptchaService) GenerateCaptcha() (string, string, error) {
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driver := base64Captcha.NewDriverString(CaptchaHeight, CaptchaWidth, CaptchaNoiseCount, CaptcahShowLineOptions, CaptchaLength, CaptcahSource, &color.RGBA{R: 3, G: 102, B: 214, A: 125}, nil, nil)
	driver = driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.Store)
	return c.Generate()
}

func (s *CaptchaService) VerifyCaptcha(id, answer string) bool {
	return s.Store.Verify(id, answer, true)
}

type RedisStore struct{}

func (s *RedisStore) Set(id string, value string) error {
	key := global.Captcha + id
	return global.SetRedis(key, []byte(value), global.CaptchaExp) //Redis.Set(ctx, key, value, global.CaptchaExp).Err()
}
func (s *RedisStore) Get(id string, clear bool) string {
	key := global.Captcha + id
	val := global.GetRedis(key)
	err := global.DelRedis(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
