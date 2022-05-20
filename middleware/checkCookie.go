package middleware

import (
	"GoShopping/database"
	"GoShopping/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

func CheckCookie(c *fiber.Ctx) string {
	// return current user
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "Invalid cookie" // Unauthorized
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	if user.Id == 0 {
		return "Invalid Cookie"
	}

	return "Valid cookie"
}
