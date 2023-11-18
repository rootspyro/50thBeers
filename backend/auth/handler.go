package auth

import (
	"50thbeers/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func( ah *AuthHandler ) AuthMiddleware() gin.HandlerFunc {
   return func(ctx *gin.Context) {

      tokenString := ctx.GetHeader("Authorization")

      if tokenString == "" {

         models.Unauthorized(ctx)
         ctx.Abort()
         return
      }

      tokenFiltred := strings.Split(tokenString, " ")

      token, err := jwt.Parse(tokenFiltred[1], func(token *jwt.Token) (interface{}, error) {

         return []byte(os.Getenv("SECRET")), nil
      })

      if err != nil || !token.Valid {

         models.Unauthorized(ctx)

         log.Println(err)
         ctx.Abort()
         return
      }

      ctx.Next()
   }
}
