package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"

	"github.com/gin-gonic/gin"
)

type CountriesRouter struct {
   group   *gin.RouterGroup
   handler *handlers.CountriesHandler
   auth    *auth.AuthHandler
}

func NewCountriesRouter( 

   g *gin.RouterGroup, 
   ch *handlers.CountriesHandler, 
   au *auth.AuthHandler,

) *CountriesRouter {

   return &CountriesRouter{
      group: g,
      handler: ch,
      auth: au,
   }
}

func( cr *CountriesRouter ) Setup() {

   cr.group.GET("/countries", cr.auth.APIKeyMiddleware(), cr.handler.GetItems)
   cr.group.GET("/countries/:id", cr.auth.APIKeyMiddleware(), cr.handler.GetItem )
   cr.group.POST("/countries", cr.auth.AuthMiddleware(), cr.handler.CreateItem)
   cr.group.PATCH("/countries/:id", cr.auth.AuthMiddleware(), cr.handler.UpdateItem)
   cr.group.DELETE("/countries/:id", cr.auth.AuthMiddleware(), cr.handler.DeleteItem)
}
