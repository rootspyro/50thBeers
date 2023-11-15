// Health path is used to verify the state of the server

package router

import "github.com/gin-gonic/gin"

type HealthRouter struct {
   group *gin.RouterGroup
}

func NewHealthRouter( g *gin.RouterGroup ) *HealthRouter {
   return &HealthRouter{
      group: g,
   }
}

func ( hr *HealthRouter ) Setup() {
  
   hr.group.GET("/health", func(ctx *gin.Context) {
      
      ctx.JSON(200, gin.H{
         "success": true,
         "data": "Server is running!",
      })
   });
}

