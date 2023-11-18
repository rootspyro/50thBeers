// The basic JSON responses structs

package models

import "github.com/gin-gonic/gin"

type BasicResponse struct {
   Success bool   `json:"success"`
   Data    any `json:"data"`
}

// HTTP 200 OK

func OK( ctx *gin.Context, data any ) {

   response := BasicResponse {
      Success: true,
      Data: data,
   }

   ctx.JSON(200, response) 
}

// HTTP 400 INVALID REQUEST 
func InvalidRequest( ctx *gin.Context ) {
   
   response := BasicResponse{
      Success: false,
      Data: "Invalid request!",
   }

   ctx.JSON(400, response)
}

// HTTP 401 UNAUTHORIZED 
func Unauthorized( ctx *gin.Context ) {
   
   response := BasicResponse{
      Success: false,
      Data: "Unauthorized",
   }

   ctx.JSON(401, response)
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


