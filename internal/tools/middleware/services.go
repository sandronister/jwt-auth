package middleware

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SignedDetails struct {
	Id         int
	Authorized bool
	Data       interface{}
	jwt.StandardClaims
}

func getEnvs() (*int, *string) {
	TOKEN_HOUR_LIFESPAN, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		panic("token lifespan not set")
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	if JWT_SECRET == "" {
		panic("jwt secret not set")
	}

	return &TOKEN_HOUR_LIFESPAN, &JWT_SECRET
}

func GenerateToken(id int, data interface{}) (string, error) {

	TOKEN_HOUR_LIFESPAN, JWT_SECRET := getEnvs()

	claims := &SignedDetails{
		Id:         id,
		Authorized: true,
		Data:       data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(*TOKEN_HOUR_LIFESPAN)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(*JWT_SECRET))

}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		_, JWT_SECRET := getEnvs()
		return []byte(*JWT_SECRET), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok && !token.Valid {
		return err
	}

	c.Set("user_id", claims.Id)
	c.Set("user_authorized", claims.Authorized)
	c.Set("user_data", claims.Data)

	return nil

}
