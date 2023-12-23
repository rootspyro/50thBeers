package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DrinksHandler struct {
  drinksTable *db.DrinksTable
}

func NewDrinksHandler( dt *db.DrinksTable ) *DrinksHandler {
  return &DrinksHandler{
    drinksTable: dt,
  }
}

func( dh *DrinksHandler ) GetItems( params url.Values ) ( models.DrinkCollection, error ) {
  
  data, itemsFound, err := dh.drinksTable.GetAllDrinks(params)

  var drinks models.DrinkCollection

  if err != nil {
    log.Println(err)
    return drinks, err
  }

  drinks.Items = data
  drinks.ItemsFound = itemsFound
  drinks.Filters = dh.drinksTable.Filters

  return drinks, nil
} 

func( dh *DrinksHandler ) GetItem( drinkId string ) ( models.Drink, error ) {
  return dh.drinksTable.GetSingleDrink(drinkId) 
}

func( dh *DrinksHandler ) CreateItem( body models.DrinkPostBody, drinkId string )  ( models.Drink, error ) {
  
  return dh.drinksTable.CreateDrink(body, drinkId)
}

func( dh *DrinksHandler ) UpdateItem( body models.DrinkPatchBody, drinkId string ) ( models.Drink, error )  {
  return dh.drinksTable.UpdateDrink(body, drinkId)
}

func( dh *DrinksHandler ) ChangeItemStatus( drinkId string, status string ) bool {
  
  success, err := dh.drinksTable.ChangeStatus(drinkId, status)

  if err != nil {
    log.Println(err)
  }

  return success 
}

func( dh *DrinksHandler ) GetItemTags( drinkId string ) ( models.DrinkTagsCollection, error ) {
  
  var drinkTags models.DrinkTagsCollection

  tags, itemsFound, err := dh.drinksTable.GetDrinkTags(drinkId)

  if err != nil {

    log.Println(err)
    return drinkTags, err
  }

  drinkTags.Items = tags
  drinkTags.ItemsFound = itemsFound

  return drinkTags, nil
}

func( dh *DrinksHandler ) GetItemTag( drinkId string, tagId int ) ( models.DrinkTags, error ) {

  return dh.drinksTable.GetSingleDrinkTag(drinkId, tagId)
}

func( dh *DrinksHandler ) CreateItemTag( body models.DrinkTagsPostBody, drinkId string ) ( models.DrinkTags, error ) {

  return dh.drinksTable.CreateDrinkTag(body, drinkId)
}


func( dh *DrinksHandler ) UpdateItemTag( body models.DrinkTagsPostBody, tagId int, drinkId string ) ( models.DrinkTags, error ) {
  return dh.drinksTable.UpdateDrinkTag( body, tagId, drinkId )
}

func( dh *DrinksHandler ) DeleteItemTag( ctx *gin.Context ) {

  drinkId  := ctx.Param("id") 
  strTagId := ctx.Param("tagId")

  tagId, err := strconv.Atoi(strTagId)

  if err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  }

  // validate that the drink exist

  _, err = dh.GetItem(drinkId)

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
  _, err = dh.GetItemTag(drinkId, tagId)

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
