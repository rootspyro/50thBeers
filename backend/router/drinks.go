package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DrinksRouter struct {
  group       *gin.RouterGroup
  handler     *handlers.DrinksHandler
  auth        *auth.AuthHandler
  tagsHandler *handlers.TagsHandler
}

func NewDrinksRouter(
  g  *gin.RouterGroup,
  h  *handlers.DrinksHandler,
  a  *auth.AuthHandler,
  th *handlers.TagsHandler,
) *DrinksRouter {

  return &DrinksRouter{
    group: g, 
    handler: h,
    auth: a,
    tagsHandler: th,
  }
} 

func( dr *DrinksRouter ) Setup() {

  // refactor to routes - handler declarations
  // Example: GET("/drinks", auth.APIKeyMiddleware(), dr.handler.GetItems(ctx))
  dr.group.GET("/drinks/", dr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

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

  dr.group.PATCH("/drinks/:id", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id")

    var body models.DrinkPatchBody
    err := ctx.ShouldBindJSON(&body)

    if err != nil {
      models.InvalidRequest(ctx, err.Error())
      return
    }
  
    // validate if drink exist
    _, err = dr.handler.GetItem(drinkId)

    if err != nil {
      
      if err == sql.ErrNoRows {
        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    // if drink exist -> updates the data
    drink, err := dr.handler.UpdateItem(body, drinkId) 

    if err != nil {

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.OK(ctx, drink)
  })

  dr.group.PUT("/drinks/:id/publish", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id")

    success := dr.handler.ChangeItemStatus(
      drinkId,
      models.DrinksStatuses.Public,
    )

    if !success {
      models.ServerError(ctx) 
      return
    }

    models.OK(ctx, "Drink successfully publicated")

  })

  dr.group.PUT("/drinks/:id/hide", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id")

    success := dr.handler.ChangeItemStatus(
      drinkId,
      models.DrinksStatuses.Created,
    )

    if !success {
      models.ServerError(ctx) 
      return
    }

    models.OK(ctx, "Now the drink is not public!")

  })

  dr.group.DELETE("/drinks/:id", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id")

    success := dr.handler.ChangeItemStatus(
      drinkId,
      models.DrinksStatuses.Deleted,
    )

    if !success {
      models.ServerError(ctx) 
      return
    }

    models.OK(ctx, "Drink successfully deleted!")
  })

  dr.group.GET("/drinks/:id/tags/", dr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

    drinkId := ctx.Param("id")
    drinkTags, err := dr.handler.GetItemTags(drinkId) 

    if err != nil {
      models.ServerError(ctx)
      return
    }

    models.OK(ctx, drinkTags)
  })

  dr.group.GET("/drinks/:id/tags/:tagId", dr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

    drinkId  := ctx.Param("id")
    strTagId := ctx.Param("tagId")

    tagId, err := strconv.Atoi( strTagId )

    if err != nil {
      models.InvalidRequest(ctx, err.Error())
      return
    }

    tag, err := dr.handler.GetItemTag(drinkId, tagId)

    if err != nil {

      if err == sql.ErrNoRows {
        models.NotFound(ctx)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return 
    }

    models.OK(ctx, tag)
  })

  dr.group.POST("/drinks/:id/tags/", dr.auth.AuthMiddleware(), func(ctx *gin.Context) {
  
    var body models.DrinkTagsPostBody

    if err := ctx.ShouldBindJSON(&body); err != nil {
      models.InvalidRequest(ctx, err.Error())
      return
    }

    // validate if tag exist
    
    if _, err := dr.tagsHandler.GetItem(fmt.Sprint(body.TagId)); err != nil {

      ctx.JSON(404, models.BasicResponse{
        Success: false,
        Data: "This tag doesn't exist!",
      })
      return
    } 

    drinkId := ctx.Param("id")

    // validate if the tag is already assigned to the drink
     _, err := dr.handler.GetItemTag(drinkId, body.TagId)

    if err != nil {

      if err == sql.ErrNoRows {

        // add the tag
        newTag, err := dr.handler.CreateItemTag(body, drinkId) 

        if err != nil {

          log.Println(err)
          models.ServerError(ctx)
          return
        }

        models.Created(ctx, newTag)
        return
      }

      log.Println(err)
      models.ServerError(ctx)
      return
    }

    models.Conflict(ctx)
  })
}
