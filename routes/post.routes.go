package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Praveenkusuluri08/controller"
)

func TestRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messagee": "Hello world!",
	})
}

func PostRoutes(incommingRoutes *gin.Engine) {
	//TODO:Test route
	v := incommingRoutes.Group("/api")
	{
		v.GET("test", controller.TestRoute)
	}
	incommingRoutes.POST("/createpost", controller.CreatePost)
	incommingRoutes.GET("/getallposts", controller.GetAllPosts)
	incommingRoutes.GET("/getsinglepost", controller.GetSinglePost)
	incommingRoutes.GET("/getsinglepost/:id", controller.GetSinglePostParams)
	incommingRoutes.PUT("/updatepost", controller.UpdateSinglePost)
	incommingRoutes.DELETE("/deletepost", controller.DeleteSinglePost)
	incommingRoutes.DELETE("/deletepostpermanently", controller.DeletePostPermanently)
}
