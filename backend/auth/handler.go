package auth

import (
	"50thbeers/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
   secret []byte
}

func NewAuthHandler( s string ) *AuthHandler {
   return &AuthHandler{
      secret: []byte(s),
   }
}

func( ah *AuthHandler ) GenerateToken( data models.User ) (string, error) {

   token := jwt.New(jwt.SigningMethodHS256)
   claims := token.Claims.(jwt.MapClaims)
   claims["name"] = data.Username
   claims["sub"] = data.UserID
   claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

   tokenString, err := token.SignedString(ah.secret)

   if err != nil {
      return "", err
   }

   return tokenString, nil
}
