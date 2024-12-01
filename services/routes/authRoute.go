package routes

import (
	"github.com/Danielyilma/Food_Recipes/services/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.POST("/signup", controllers.Signup())
	incomingRoutes.GET("/.well-known/jwks", controllers.ServeJwk())
}
