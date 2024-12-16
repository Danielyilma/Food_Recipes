package routes

import (
	"github.com/Danielyilma/Food_Recipes/services/controllers"
	"github.com/Danielyilma/Food_Recipes/services/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.POST("/signup", controllers.Signup())
	incomingRoutes.GET("/.well-known/jwks", controllers.ServeJwk())
	incomingRoutes.POST("/upload", middleware.JWTMiddleware(), controllers.FileUpload())
	incomingRoutes.POST("/single-upload", middleware.JWTMiddleware(), controllers.Upload())
	incomingRoutes.POST("/update-recipe", middleware.JWTMiddleware(), controllers.UpdateRecipe())
	incomingRoutes.POST("/payment", middleware.JWTMiddleware(), controllers.Payment())
}
