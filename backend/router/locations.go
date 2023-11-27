package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"

	"github.com/gin-gonic/gin"
)

type LocationsRouter struct {
   group    *gin.RouterGroup
   handler  *handlers.LocationsHandler
   auth     *auth.AuthHandler
}

func NewLocationsRouter( 
   g *gin.RouterGroup, 
   h *handlers.LocationsHandler, 
   a *auth.AuthHandler,
) *LocationsRouter {
   return &LocationsRouter{
      group:  g,
      handler: h,
      auth: a,
   }
}

func( lr *LocationsRouter ) Setup() {

   lr.group.GET("/locations", lr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

      params := ctx.Request.URL.Query()

      data, err := lr.handler.GetItems(params)

      if err != nil {

         models.ServerError(ctx)
         return
      }

      models.OK(ctx, data)
   })
}
