package resDto

type PetInfo struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Age         string  `json:"age"`
	Description string  `json:"description"`
	Number      int     `json:"number"`
	Price       float32 `json:"price"`
	Picture     string  `json:"picture"` //轮播显示图片
	ShowPicture string  `json:"show_picture"`
	TypeId      uint    `json:"type_id"`  //绑定PetType
	ClassId     uint    `json:"class_id"` //绑定PetClass
	TypeName    string  `json:"type_name"`
	ClassName   string  `json:"class_name"`
}

type PetTypeInfo struct {
	Id   uint
	Name string
}

type PetClassInfo struct {
	Id     uint
	TypeId uint
	Name   string
}
