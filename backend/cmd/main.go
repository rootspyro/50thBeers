package main

import (
	"50thbeers/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

   // load env variables
   err := godotenv.Load();

   if err != nil {
      log.Fatal("Error loading env variables!");
   }

   server := gin.Default()

   // ROUTER SETUP

   sRouter := router.NewSetupRouter(server);

   sRouter.Setup();

   server.Run();
}
