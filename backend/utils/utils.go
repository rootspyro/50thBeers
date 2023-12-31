package utils

import (
	"50thbeers/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func NameToId( name string ) string {

  name = strings.ToLower(name)
  name = strings.ReplaceAll(name, " ", "_")

  return name
}

// this function sends the json status and response for a server error
// and prints in log the error
func ServerError( ctx *gin.Context, err error ) {
  
  log.Println(err.Error())
  models.ServerError(ctx)

}
