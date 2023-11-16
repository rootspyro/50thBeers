// Health path is used to verify the state of the server

package router

import (
	"50thbeers/db"
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

         ctx.JSON(500, gin.H{
            "success": false,
            "data": "Something went wrong!",
         })

         return
      }
      
      ctx.JSON(200, gin.H{
         "success": true,
         "data": "Server is running!",
      })
   })
}

