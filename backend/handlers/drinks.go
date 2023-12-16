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
