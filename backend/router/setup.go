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

   authHandler := auth.NewAuthHandler(
      os.Getenv("SECRET"),
      os.Getenv("API_KEY"),
   )

   // TABLES SETUP

   usersTable     := db.NewUsersTable(sr.db)
   tagsTable      := db.NewTagsTable(sr.db)
   countriesTable := db.NewCountriesTable(sr.db)
   locationsTable := db.NewLocationsTable(sr.db)
   drinksTable    := db.NewDrinksTable(sr.db, tagsTable.Table, countriesTable.Table, locationsTable.Table)

   // HANDLERS SETUP

   userHandler      := handlers.NewUsersHandler(usersTable);
   tagsHandler      := handlers.NewTagsHandler(tagsTable)
   countriesHandler := handlers.NewCountriesHandler(countriesTable)
   locationsHandler := handlers.NewLocationsHandler(locationsTable) 
   drinksHandler    := handlers.NewDrinksHandler(drinksTable)

   // PATHS SETUP
   healthPath    := NewHealthRouter(v1, sr.db)
   usersPath     := NewUsersRouter(v1, userHandler, authHandler)
   tagsPath      := NewTagsRouter(v1, tagsHandler, authHandler)
   countriesPath := NewCountriesRouter(v1, countriesHandler, authHandler)
   locationsPath := NewLocationsRouter(v1, locationsHandler, authHandler)
   drinksPath    := NewDrinksRouter(v1, drinksHandler, authHandler)

   healthPath.Setup()
   usersPath.Setup()
   tagsPath.Setup()
   countriesPath.Setup()
   locationsPath.Setup()
   drinksPath.Setup()

   // NOT FOUND
   sr.server.NoRoute(func(ctx *gin.Context) {

      ctx.JSON(404, models.BasicResponse{
         Success: false,
         Data: "404 Page not found...",
      })
   })
}
