package global

import (
	"crypto/sha256"
	"fmt"
	"github.com/robfig/cron"
	"math/rand"
)

/**
 * @ClassName cron
 * @Description 定时任务
 * @Author khr
 * @Date 2023/4/14 9:19
 * @Version 1.0
 */
var CronTesk *cron.Cron

func init() {
	//CronTesk = cron.New()
	//
	//CronTesk.AddFunc("*/5 * * * * *", addCron1)
	//CronTesk.Start()
	//log.Println("定时任务初始化成功")
}
func addCron1() {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*-=+")
	b := make([]rune, 14)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	combined := []byte(string(b) + "==")
	hash := sha256.Sum256(combined)
	//return fmt.Sprintf("%x", hash)
	fmt.Printf("salt is %x\n", hash)
	//util.DtoToStruct(reqDto.RuleList{}, pojo.Rule{})
	//fmt.Println("Task executed at", time.Now())

}
