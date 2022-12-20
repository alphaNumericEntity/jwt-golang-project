package controllers

import (
	"net/http"

	"github.com/alphanumericentity/jwt-example/initializers"
	"github.com/alphanumericentity/jwt-example/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	// get email and password
	type User struct {
		Email    string
		Password string
	}

	var user User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the request"})
		return
	}

	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password"})

		return
	}

	// save the user

	err = initializers.DB.Create(&models.User{Email: user.Email, Password: string(hashed)}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not save user"})

		return
	}

	c.JSON(http.StatusOK, nil)

}
