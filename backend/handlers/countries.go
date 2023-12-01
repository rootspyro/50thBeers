package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"
	"net/url"
	"strconv"
)

type CountriesHandler struct {
   countriesTable *db.CountriesTable
}

func NewCountriesHandler( table *db.CountriesTable ) *CountriesHandler {
   return &CountriesHandler{
      countriesTable: table,
   }
}

func( ch *CountriesHandler ) GetItems(params url.Values) ( models.CountryCollection, error ) {

   data, itemsFound, err := ch.countriesTable.GetAllCountries( params )

   var countries models.CountryCollection

   if err != nil {
      log.Println(err)
      return countries, err
   }

   countries.ItemsFound = itemsFound
   countries.Items = data
   countries.Filters = ch.countriesTable.Filters

   return countries, err
}

func( ch *CountriesHandler ) GetItem( countryId string ) ( models.Country, error ) {

   countryIdInt, _ := strconv.Atoi(countryId)
   return ch.countriesTable.GetSingleCountry(countryIdInt)
}

func( ch *CountriesHandler ) SearchItemByName( name string ) ( models.Country, error ) {
   
   return ch.countriesTable.SearchCountryByName(name)
}

func( ch *CountriesHandler ) CreateItem( body models.CountryBody ) ( models.Country, error ) {

   return ch.countriesTable.CreateCountry(body)
}

func( ch *CountriesHandler ) UpdateItem( countryId string, body models.CountryBody ) ( models.Country, error ) {

   countryIdInt, _ := strconv.Atoi(countryId)
   return ch.countriesTable.UpdateCountry(body, countryIdInt)
}

func( ch *CountriesHandler ) DeleteItem(countryId string) bool {

   countryIdInt, _ := strconv.Atoi(countryId)
   success, err := ch.countriesTable.DeleteCountry(countryIdInt)

   if err != nil {
      log.Println(err)
   }

   return success
}
