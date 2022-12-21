package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alphanumericentity/jwt-example/initializers"
	"github.com/alphanumericentity/jwt-example/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func Login(c *gin.Context) {
	// get email and password
	type Request struct {
		Email    string
		Password string
	}

	var request Request

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the request"})
		return
	}

	var user models.User

	err := initializers.DB.Where("email = ?", request.Email).First(&user).Error

	if err != nil || user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email/password is not correct"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email/password is not correct"})
		return
	}

	tokenPeriod := time.Hour
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user.ID,
		"expiry": time.Now().Add(tokenPeriod).Unix()})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not generate jwt token"})
		return
	}

	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", tokenString, int(tokenPeriod.Seconds()), "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "token successfully generated"})

}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}
