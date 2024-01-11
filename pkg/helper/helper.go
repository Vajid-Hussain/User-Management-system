package helper

import (
	"fmt"
	"sample/pkg/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SetToken(data models.GenerateToken, c *gin.Context) {
	claims := jwt.MapClaims{
		"email": data.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("qwertyacid12345acidqwerty"))
	if err != nil {
		fmt.Println(err, "error at token generation")
	}
	c.SetCookie("token", tokenString, 3600, "/", "", false, true)
}

func CheckCookie(c *gin.Context) (string, bool) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		// c.String(http.StatusBadRequest, "Cookie not found")
		fmt.Println("cookie not fount")
		return "", false
	}

	secret := []byte("qwertyacid12345acidqwerty")
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		fmt.Println(err, "wronge user with wrong token")
		return "", false
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return email, true
}

// Delete cookie
func DeleteToken(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	_, err := c.Cookie("token")
	if err != nil {
		fmt.Println("wrong in DeleteCookie function", err)
	}
}
