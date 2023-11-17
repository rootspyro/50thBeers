package router

import (
	"50thbeers/handlers"
	"50thbeers/models"

	"github.com/gin-gonic/gin"
)

type UsersRouter  struct {
   group   *gin.RouterGroup
   handler *handlers.UsersHandler
}


func NewUsersRouter( g *gin.RouterGroup, uh *handlers.UsersHandler ) *UsersRouter {
   return &UsersRouter{
      group: g,
      handler: uh,
   }
}


func ( ur *UsersRouter ) Setup() {

   ur.group.GET("/users", func(ctx *gin.Context) {

      data := ur.handler.GetItems()

      models.OK(ctx, data)

   })
}
