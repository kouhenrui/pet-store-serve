package reqDto

type ResourceMethod string

const (
	Delete ResourceMethod = "DELETE"
	Get    ResourceMethod = "GET"
	Post   ResourceMethod = "POST"
	Put    ResourceMethod = "PUT"
)

type OperateResourceAdd struct {
	Description    string         `json:"description"`
	ResourceMethod ResourceMethod `json:"method"`
	Name           string         `json:"name"`
	URL            string         `json:"url"`
}
type OperateResourceUpd struct {
	Id             uint
	Description    string         `json:"description"`
	ResourceMethod ResourceMethod `json:"method"`
	Name           string         `json:"name"`
	URL            string         `json:"url"`
}

type OperateResourceList struct {
	Take           int            `json:"take,omitempty"`
	Skip           int            `json:"skip,omitempty"`
	Name           string         `json:"name,omitempty"`
	ResourceMethod ResourceMethod `json:"method,omitempty"`
}
