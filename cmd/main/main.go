package main

import (
	"URLServer/DB"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func main() {
	DB.InitDB()

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("hello"))
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/Task", V1_Get_Task_Handler)
		v1.POST("/Task", V1_Post_Task_Handler)

	}

	router.Run(":80")
}

func V1_Get_Task_Handler(ctx *gin.Context) {
	task, _ := DB.GetTask()
	ctx.Render(200, render.String{
		Format: task,
		Data:   nil,
	})
}

func V1_Post_Task_Handler(ctx *gin.Context) {

}
