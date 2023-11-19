package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"

	"github.com/gin-gonic/gin"
)

type TagRouter struct {
   group   *gin.RouterGroup
   handler *handlers.TagsHandler
   auth    *auth.AuthHandler
}

func NewTagsRouter( g *gin.RouterGroup, th *handlers.TagsHandler, au *auth.AuthHandler ) *TagRouter {
   return &TagRouter{
      group: g,
      handler: th,
      auth: au,
   }
}

func( tr *TagRouter ) Setup() {

   tr.group.GET("/tags", func(ctx *gin.Context) {

      params := ctx.Request.URL.Query()

      data, err := tr.handler.GetItems(params)

      if err != nil {
         models.ServerError(ctx)
         return
      }
      
      models.OK(ctx, data)
   })
}
