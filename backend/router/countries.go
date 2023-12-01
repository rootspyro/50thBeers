package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"database/sql"
	"log"

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

   cr.group.GET("/countries", cr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {
      
      params := ctx.Request.URL.Query()

      data, err :=  cr.handler.GetItems(params)

      if err != nil {
         models.ServerError(ctx)
         return
      }

      models.OK(ctx, data)
   })

   cr.group.GET("/countries/:id", cr.auth.APIKeyMiddleware(), func(ctx *gin.Context) {

      countryId := ctx.Param("id")

      data, err := cr.handler.GetItem(countryId)

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

   cr.group.POST("/countries", cr.auth.AuthMiddleware(), func(ctx *gin.Context) {

      var body models.CountryBody

      if err := ctx.ShouldBindJSON(&body); err != nil {
         models.InvalidRequest(ctx, err.Error())
         return
      }

      // Validate if country already exits
      _, err := cr.handler.SearchItemByName(body.CountryName)

      if err != nil {
         
         // if the country doesn't exists then creates it
         if err == sql.ErrNoRows {

            data, err := cr.handler.CreateItem(body)

            if err != nil {
               
               log.Println(err)
               models.ServerError(ctx)
               return
            }

            models.Created(ctx, data)
            return
         }

         // ServerError
         log.Println(err)
         models.ServerError(ctx)
         return
      }

      models.Conflict(ctx)
   })

   cr.group.PATCH("/countries/:id", cr.auth.AuthMiddleware(), func(ctx *gin.Context) {

      var (
         countryId string
         body      models.CountryBody
      )

      countryId = ctx.Param("id")
      
      if err := ctx.ShouldBindJSON(&body); err != nil {

         models.InvalidRequest(ctx, err.Error())
         return
      }

      // country validation searching

      _, err := cr.handler.GetItem(countryId)
      
      if err != nil {

         if err == sql.ErrNoRows {

            models.NotFound(ctx)
            return
         }

         log.Println(err)
         models.ServerError(ctx)
         return
      }

      // if country exist then validate that the new countryName is unique
      _, err = cr.handler.SearchItemByName(body.CountryName)

      if err != nil {

         if err == sql.ErrNoRows {

            // Updates the country
            newCountry, err := cr.handler.UpdateItem(countryId, body)

            if err != nil {
               log.Println(err)
               models.ServerError(ctx)
               return
            }

            models.OK(ctx, newCountry)
            return
         }

         log.Println(err)
         models.ServerError(ctx)
         return
      }
      
      models.Conflict(ctx)
      return
   })

   cr.group.DELETE("/countries/:id", cr.auth.AuthMiddleware(), func(ctx *gin.Context) {

      countryId := ctx.Param("id")
      _, err := cr.handler.GetItem(countryId)

      if err != nil {
         
         if err == sql.ErrNoRows {

            models.NotFound(ctx)
            return
         }

         log.Println(err)
         models.ServerError(ctx)
         return
      }

      success := cr.handler.DeleteItem(countryId)

      if !success {
         models.ServerError(ctx)
         return
      }

      models.OK(ctx, "item deleted!")
   })
}
