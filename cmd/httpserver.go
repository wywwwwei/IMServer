package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wywwwwei/IMServer/Service"
)
func setupHttp(){
	router := gin.Default()
	router.GET("/list/:uid",Service.ListHandler)
	router.GET("/profile/:uid",Service.ProfileHandler)
	router.GET("/username/:uid",Service.UsernameHandler)
	router.POST("/login",Service.LoginHandler)
	router.POST("/regist",Service.RegistHandler)

	router.Run(fmt.Sprintf(":%d",Service.SERVER_HTTP_PORT))
}
