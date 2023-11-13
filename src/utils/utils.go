package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"path/filepath"
	"pet-store-serve/src/msg"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// json格式化数据
func Marshal(user interface{}) []byte {
	ub, _ := json.Marshal(user)
	return ub
}
func UnMarshal(r []byte, res interface{}) (bool, interface{}) {
	err := json.Unmarshal(r, &res)
	if err != nil {
		return false, msg.REDIS_INFORMATION_ERROR
	}
	return true, res
}

/*
 * @MethodName 参数验证
 * @Description
 * @Author khr
 * @Date 2023/8/21 10:21
 */
func GetValidate(err error, obj any) error {

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return invalid
	}
	//反射获取标签的注释
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {

		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				fmt.Println(msg, "////////////////////////////////")
				return errors.New(msg)
			}
		}
		fmt.Println(errs, "=======================")
		return errs
	}

	fmt.Println(err, "********************")
	return err
}

/*
 * @MethodName ExistIn
 * @Description 判断参数是否存在
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func ExistIn(param string, paths []string) bool {
	for _, v := range paths {
		if param == v {
			return true
		}
	}
	return false
}

/*
 * @MethodName FuzzyMatch
 * @Description 正则模糊匹配路径
 * @Author khr
 * @Date 2023/5/9 16:25
 */
func FuzzyMatch(param string, paths []string) bool {
	for _, y := range paths {
		if regexp.MustCompile(y).MatchString(param) {

			//fmt.Print("匹配道路进了")
			return true
		}

	}
	return false
}

/**
 * @Author Khr
 * @Description //TODO 根据时间戳生成名称
 * @Date 15:41 2023/9/27
 * @Param
 * @return
 **/
func GenerateUniqueFileName(originalFileName string) string {
	// 生成唯一的文件名，可以使用时间戳或随机数等方式
	timestamp := time.Now().UnixNano()
	extension := filepath.Ext(originalFileName)
	uniqueFileName := strconv.FormatInt(timestamp, 10) + extension
	return uniqueFileName
}

// TODO 结构体转切片
func StructToArrayString(s interface{}) []string {
	value := reflect.ValueOf(s)
	if value.Kind() != reflect.Struct {
		return nil
	}

	var result []string
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i).String()
		result = append(result, fieldValue)
	}

	return result
}

// TODO 反射判断结构体字段是否为空
func IsFieldEmpty(s interface{}, fieldName string) bool {
	val := reflect.ValueOf(s).Elem() // 获取结构体的值（指针类型需要使用Elem()方法）

	// 获取字段的值
	fieldValue := val.FieldByName(fieldName)

	// 判断字段是否为空或零值
	return !fieldValue.IsValid() || reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
}
