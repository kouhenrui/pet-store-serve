package pet

import (
	"github.com/gin-gonic/gin"
	"pet-store-serve/src/service/pet"
)

var (
	petService pet.PetInter = &pet.PetService{}
	err        error
)

func PetTypeAdd(c *gin.Context) {

	//petService.
}
