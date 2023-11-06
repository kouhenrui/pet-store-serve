package comDto

import "pet-store-serve/src/pojo"

/**
 * @ClassName common
 * @Description TODO
 * @Author khr
 * @Date 2023/5/8 13:37
 * @Version 1.0
 */

type TokenClaims struct {
	Id       uint        `json:"id"`
	Phone    string      `json:"phone"`
	Name     string      `json:"name"`
	NickName string      `json:"nick_name"`
	Class    string      `json:"class"`
	Role     []pojo.Role `json:"role"`
}
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Cabin struct {
	Sub string `json:"sub,omitempty"`
	Obj string `json:"obj,omitempty"`
	Act string `json:"act,omitempty"`
}
type GCabin struct {
	Type string `json:"type" binding:"oneof=g g2"`
	Sub  string `json:"sub,omitempty"`
	Obj  string `json:"obj,omitempty"`
}
