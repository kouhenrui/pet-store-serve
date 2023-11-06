package reqDto

type AddRole struct {
	Name string `json:"name" binding:"required"`
}
