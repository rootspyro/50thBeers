package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type LocationsHandler struct {
   locationsTable *db.LocationsTable
}

func NewLocationsHandler(lh *db.LocationsTable) *LocationsHandler {
   return &LocationsHandler{
      locationsTable: lh,
   }
}

func( lh *LocationsHandler ) GetItems( ctx *gin.Context ) {

  params := ctx.Request.URL.Query()

  data, itemsFound, err := lh.locationsTable.GetAllLocations(params)

  if err != nil {
    utils.ServerError(ctx, err)
  }

  models.OK(ctx, models.LocationsCollection{
    ItemsFound: itemsFound,
    Items: data,
    Filters: lh.locationsTable.Filters,
  })
}

func( lh *LocationsHandler ) GetItem( ctx *gin.Context ) {

  locationId := ctx.Param("id")

  data, err := lh.locationsTable.GetSingleLocation(locationId)

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

func( lh *LocationsHandler ) CreateItem( ctx *gin.Context ) {
  
  var body models.LocationBody

  if err := ctx.ShouldBindJSON(&body); err != nil {
    
    models.InvalidRequest(ctx, err)
    return
  }

  newId := utils.NameToId(body.LocationName)

  // search if the item already exist
  _, err := lh.locationsTable.GetSingleLocation(newId)

  if err == nil {
    models.Conflict(ctx)
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  location, err := lh.locationsTable.CreateLocation(body)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.Created(ctx, location)
}

func( lh *LocationsHandler) UpdateItem( ctx *gin.Context ) {
  
  locationId := ctx.Param("id")

  var body models.LocationPatchBody

  if err := ctx.ShouldBindJSON(&body); err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  } 

  // validate if the location actually exist
  _, err := lh.locationsTable.GetSingleLocation(locationId)
  
  if err != nil {
    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // if item exist then update the data
  location, err := lh.locationsTable.UpdateLocation(body, locationId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, location)
}

func( lh *LocationsHandler ) PublishItem( ctx *gin.Context ) {

  locationId := ctx.Param("id")

  // validate if location exist
  _, err := lh.locationsTable.GetSingleLocation(locationId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  err = lh.locationsTable.ChangeStatus(locationId, models.LocationsStatuses.Public)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "Locations successfully published")
    
}

func( lh *LocationsHandler ) HideItem( ctx *gin.Context )  {

  locationId := ctx.Param("id")

  // validate if location exist
  _, err := lh.locationsTable.GetSingleLocation(locationId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  err = lh.locationsTable.ChangeStatus(locationId, models.LocationsStatuses.Created)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "This location is not public now")
  
}

func( lh *LocationsHandler ) DeleteItem( ctx *gin.Context ) {
  
  locationId := ctx.Param("id")

  // validate if location exist
  _, err := lh.locationsTable.GetSingleLocation(locationId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  err = lh.locationsTable.ChangeStatus(locationId, models.LocationsStatuses.Deleted)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "Location successfully deleted")

}
