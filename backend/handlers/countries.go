package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CountriesHandler struct {
   countriesTable *db.CountriesTable
}

func NewCountriesHandler( table *db.CountriesTable ) *CountriesHandler {
   return &CountriesHandler{
      countriesTable: table,
   }
}

func( ch *CountriesHandler ) GetItems(ctx *gin.Context) {
  
  params := ctx.Request.URL.Query()

  data, itemsFound, err := ch.countriesTable.GetAllCountries(params)

  if err != nil {

    utils.ServerError(ctx, err)
    return
  }

  // build response
  countries := models.CountryCollection {
    ItemsFound: itemsFound,
    Items: data,
    Filters: ch.countriesTable.Filters,
  }

  models.OK(ctx, countries)
}

func( ch *CountriesHandler ) GetItem( ctx *gin.Context ) {

  countryId, err := strconv.Atoi(ctx.Param("id"))

  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  data, err := ch.countriesTable.GetSingleCountry(countryId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, data)
}


func( ch *CountriesHandler ) CreateItem( ctx *gin.Context ) {

  var body models.CountryBody

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  // Validate if country actually exist
  _, err := ch.countriesTable.SearchCountryByName(body.CountryName)

  if err == nil {
    models.Conflict(ctx)
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  newCountry, err := ch.countriesTable.CreateCountry(body)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.Created(ctx, newCountry)
}

func( ch *CountriesHandler ) UpdateItem( ctx *gin.Context ) {

  var body      models.CountryBody

  countryId, err := strconv.Atoi(ctx.Param("id"))

  // get and validate params and body
  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  // country validation
  _, err = ch.countriesTable.GetSingleCountry(countryId)

  if err != nil {
    
    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // if the country actually exist -> validate that the new country name is unique
  _, err = ch.countriesTable.SearchCountryByName(body.CountryName)

  if err != nil {
    
    if err == sql.ErrNoRows {

      country, err := ch.countriesTable.UpdateCountry(body, countryId)

      if err != nil {
        utils.ServerError(ctx, err)
        return
      }

      models.OK(ctx, country)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  models.Conflict(ctx)
  return
}


func( ch *CountriesHandler ) DeleteItem( ctx *gin.Context ) {

  countryId, err := strconv.Atoi(ctx.Param("id"))

  if err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  }

  _, err = ch.countriesTable.GetSingleCountry(countryId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  err = ch.countriesTable.DeleteCountry(countryId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "Country successfully deleted")
}
