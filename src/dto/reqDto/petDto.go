package reqDto

type PetList struct {
	Take    int    `json:"take,omitempty" validate:"required"`
	Skip    uint   `json:"skip,omitempty"`
	TypeId  string `json:"type_id"`
	ClassId string `json:"class_id"`
	Title   string `json:"title"`
}
type PetAdd struct {
	Title       string  `json:"title"`
	Age         string  `json:"age"`
	Description string  `json:"description"`
	Number      int     `json:"number"`
	Price       float32 `json:"price"`
	Picture     string  `json:"picture"` //轮播显示图片
	ShowPicture string  `json:"show_picture"`
	TypeId      uint    `json:"type_id"`  //绑定PetType
	ClassId     uint    `json:"class_id"` //绑定PetClass
}
type PetUpd struct {
	Id          uint    `json:"id" validate:"required"`
	Title       string  `json:"title"`
	Age         string  `json:"age"`
	Description string  `json:"description"`
	Number      int     `json:"number"`
	Price       float32 `json:"price"`
	Picture     string  `json:"picture"` //轮播显示图片
	ShowPicture string  `json:"show_picture"`
	TypeId      uint    `json:"type_id"`  //绑定PetType
	ClassId     uint    `json:"class_id"` //绑定PetClass
}
type PetTypeList struct {
	Take int    `json:"take,omitempty" validate:"required"`
	Skip int    `json:"skip,omitempty"`
	Name string `json:"name"`
}
type PetTypeUpd struct {
	Id   uint   `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}
type PetTypeAdd struct {
	Name string `json:"name" validate:"required"`
}

type PetClassAdd struct {
	Name   string `json:"name" validate:"required"`
	TypeId uint   `json:"type_id" validate:"required"`
}
type PetClassList struct {
	Take   int    `json:"take,omitempty" validate:"required"`
	Skip   int    `json:"skip,omitempty"`
	Name   string `json:"name" validate:"required"`
	TypeId uint   `json:"type_id" validate:"required"`
}
type PetClassUpd struct {
	Id     uint   `json:"id,omitempty" validate:"required"`
	Name   string `json:"name,omitempty" validate:"required"`
	TypeId uint   `json:"type_id"`
}
