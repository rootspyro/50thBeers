package router

import (
	"50thbeers/auth"
	"50thbeers/db"
	"50thbeers/handlers"
	"50thbeers/models"
	"os"

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

   // AUTH SETUP

   authHandler := auth.NewAuthHandler(os.Getenv("SECRET"))

   // TABLES SETUP

   usersTable := db.NewUsersTable(sr.db)
   tagsTable := db.NewTagsTable(sr.db)

   // HANDLERS SETUP

   userHandler := handlers.NewUsersHandler(usersTable);
   tagsHandler := handlers.NewTagsHandler(tagsTable)

   // PATHS SETUP
   healthPath := NewHealthRouter(v1, sr.db)
   usersPath  := NewUsersRouter(v1, userHandler, authHandler)
   tagsPath   := NewTagsRouter(v1, tagsHandler, authHandler)

   healthPath.Setup()
   usersPath.Setup()
   tagsPath.Setup()

   // NOT FOUND
   sr.server.NoRoute(func(ctx *gin.Context) {

      ctx.JSON(404, models.BasicResponse{
         Success: false,
         Data: "404 Page not found...",
      })
   })
}
