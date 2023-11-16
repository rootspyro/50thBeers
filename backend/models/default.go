// The basic JSON responses structs

package models

import "github.com/gin-gonic/gin"

type BasicResponse struct {
   Success bool   `json:"success"`
   Data    string `json:"data"`
}

// HTTP 404 NOT FOUND
func NotFound( ctx *gin.Context ) {
   
   response := BasicResponse{
      Success: false,
      Data: "Item not found!",
   }

   ctx.JSON(404, response)
}

// HTTP 409 CONFLICT
func Conflict( ctx *gin.Context ) {
   
   response := BasicResponse{
      Success: false,
      Data: "Item already exist!",
   }

   ctx.JSON(409, response)
}

// HTTP 500 SERVER ERROR 
func ServerError( ctx *gin.Context ) {

   response := BasicResponse{
      Success: false,
      Data: "Something Went Wrong!",
   }

   ctx.JSON(500, response) 
}


