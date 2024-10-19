package routes 

import (
	"github.com/gin-gonic/gin"
	"games-info-api/src/controllers"
)

func SetRoutes(router *gin.Engine){
	gamesRoutes := router.Group("/games")
	{
		gamesRoutes.GET("/", controllers.GetGames)
	}
}