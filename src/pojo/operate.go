package pojo

import "gorm.io/gorm"

type OperationLog struct {
	gorm.Model `json:"gorm.Model"`
	UserID     uint   `json:"userID,omitempty" gorm:"comment:'执行操作的用户ID'"`                                             // 执行操作的用户ID
	UserName   string `json:"userName,omitempty" gorm:"comment:'名称'"`                                                  //名称
	Way        string `json:"way,omitempty" gorm:"not null;comment:'方法'" binding:"oneof=POST GET PUT DELETE,required"` //方法
	Path       string `json:"path,omitempty" gorm:"comment:'请求路基'"`                                                    //请求路基
	Details    string `json:"details,omitempty" gorm:"comment:'操作详情（可以是JSON格式的操作参数、描述等）'"`                             // 操作详情（可以是JSON格式的操作参数、描述等）
	IP         string `json:"IP,omitempty" gorm:"comment:'用户IP地址'"`                                                    // 用户IP地址
	UserAgent  string `json:"userAgent,omitempty" gorm:"comment:'用户代理信息（浏览器、设备等）'"`                                    // 用户代理信息（浏览器、设备等）
}
