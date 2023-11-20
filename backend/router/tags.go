package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"database/sql"
	"log"

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

   tr.group.GET("/tags/:id", func(ctx *gin.Context) {

      tagId := ctx.Param("id")

      data, err := tr.handler.GetItem(tagId)

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

   tr.group.POST("/tags", tr.auth.AuthMiddleware(), func(ctx *gin.Context) {

      var body models.TagBody

      if err := ctx.ShouldBindJSON(&body); err != nil {
         models.InvalidRequest(ctx)
         return
      }

      // Validate if the tag already exist
      _, err := tr.handler.SearchItemByName(body.TagName)

      if err != nil {

         if err == sql.ErrNoRows {
            // if the tag doesn't exists then creates the tag

            data, err := tr.handler.CreateItem(body)

            if err != nil {

               log.Println(err)
               models.ServerError(ctx)
               return
            }

            models.Created(ctx, data)
            return
         }

         // server error
         log.Println(err)
         models.ServerError(ctx)
         return
      } 

      // if the tag exists send 409 CONFLICT
      models.Conflict(ctx) 
   })

   tr.group.PATCH("/tags/:id", tr.auth.AuthMiddleware(), func(ctx *gin.Context) {

      // We get the request body and path Id
      var (
         tagId string
         body  models.TagBody
      ) 

      tagId = ctx.Param("id")

      err := ctx.ShouldBindJSON(&body)

      if err != nil {

         log.Println(err)
         models.InvalidRequest(ctx)
         return
      }

      // tag validation searching
      tag, err := tr.handler.GetItem(tagId)

      if err != nil {

         if err == sql.ErrNoRows {
            models.NotFound(ctx)
            return
         }

         log.Println(tag)
         models.ServerError(ctx)
         return
      }

      // if tag exists then validate that the new tagname is unique
      _, err = tr.handler.SearchItemByName(body.TagName)

      if err != nil {

         if err == sql.ErrNoRows {

            // the updates the tag
            newTag, err := tr.handler.UpdateItem(body, tagId)

            if err != nil {
               log.Println(err)
               models.ServerError(ctx)
               return
            }

            models.OK(ctx, newTag)
            return
         }

         log.Println(err)
         models.ServerError(ctx)
      }

      models.Conflict(ctx)
   })
}

