package router

import (
	"50thbeers/db"
	"50thbeers/handlers"
	"50thbeers/models"

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

   // HANDLERS SETUP

   userHandler := handlers.NewUsersHandler();

   // PATHS SETUP
   healthPath := NewHealthRouter(v1, sr.db)
   usersPath  := NewUsersRouter(v1, userHandler)

   healthPath.Setup()
   usersPath.Setup()

   // NOT FOUND
   sr.server.NoRoute(func(ctx *gin.Context) {

      ctx.JSON(404, models.BasicResponse{
         Success: false,
         Data: "404 Page not found...",
      })
   })
}
