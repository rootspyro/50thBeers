package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/dict"
)

type DrinksHandler struct {
  drinksTable *db.DrinksTable
  tagsTable   *db.TagsTable
}

func NewDrinksHandler( dt *db.DrinksTable, tt *db.TagsTable ) *DrinksHandler {
  return &DrinksHandler{
    drinksTable: dt,
    tagsTable: tt,
  }
}

func( dh *DrinksHandler ) GetItems( ctx *gin.Context ) {
  
  params := ctx.Request.URL.Query()

  items, itemsFound, err := dh.drinksTable.GetAllDrinks(params)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, models.DrinkCollection{
    ItemsFound: itemsFound,
    Items: items,
    Filters: dh.drinksTable.Filters,
  })
} 

func( dh *DrinksHandler ) GetItem( ctx *gin.Context ) {
  drinkId := ctx.Param("id")

  drink, err := dh.drinksTable.GetSingleDrink(drinkId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, drink)
}

func( dh *DrinksHandler ) CreateItem( ctx *gin.Context ) {
  
  var body models.DrinkPostBody

  if err := ctx.ShouldBindJSON(&body); err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  }

  // generate drinkId
  drinkId := utils.NameToId(body.DrinkName)  

  // validate if the drink already exist
  _, err := dh.drinksTable.GetSingleDrink(drinkId)

  if err == nil {
    models.Conflict(ctx) // item already exist
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  // create the new drink

  newDrink, err := dh.drinksTable.CreateDrink(body, drinkId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.Created(ctx, newDrink)
}


func( dh *DrinksHandler ) UpdateItem( ctx *gin.Context) {

  drinkId := ctx.Param("id")
  var body models.DrinkPatchBody

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
  }

  // validate if the drink exist
  _, err := dh.drinksTable.GetSingleDrink(drinkId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // update the data

  drink, err := dh.drinksTable.UpdateDrink(body, drinkId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, drink)
}

func( dh *DrinksHandler ) PublishDrink( ctx *gin.Context ) {

  drinkId := ctx.Param("id")

  err := dh.drinksTable.ChangeStatus(drinkId, models.DrinksStatuses.Public)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "Drink successfully published")
}

func( dh *DrinksHandler ) HideDrink( ctx *gin.Context ) {

  drinkId := ctx.Param("id")

  err := dh.drinksTable.ChangeStatus(drinkId, models.DrinksStatuses.Created)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "This drink is not public now")
}

func( dh *DrinksHandler ) DeleteDrink( ctx *gin.Context ) {

  drinkId := ctx.Param("id")

  err := dh.drinksTable.ChangeStatus(drinkId, models.DrinksStatuses.Deleted)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "This drink is not public now")
}

func( dh *DrinksHandler ) GetItemTags( ctx *gin.Context ) {
  
  drinkId := ctx.Param("id")
  items, itemsFound, err := dh.drinksTable.GetDrinkTags(drinkId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, models.DrinkTagsCollection {
    ItemsFound: itemsFound,
    Items: items,
  })
}

func( dh *DrinksHandler ) GetItemTag( ctx *gin.Context ) {

  drinkId := ctx.Param("id")
  tagID, err := strconv.Atoi(ctx.Param("tagId"))
  
  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }
  
  tag, err := dh.drinksTable.GetSingleDrinkTag(drinkId, tagID)

  if err != nil {
    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, tag)
}

func( dh *DrinksHandler ) CreateItemTag( ctx *gin.Context ) {
  
  var body models.DrinkTagsPostBody

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  // validate if tag exist
  if _, err := dh.tagsTable.GetSingleTag(body.TagId); err != nil {

    if err == sql.ErrNoRows {
      ctx.JSON(http.StatusNotFound, models.BasicResponse {
        Success: false,
        Data: "This tag doesn't exist",
      })
      return
    }
    utils.ServerError(ctx, err)
    return
  }

  drinkId := ctx.Param("id")

  // validate if the drink exist

  if _, err := dh.drinksTable.GetSingleDrink(drinkId); err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // validate if the tag is already assigned to the drink
  _, err := dh.drinksTable.GetSingleDrinkTag(drinkId, body.TagId) 

  if err == nil {
    models.Conflict(ctx)
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  // add the tag to the drink
  newTag, err := dh.drinksTable.CreateDrinkTag(body, drinkId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.Created(ctx, newTag)
}


func( dh *DrinksHandler ) UpdateItemTag( ctx *gin.Context ) {
    
  drinkId := ctx.Param("id")
  tagId, err := strconv.Atoi(ctx.Param("tagID"))
  var body models.DrinkTagsPostBody

  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  // validate the tag exist 
  if _, err := dh.tagsTable.GetSingleTag(tagId); err != nil {
    
    if err == sql.ErrNoRows {
      ctx.JSON(http.StatusNotFound, models.BasicResponse{
        Success: false,
        Data: "This tag doesn't exist",
      })
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // validate the drink exist
  if _, err := dh.drinksTable.GetSingleDrink(drinkId); err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // validate if the new tag is already assigned to the drink ( by body.tagId )
  _, err = dh.drinksTable.GetSingleDrinkTag(drinkId, body.TagId)
  
  if err == nil {
    ctx.JSON(http.StatusConflict, models.BasicResponse {
      Success: false,
      Data: "The drink already has this tag",
    })
    return
  }

  if err != sql.ErrNoRows {

    utils.ServerError(ctx, err)
    return
  }


  // validate if the drink has the tag to be changed ( by tagID )
  if _, err := dh.drinksTable.GetSingleDrinkTag(drinkId, tagId); err != nil {

    if err == sql.ErrNoRows {
      
      ctx.JSON(http.StatusNotFound, models.BasicResponse {
        Success: true,
        Data: "The drink does not have the tag to be changed",
      })
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // if pass the validations then update the drink tag 

  tag, err := dh.drinksTable.UpdateDrinkTag(body, tagId, drinkId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, tag)
}

func( dh *DrinksHandler ) DeleteItemTag( ctx *gin.Context ) {

  drinkId  := ctx.Param("id") 
  tagId, err := strconv.Atoi(ctx.Param("tagId"))

  if err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  }

  // validate that the drink exist

  _, err = dh.drinksTable.GetSingleDrink(drinkId)

  if err != nil {

    if err == sql.ErrNoRows {

      models.NotFound(ctx)
      return
    }

    models.ServerError(ctx)
    log.Println(err)
    return
  }

  // validate that the tag is actually assigned to the drink
  _, err = dh.drinksTable.GetSingleDrinkTag(drinkId, tagId)

  if err != nil {

    if err == sql.ErrNoRows {

      ctx.JSON(http.StatusNotFound, models.BasicResponse {
        Success: false,
        Data: "The tag is not assigned to this drink",
      })
      return
    }

    models.ServerError(ctx)
    log.Println(err)
    return
  } 

  err = dh.drinksTable.DeleteDrinkTag(tagId, drinkId)

  if err != nil {

    models.ServerError(ctx)
    log.Println(err)
    return
  }

  models.OK(ctx, "Tag removed from the drink")
}
