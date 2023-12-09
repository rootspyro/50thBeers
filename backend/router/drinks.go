package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"

	"github.com/gin-gonic/gin"
)

type DrinksRouter struct {
  group   *gin.RouterGroup
  handler *handlers.DrinksHandler
  auth    *auth.AuthHandler
}

func NewDrinksRouter(
  g *gin.RouterGroup,
  h *handlers.DrinksHandler,
  a *auth.AuthHandler,
) *DrinksRouter {

  return &DrinksRouter{
    group: g, 
    handler: h,
    auth: a,
  }
} 

func( dr *DrinksRouter ) Setup() {

  dr.group.GET("/drinks", func(ctx *gin.Context) {

    params := ctx.Request.URL.Query()

    data, err := dr.handler.GetItems( params )

    if err != nil {
      models.ServerError(ctx) 
      return
    }

    models.OK(ctx, data)
  })
}
