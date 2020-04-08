package Service

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("hello"))
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/GetTask", V1_Get_Task_Handler)
	}
}

func V1_Get_Task_Handler(ctx *gin.Context) {

}
