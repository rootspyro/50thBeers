package main

import (
	"50thbeers/db"
	"50thbeers/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

   // load env variables
   err := godotenv.Load()

   if err != nil {
      log.Fatal("Error loading env variables!")
   }

   server := gin.Default()

   // DATABASE SETUP

   dbconn := db.NewDBConnection(
      os.Getenv("PG_HOST"),
      os.Getenv("PG_USER"),
      os.Getenv("PG_PASSWORD"),
      os.Getenv("PG_DATABASE"),
      os.Getenv("PG_PORT"),
   )

   defer dbconn.Conn.Close()

   // ROUTER SETUP

   sRouter := router.NewSetupRouter(server, dbconn)

   sRouter.Setup()

   server.Run()
}
