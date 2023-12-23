package db

import (
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

type LocationsTable struct {
   db       *DB
   Table    string
   Filters  []models.Filter
}

func NewLocationsTable( db *DB ) *LocationsTable {
   return &LocationsTable{
      db: db,
      Table: "locations",
      Filters: []models.Filter{
         {
            Name: "location_name",
            Type: models.FilterTypes.Like,
         },
         {
            Name: "status",
            Type: models.FilterTypes.EqualString,
         },
         {
            Name: "limit",
            Type: models.FilterTypes.Limit,
            DefaultVal: "10",
         },
         {
           Name: "offset",
           Type: models.FilterTypes.Offset,
           DefaultVal: "0",
         },
         {
           Name: "orderBy",
           Type: models.FilterTypes.OrderBy,
           DefaultVal: "id",
         }, 
         {
           Name: "direction",
           Type: models.FilterTypes.Direction,
           DefaultVal: "ASC",
         },
      },
   }
}

func( lt *LocationsTable ) GetAllLocations( params url.Values ) ( []models.Location, int, error ) {

   var (
      location     models.Location
      updatedAt    sql.NullString
      publicatedAt sql.NullString

      locations    []models.Location
   )

   itemsFound := 0

   whereScript := lt.db.BuildWhere( params, lt.Filters )
   pagScript   := lt.db.BuildPagination(params, lt.Filters)

   // GET COUNT

   countQuery := fmt.Sprintf(
    `
      Select
        count(id)
      From
        %s
      %s
    `,
    lt.Table,
    whereScript,
  )

  err := lt.db.Conn.QueryRow(countQuery).Scan(&itemsFound)
  
  if err != nil {
    return locations, 0, err
  }

  // GET FILTERED DATA
   
  query := fmt.Sprintf(
      `
         Select
            *
         From
            %s
         %s
         %s
      `,
      lt.Table,
      whereScript,
      pagScript,
   )

   rows, err := lt.db.Conn.Query(query)

   if err != nil {
      return locations, itemsFound, err
   }

   for rows.Next() {

      err := rows.Scan(
         &location.LocationId,
         &location.LocationName,
         &location.MapsLink,
         &location.CreatedAt,
         &publicatedAt,
         &updatedAt,
         &location.Comments,
         &location.Status,
      )

      if err != nil {
         return locations, 0, err
      }
      
      location.PublicatedAt = publicatedAt.String
      location.UpdatedAt    = updatedAt.String

      locations = append(locations, location)
   }

   return locations, itemsFound, err
}

func (lt *LocationsTable) GetSingleLocation( locationId string ) ( models.Location, error ) {

  var (
    location     models.Location

    publicatedAt sql.NullString
    updatedAt    sql.NullString
  )

  query := fmt.Sprintf(
    `
      Select
        *
      From %s
      Where id = '%s'
    `,
    lt.Table,
    locationId,
  )

  err := lt.db.Conn.QueryRow(query).Scan(
    &location.LocationId,
    &location.LocationName,
    &location.MapsLink,
    &location.CreatedAt,
    &publicatedAt,
    &updatedAt,
    &location.Comments,
    &location.Status,
  )

  if err != nil {
    return location, err
  }

  location.PublicatedAt = publicatedAt.String
  location.UpdatedAt    = updatedAt.String

  return location, nil
}

func( lt *LocationsTable ) CreateLocation( body models.LocationBody ) ( models.Location, error ) {

  var (
    location     models.Location
    locationId   string
    updatedAt    sql.NullString
    publicatedAt sql.NullString
  )

  locationId = utils.NameToId(body.LocationName)

  query := fmt.Sprintf(
    `
      Insert
      into %s
      (
        id,
        location_name,
        google_maps,
        comments
      )
      Values(
        '%s',
        '%s',
        '%s',
        '%s'
      )
      Returning *
    `,
    lt.Table,
    locationId,
    body.LocationName,
    body.MapsLink,
    body.Comments,
  )

  err := lt.db.Conn.QueryRow(query).Scan(
    &location.LocationId,
    &location.LocationName,
    &location.MapsLink,
    &location.CreatedAt,
    &publicatedAt,
    &updatedAt,
    &location.Comments,
    &location.Status,
  )

  if err != nil {

    return location, err
  }

  location.PublicatedAt = publicatedAt.String
  location.UpdatedAt    = updatedAt.String

  return location, nil
}

func( lt *LocationsTable ) UpdateLocation( body models.LocationPatchBody, locationId string) ( models.Location, error ) {

  var location models.Location

  timestamp := time.Now()
  formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

  // creates the query
  script := ""
  script = lt.db.BuildUpdate(body, models.LocationsTable)
  script += fmt.Sprintf("updated_at = '%s'", formattedTimestamp)

  query := fmt.Sprintf(
    `
      Update %s
      Set
        %s
      Where id = '%s'
    `,
    lt.Table,
    script,
    locationId,
  )

  _, err := lt.db.Conn.Exec(query)

  if err != nil {
    return location, err
  }

  return lt.GetSingleLocation(locationId)
}

func( lt *LocationsTable ) PublicateLocation( locationId string ) ( bool, error ) {

  timestamp := time.Now()
  formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

  query := fmt.Sprintf(
    `
      Update %s
      Set 
        status = '%s',
        publicated_at = '%s',
        updated_at = '%s'
      Where
        id = '%s'
    `,
    lt.Table,
    models.LocationsStatuses.Public,
    formattedTimestamp,
    formattedTimestamp,
    locationId,
  )

  _, err := lt.db.Conn.Exec(query)

  if err != nil {
    return false, err
  }

  return true, nil
  
}

func ( lt *LocationsTable ) HideLocation( locationId string ) ( bool, error ) {

  timestamp := time.Now()
  formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

  query := fmt.Sprintf(
    `
      Update %s
      Set 
        status = '%s',
        updated_at = '%s'
      Where
        id = '%s'
    `,
    lt.Table,
    models.LocationsStatuses.Created,
    formattedTimestamp,
    locationId,
  )

  _, err := lt.db.Conn.Exec(query)

  if err != nil {
    return false, err
  }

  return true, nil
}

func ( lt *LocationsTable ) DeleteLocation( locationId string ) ( bool, error ) {

  timestamp := time.Now()
  formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

  query := fmt.Sprintf(
    `
      Update %s
      Set 
        status = '%s',
        updated_at = '%s'
      Where
        id = '%s'
    `,
    lt.Table,
    models.LocationsStatuses.Deleted,
    formattedTimestamp,
    locationId,
  )

  _, err := lt.db.Conn.Exec(query)

  if err != nil {
    return false, err
  }

  return true, nil
}
