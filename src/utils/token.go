package util

import (
	"errors"
	"github.com/go-pay/gopay/pkg/jwt"
	"pet-store-serve/src/dto/comDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/global"
	"pet-store-serve/src/msg"
	"time"
)

var jwtkey = []byte(global.JWTKEY)

type AllClaims struct {
	jwt.StandardClaims
	User comDto.TokenClaims
}
type JwtService struct {
}

// 颁发token inter
func (j *JwtService) SignToken(infoClaims comDto.TokenClaims, day time.Duration) (t resDto.TokenAndExp) {
	expireTime := time.Now().Add(day) //7天过期时间
	claims := &AllClaims{
		User: infoClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "khr",  // 签名颁发者
			Subject:   "sign", //签名主题
		},
	}
	//fmt.Println(claims, "封装的信息")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		errors.New(msg.MAKE_TOKEN_ERROR)
		//fmt.Println(err, "生成token错误")
	}
	tFStr := expireTime.Format("2006-01-02 15:04:05")
	t.Token = tokenString
	t.Exptime = tFStr
	return t
}

//// 验证token
//func AnalysyToken(c *gin.Context) bool {
//	//fmt.Println("进入token验证")
//	tokenString := c.GetHeader("Authorization")
//	if tokenString == "" {
//		return false
//	}
//	return true
//}

// 解析Token
func (j *JwtService) ParseToken(tokenString string) comDto.TokenClaims {
	//fmt.Println(tokenString, "tokenstring")
	//解析token
	token, _ := jwt.ParseWithClaims(tokenString, &AllClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	//fmt.Println(token, "===============================")
	user, _ := token.Claims.(*AllClaims)

	return user.User
}
func (j *JwtService) RefreshToken() {

}
