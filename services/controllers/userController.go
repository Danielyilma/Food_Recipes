package controllers

import (
	"log"
	"net/http"

	"github.com/Danielyilma/Food_Recipes/services/databases"
	"github.com/Danielyilma/Food_Recipes/services/helpers"
	"github.com/Danielyilma/Food_Recipes/services/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			log.Println("Invalid JSON received:", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		validatorErr := validate.Struct(user)

		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr.Error()})
			return
		}

		hashed_password := helpers.HashPassword(user.Password)
		user.Password = &hashed_password

		query := `INSERT INTO users (username, email, password, phone_number)
		VALUES ($1, $2, $3, $4) RETURNING id`
		var userID int
		err := databases.DBConnection.QueryRow(query, user.Username, user.Email, user.Password, user.Phone_number).Scan(&userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Id = userID

		token, err := helpers.CreateJwt(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user, "token": token})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.LoginForm

		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validatorErr := validate.Struct(login)

		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr.Error()})
			return
		}

		user, err := helpers.GetUserByEmail(*login.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not found with that email"})
			return
		}

		if helpers.CheckPassword(*user.Password, *login.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password or email"})
			return
		}

		token, err := helpers.CreateJwt(*user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func ServeJwk() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwk, err := helpers.GetJWK()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, jwk)
	}
}
