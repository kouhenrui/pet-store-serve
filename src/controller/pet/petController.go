package pet

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/service/pet"
	util "pet-store-serve/src/utils"
)

var (
	petService pet.PetInter = &pet.PetService{}
	err        error
)

func PetTypeList(c *gin.Context) {
	var list reqDto.PetTypeList
	if err = c.Bind(&list); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &list))
		return
	}
	petTypeList, err := petService.PetTypeList(list)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", petTypeList)
	return
}
func PetTypeInfo(c *gin.Context) {
	log.Println("--------------------------------------------------")
	var id = c.Query("id")
	fmt.Println(id, "********************************")
	c.Set("res", id)
	return
	//petService.PetTypeInfo(id)
}
