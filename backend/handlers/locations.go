package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"
	"net/url"
)

type LocationsHandler struct {
   locationsTable *db.LocationsTable
}

func NewLocationsHandler(lh *db.LocationsTable) *LocationsHandler {
   return &LocationsHandler{
      locationsTable: lh,
   }
}

func( lh *LocationsHandler ) GetItems( params url.Values ) ( models.LocationsCollection, error ) {

   data, itemsFound, err := lh.locationsTable.GetAllLocations(params) 

   var locations models.LocationsCollection

   if err != nil {
      log.Println(err)
      return locations, err
   }

   locations.ItemsFound = itemsFound
   locations.Items = data

   return locations, nil
}
