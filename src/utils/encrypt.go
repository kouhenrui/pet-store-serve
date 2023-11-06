package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"math/rand"
	"pet-store-serve/src/msg"
	"time"
)

var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*-=+")
var strByteLen = len(strByte)

// 生成12位随机字符串 加两位==
func RandAllString() string {

	bytes := make([]byte, 14)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 14; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return string(bytes) + "=="
	//str := strings.Builder{}
	//length := len(CHARS)
	//for i := 0; i < 14; i++ {
	//	l := CHARS[rand.Intn(length)]
	//	str.WriteString(l)
	//}
	//return str.String() + "=="
}

/**
 * @Author Khr
 * @Description redis缓存名称
 * @Date 14:08 2023/8/29
 * @Param
 * @return
 **/
func Rand6String() string {
	bytes := make([]byte, 4)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}
	return string(bytes) + "=="
}

/**
 * @Author Khr
 * @Description 生成16位密钥
 * @Date 10:52 2023/8/29
 * @Param
 * @return
 **/
func MakeSalt() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 14)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + "=="

}

// PKCS7 填充模式
func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func pKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// 实现加密
func aesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("密钥长度错误：either 16, 24, or 32 bytes to select")
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = pKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 实现解密
func aesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = pKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// 加密base64
func EnPwdCode(Pwd string, pwdKey string) (string, error) {
	pwd := []byte(Pwd)
	PwdKey := []byte(pwdKey)
	result, err := aesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// 解密
func DePwdCode(pwd string, pwdKey string) (string, error) {
	PwdKey := []byte(pwdKey)
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		//errors.New(INTERNAL_ERROR)
		return msg.INTERNAL_ERROR, err
	}
	//执行AES解密
	ecpwsd, erro := aesDeCrypt(pwdByte, PwdKey)
	if erro != nil {
		return msg.PASSWORD_RESOLUTION_ERROR, erro
	}
	return string(ecpwsd), erro

}
