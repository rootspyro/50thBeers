package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"log"

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

  dr.group.GET("/drinks", dr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

    params := ctx.Request.URL.Query()

    data, err := dr.handler.GetItems( params )

    if err != nil {
      models.ServerError(ctx) 
      return
    }

    models.OK(ctx, data)
  })

  dr.group.GET("/drinks/:id", dr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id") 

    data, err := dr.handler.GetItem(drinkId) 
    
    if err != nil {

      if err == sql.ErrNoRows {
        models.NotFound(ctx)
        return
      } 

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.OK(ctx, data)
  })

  
  dr.group.POST("/drinks", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    var body models.DrinkPostBody

    err := ctx.ShouldBindJSON(&body)

    if err != nil {
      log.Println(err)
      models.InvalidRequest(ctx, err.Error())
      return
    }

    // generate drinkId
    drinkId := utils.NameToId(body.DrinkName)

    // validate if drink already exits
    _, err = dr.handler.GetItem(drinkId)

    if err != nil {
      
      // if item don't exist then create it
      if err == sql.ErrNoRows {

        data, err := dr.handler.CreateItem(body, drinkId)

        if err != nil {
          log.Println(err)
          models.ServerError(ctx)
          return
        }

        models.Created(ctx, data)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.Conflict(ctx)
    return
  })
}
