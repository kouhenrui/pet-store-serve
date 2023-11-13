package global

/**
 * @ClassName clickhouse
 * @Description TODO
 * @Author khr
 * @Date 2023/5/6 9:58
 * @Version 1.0
 */
import (
	"database/sql"
	_ "github.com/ClickHouse/clickhouse-go"
	"log"
	//"sync"
	"time"
)

var cdb *sql.DB

//var dbOnce sync.Once

func ClickhouseInit() {
	db := "tcp://" + clickConfig.Host + ":" + clickConfig.Port + "?" + "username=default&password=" + clickConfig.Password
	// 连接ClickHouse数据库
	//dbOnce.Do(func() {
	cdb, err = sql.Open("clickhouse", db)
	if err != nil {
		return
	}
	cdb.SetMaxIdleConns(10)
	cdb.SetMaxOpenConns(100)
	cdb.SetConnMaxLifetime(30)
	//})
	//defer clickhouseDb.Close()
	log.Printf("click初始化连接成功")
}

type OperationLog struct {
	Id            int32     `json:"id"`
	UserID        uint      `json:"user_id,omitempty"`   // 执行操作的用户ID
	UserName      string    `json:"user_name,omitempty"` //名称
	OperateDate   time.Time `json:"operate_date"`
	OperateType   string    `json:"operate_type" comment:"操作方式 read delete update create"`         //操作类型，读取，创建，删除，修改
	OperateUrl    string    `json:"operate_url"`                                                   //操作路径
	OperateMethod string    `json:"operate_method,omitempty" validate:"oneof=POST GET PUT DELETE"` //方法
	ResourceType  string    `json:"resource_type"`                                                 //资源类型：文件，目录，用户
	Details       string    `json:"details,omitempty" gorm:"comment:'操作详情（可以是JSON格式的操作参数、描述等）'"`   // 操作详情（可以是JSON格式的操作参数、描述等）
	IP            string    `json:"ip,omitempty" gorm:"comment:'用户IP地址'"`                          // 用户IP地址
	UserAgent     string    `json:"userAgent,omitempty" gorm:"comment:'用户代理信息（浏览器、设备等）'"`          // 用户代理信息（浏览器、设备等）
}
type OperateInter interface {
}

func (o OperationLog) OperateSave() {
	//cdb.Exec()
	//cdb.NamedExec("INSERT INTO  operation_log (id,user_id,user_name,operate_date,operate_type,operate_url,operate_method,resource_type,details,ip,userAgent) values " +
	//	(),)
	////cdb.NamedExec("INSERT INTO users (id, name, age) VALUES (:id, :name, :age)", &user)
	//return err
}
