package router

import "github.com/gin-gonic/gin"

type SetupRouter struct {
   server *gin.Engine
}

func NewSetupRouter(s *gin.Engine) *SetupRouter {
   return &SetupRouter{
      server: s,
   }
}

func( sr *SetupRouter ) Setup() {

   v1 := sr.server.Group("v1");

   // PATHS SETUP
   healthPath := NewHealthRouter(v1);

   healthPath.Setup();

   // NOT FOUND
   sr.server.NoRoute(func(ctx *gin.Context) {

      ctx.JSON(404, gin.H{
         "status": false,
         "data": "404 Page not found",
      })
   });
}
