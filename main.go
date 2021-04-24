package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"qq_black_user_list/models"
	"qq_black_user_list/pkg/casbin"
	"qq_black_user_list/router"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := models.InitDB(); err != nil {
		fmt.Println("init db err:", err.Error())
		os.Exit(-999)
	}
	if err := casbin.InitCasbin(); err != nil {
		fmt.Println("init casbin err:", err)
		os.Exit(-6666)
	}
	r := router.InitRouter()
	if err := r.Run(); err != nil {
		fmt.Println("init gin Run err:", err.Error())
		os.Exit(-666)
	}
}
