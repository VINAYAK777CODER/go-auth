package controllers

import (
	"net/http"

	"github.com/VINAYAK777CODER/go-auth/initializers"
	"github.com/VINAYAK777CODER/go-auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get the email/pass off req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return

	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// create the user
	user := models.User{Email: body.Email,Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}





	//respond
	c.JSON(http.StatusOK,gin.H{})

}
