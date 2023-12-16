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

   lr.group.GET("/locations/:id", lr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

    locationId := ctx.Param("id")

    data, err := lr.handler.GetItem(locationId)

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

   lr.group.POST("/locations", lr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    var body models.LocationBody

    if err := ctx.ShouldBindJSON(&body); err != nil {

      models.InvalidRequest(ctx, err.Error())
      return
    } 

    newId := utils.NameToId(body.LocationName)

    // search if the item already exist
    _, err := lr.handler.GetItem(newId)

    if err != nil {

      // if location doesn't exist then create it
      if err == sql.ErrNoRows {

        location, err := lr.handler.CreateItem(body)
        
        if err != nil {

          log.Println(err)
          models.ServerError(ctx)
          return
        }

        models.Created(ctx, location)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.Conflict(ctx)
  })

  lr.group.PATCH("/locations/:id", lr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    locationId := ctx.Param("id")
    
    var body models.LocationBody 
    err := ctx.ShouldBindJSON(&body)
    
    if err != nil {
      models.InvalidRequest(ctx, err.Error()) 
    }

    // validate if location exist

    _, err = lr.handler.GetItem(locationId)

    if err != nil {

      if err == sql.ErrNoRows {
        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    // if item exist then update the data
    location, err := lr.handler.UpdateItem(body, locationId)

    if err != nil {

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.OK(ctx, location)
  })

  lr.group.PUT("/locations/:id/publicate", func(ctx *gin.Context) {

    locationId := ctx.Param("id")

    _, err := lr.handler.GetItem(locationId)

    if err != nil {

      if err == sql.ErrNoRows {

        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    success := lr.handler.PublicateItem(locationId)

    if !success {

      models.ServerError(ctx)
      return
    }

    models.OK(ctx, "Item successfully publicated!")
  })

  lr.group.PUT("/locations/:id/hide", func(ctx *gin.Context) {

    locationId := ctx.Param("id")

    _, err := lr.handler.GetItem(locationId)

    if err != nil {

      if err == sql.ErrNoRows {

        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    success := lr.handler.HideItem(locationId)

    if !success {

      models.ServerError(ctx)
      return
    }

    models.OK(ctx, "This location is not public now!")
  })

  lr.group.DELETE("/locations/:id", func(ctx *gin.Context) {

    locationId := ctx.Param("id")

    _, err := lr.handler.GetItem(locationId)

    if err != nil {

      if err == sql.ErrNoRows {

        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    success := lr.handler.DeleteItem(locationId)

    if !success {

      models.ServerError(ctx)
      return
    }

    models.OK(ctx, "Item successfully deleted!")
  })
}

