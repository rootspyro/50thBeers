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
   locations.Filters = lh.locationsTable.Filters

   return locations, nil
}

func( lh *LocationsHandler ) GetItem( locationId string ) ( models.Location, error ) {

  return lh.locationsTable.GetSingleLocation(locationId)

} 

func( lh *LocationsHandler ) CreateItem( body models.LocationBody ) ( models.Location, error ) {

  return lh.locationsTable.CreateLocation(body)
}

func( lh *LocationsHandler) UpdateItem( body models.LocationPatchBody, locationId string ) ( models.Location, error ) {

  return lh.locationsTable.UpdateLocation(body, locationId)
}

func( lh *LocationsHandler ) PublicateItem( locationId string ) bool {

  success, err := lh.locationsTable.PublicateLocation(locationId)
  
  if  !success {
    log.Println(err)
    return false
  }

  return success 
}

func( lh *LocationsHandler ) HideItem( locationId string ) bool {

  success, err := lh.locationsTable.HideLocation(locationId)
  
  if  !success {
    log.Println(err)
    return false
  }

  return success 
}

func( lh *LocationsHandler ) DeleteItem( locationId string ) bool {

  success, err := lh.locationsTable.DeleteLocation(locationId)
  
  if  !success {
    log.Println(err)
    return false
  }

  return success 
}
