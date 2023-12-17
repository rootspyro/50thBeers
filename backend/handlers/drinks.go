package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"
	"net/url"
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

  return dh.drinksTable.GetDrinkSingleTag(drinkId, tagId)
}

func( dh *DrinksHandler ) CreateItemTag( body models.DrinkTagsPostBody, drinkId string ) ( models.DrinkTags, error ) {

  return dh.drinksTable.CreateDrinkTag(body, drinkId)
}

