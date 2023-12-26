package router

import (
	"50thbeers/auth"
	"50thbeers/handlers"
	"50thbeers/models"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DrinksRouter struct {
  group       *gin.RouterGroup
  handler     *handlers.DrinksHandler
  auth        *auth.AuthHandler
}

func NewDrinksRouter(
  g  *gin.RouterGroup,
  h  *handlers.DrinksHandler,
  a  *auth.AuthHandler,
) *DrinksRouter {

  return &DrinksRouter{
    group: g, 
    handler: h,
    auth: a,
  }
} 

func( dr *DrinksRouter ) Setup() {

  dr.group.GET("/drinks/", dr.auth.APIKeyMiddleware(), dr.handler.GetItems) 

  dr.group.GET("/drinks/:id", dr.auth.APIKeyMiddleware(), dr.handler.GetItem ) 

  dr.group.POST("/drinks", dr.auth.AuthMiddleware(), dr.handler.CreateItem )

  dr.group.PATCH("/drinks/:id", dr.auth.AuthMiddleware(), dr.handler.UpdateItem ) 

  dr.group.PUT("/drinks/:id/publish", dr.auth.AuthMiddleware(), dr.handler.PublishDrink ) 

  dr.group.PUT("/drinks/:id/hide", dr.auth.AuthMiddleware(), dr.handler.HideDrink ) 

  dr.group.DELETE("/drinks/:id", dr.auth.AuthMiddleware(), dr.handler.DeleteDrink ) 

  dr.group.GET("/drinks/:id/tags/", dr.auth.APIKeyMiddleware(), dr.handler.GetItemTags) 

  dr.group.GET("/drinks/:id/tags/:tagId", dr.auth.APIKeyMiddleware(), dr.handler.GetItemTag) 

  dr.group.POST("/drinks/:id/tags/", dr.auth.AuthMiddleware(), dr.handler.CreateItemTag )

  dr.group.PATCH("/drinks/:id/tags/:tagId", dr.auth.AuthMiddleware(), dr.handler.UpdateItemTag ) 

  dr.group.DELETE("/drinks/:id/tags/:tagId", dr.auth.AuthMiddleware(), dr.handler.DeleteItemTag )
}
