// Health path is used to verify the state of the server

package router

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"

	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
   group *gin.RouterGroup
   db    *db.DB
}

func NewHealthRouter( g *gin.RouterGroup, db *db.DB ) *HealthRouter {
   return &HealthRouter{
      group: g,
      db: db,
   }
}

func ( hr *HealthRouter ) Setup() {
  
   hr.group.GET("/health", func(ctx *gin.Context) {

      err := hr.db.Conn.Ping()

      if err != nil {

         log.Printf("Database PING error: %v", err)

         models.ServerError(ctx);

         return
      }
      
      ctx.JSON(200, gin.H{
         "success": true,
         "data": "Server is running!",
      })
   })
}

