package resDto

import "pet-store-serve/src/dto/reqDto"

type OperateResourceInfo struct {
	Id             uint
	Description    string                `json:"description"`
	ResourceMethod reqDto.ResourceMethod `json:"method"`
	Name           string                `json:"name"`
	URL            string                `json:"url"`
}
