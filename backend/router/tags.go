package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"

	"github.com/gin-gonic/gin"
)

type TagRouter struct {
   group   *gin.RouterGroup
   handler *handlers.TagsHandler
   auth    *auth.AuthHandler
}

func NewTagsRouter( 

   g *gin.RouterGroup, 
   th *handlers.TagsHandler, 
   au *auth.AuthHandler,

) *TagRouter {

   return &TagRouter{
      group: g,
      handler: th,
      auth: au,
   }
}

func( tr *TagRouter ) Setup() {
   
   tr.group.GET("/tags", tr.auth.APIKeyMiddleware(), tr.handler.GetItems) 

   tr.group.GET("/tags/:id", tr.auth.APIKeyMiddleware(), tr.handler.GetItem )

   tr.group.POST("/tags", tr.auth.AuthMiddleware(), tr.handler.CreateItem) 

   tr.group.PATCH("/tags/:id", tr.auth.AuthMiddleware(), tr.handler.UpdateItem ) 

   tr.group.DELETE("/tags/:id", tr.auth.AuthMiddleware(), tr.handler.DeleteItem )
}

