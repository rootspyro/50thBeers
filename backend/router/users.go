package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsersRouter  struct {
   group   *gin.RouterGroup
   handler *handlers.UsersHandler
   auth    *auth.AuthHandler
}


func NewUsersRouter( g *gin.RouterGroup, uh *handlers.UsersHandler, au *auth.AuthHandler ) *UsersRouter {
   return &UsersRouter{
      group: g,
      handler: uh,
      auth: au,
   }
}


func ( ur *UsersRouter ) Setup() {

   ur.group.GET("/users", ur.auth.AuthMiddleware(), func(ctx *gin.Context) {

      data, err := ur.handler.GetItems()

      if err != nil {

         models.ServerError(ctx)
         return
      }

      models.OK(ctx, data)

   })

   ur.group.POST("/auth/login", ur.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

      var loginBody models.LoginBody

      if err := ctx.ShouldBindJSON(&loginBody); err != nil {
   
         models.InvalidRequest(ctx)
         return
      }

      data, err := ur.handler.SearchItem(loginBody.User)

      if err != nil {

         if err == sql.ErrNoRows {

            models.NotFound(ctx)
            return
         }

         log.Println(err)
         models.ServerError(ctx)
         return
      }
      
      err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(loginBody.Password))

      if err != nil {
         log.Println(err)
         ctx.JSON(409, models.BasicResponse{
            Success: false,
            Data: "password not match!",
         })

         return
      }

      token, err := ur.auth.GenerateToken(data)
      
      if err != nil {

         log.Println(err)
         models.ServerError(ctx)
         return
      }

      response := models.LoginResponse{
         Token: token,
         Name: data.Username,
         Sub: data.UserID,
         Exp: time.Now().Add(time.Hour * 24).Unix(),
      }

      models.OK(ctx, response)
   })
}
