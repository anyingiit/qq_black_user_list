package router

import (
	"github.com/gin-gonic/gin"
	"qq_black_user_list/pkg/setting"
	"qq_black_user_list/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())        //,Recovery 中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌，中间件将会写入500。
	gin.SetMode(setting.RunMode) //设置gin运行模式 debugCode releaseCode testCode
	apiv1 := r.Group("api/v1")
	{
		public := apiv1.Group("public")
		{
			public.GET("/even/:qq_num", v1.GetEvenForQq)
		}
		////apiv1.POST("/evev", v1.CreateEven)
		//apiv1.GET("/even/:qq_num", v1.GetEvenForQq)
		////apiv1.DELETE("/even/:id", v1.DeleteEven)
		////apiv1.PUT("/even/:id", v1.SetEvenAuditStatus)
	}
	return r
}
