package endpoints

import "github.com/gin-gonic/gin"

func WebEndpoint(router *gin.Engine){
	router.Static("/public" , "./web/public")
	router.Static("/fonts" , "./web/public/fonts/")
	router.Static("/pages", "./web/src/pages")
	router.Static("/images" , "./web/public/images")

	web := router.Group("/")
	{
		web.GET("/home" , func(ctx *gin.Context) {
			ctx.File("./web/src/pages/home/index.html")
		})
		web.GET("/post" , func(ctx *gin.Context) {
			ctx.File("./web/src/pages/blog/post/post.html")
		})
	}
}