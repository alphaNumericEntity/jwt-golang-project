package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alphanumericentity/jwt-example/initializers"
	"github.com/alphanumericentity/jwt-example/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errr": "could not find cookie"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errr": "expired token"})
		}

		var user models.User

		initializers.DB.First(&user, claims["user"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errr": "couild not find user"})
		}

		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errr": err})
	}

}
