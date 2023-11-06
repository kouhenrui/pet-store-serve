package resDto

/**
* @program: work_space
*
* @description:返回参数格式化
*
* @author: khr
*
* @create: 2023-02-01 14:15
**/
type AccountList struct {
	Id      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Account string `json:"account,omitempty"`
	Role    int    `json:"role,omitempty"`
}

// 详情数据
type AccountInfo struct {
	Id       uint     `json:"id"`
	UserName string   `json:"name"`
	NickName string   `json:"nick_name"`
	Phone    string   `json:"phone"`
	Role     []string `json:"role"`
	Sex      int      `json:"sex" `
	Avatar   string   `json:"avatar" `
	Email    string   `json:"email"`
	Class    string   `json:"class"`
}

//type AdminInformation struct {
//
//}
