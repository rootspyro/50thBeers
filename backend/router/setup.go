package router

import (
	"50thbeers/db"

	"github.com/gin-gonic/gin"
)

type SetupRouter struct {
   server *gin.Engine
   db     *db.DB
}

func NewSetupRouter(s *gin.Engine, db *db.DB) *SetupRouter {
   return &SetupRouter{
      server: s,
      db: db,
   }
}

func( sr *SetupRouter ) Setup() {

   v1 := sr.server.Group("v1");

   // PATHS SETUP
   healthPath := NewHealthRouter(v1, sr.db);

   healthPath.Setup()

   // NOT FOUND
   sr.server.NoRoute(func(ctx *gin.Context) {

      ctx.JSON(404, gin.H{
         "status": false,
         "data": "404 Page not found",
      })
   })
}
