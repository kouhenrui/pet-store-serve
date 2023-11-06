package main

import (
	"fmt"
	"pet-store-serve/src/global"
	"pet-store-serve/src/route"
)

func main() {
	r := route.InitRoute()
	//fmt.Println("路由挂载在main程序")
	if err := r.Run(global.Port, "https/certificate.crt", "https/private.key"); err != nil {
		fmt.Println(fmt.Errorf("端口占用,err:%v\n", err))
	}
	art :=
		`
   ###      ##     #          ##     #    #   ####               ####    #    #     ###      ###    ######    ####     ####
  #   #    #  #    #         #  #    ##   #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
 #        #    #   #        #    #   # #  #    #   #            #        #    #   #        #        #        #        #
 #  ###   #    #   #        ######   #  # #    #   #             ####    #    #   #        #        ####      ####     ####
 #    #   #    #   #        #    #   #   ##    #   #                 #   #    #   #        #        #             #        #
  #   #    #  #    #        #    #   #    #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
   ###      ##     ######   #    #   #    #   ####               ####     ####      ###      ###    ######    ####     ####
`
	fmt.Println(art)
}
