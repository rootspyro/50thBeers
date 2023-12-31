package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"

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

  lr.group.GET("/locations", lr.auth.APIKeyMiddleware(), lr.handler.GetItems ) 

  lr.group.GET("/locations/:id", lr.auth.APIKeyMiddleware(), lr.handler.GetItem ) 

  lr.group.POST("/locations", lr.auth.AuthMiddleware(), lr.handler.CreateItem )

  lr.group.PATCH("/locations/:id", lr.auth.AuthMiddleware(), lr.handler.UpdateItem )

  lr.group.PUT("/locations/:id/publish", lr.auth.AuthMiddleware(), lr.handler.PublishItem ) 

  lr.group.PUT("/locations/:id/hide", lr.auth.AuthMiddleware(), lr.handler.HideItem )

  lr.group.DELETE("/locations/:id", lr.auth.AuthMiddleware(), lr.handler.DeleteItem )

}

