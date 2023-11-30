package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type LocationsTable struct {
   db       *DB
   table    string
   filters  []models.Filter
}

func NewLocationsTable( db *DB ) *LocationsTable {
   return &LocationsTable{
      db: db,
      table: "locations",
      filters: []models.Filter{
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

// This function transforms a Text: "Hello wOrld" 
// to a valid string id "hello_world"
func( lt *LocationsTable ) NameToId( name string ) string {

  name = strings.ToLower(name)
  name = strings.ReplaceAll(name, " ", "_")

  return name
}

func( lt *LocationsTable ) GetAllLocations( params url.Values ) ( []models.Location, int, error ) {

   var (
      location     models.Location
      updatedAt    sql.NullString
      publicatedAt sql.NullString

      locations    []models.Location
   )

   itemsFound := 0

   whereScript := lt.db.BuildWhere( params, lt.filters )
   pagScript   := lt.db.BuildPagination(params, lt.filters)

   // GET COUNT

   countQuery := fmt.Sprintf(
    `
      Select
        count(id)
      From
        %s
      %s
    `,
    lt.table,
    whereScript,
  )

  fmt.Println(countQuery)

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
      lt.table,
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
    lt.table,
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

  locationId = lt.NameToId(body.LocationName)

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
    lt.table,
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

func( lt *LocationsTable ) UpdateLocation( body models.LocationBody, locationId string) ( models.Location, error ) {

  var location models.Location

  timestamp := time.Now()
  formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

  // ------- REFACTOR CODE SECTION INIT ------- // 

  // creates the query
  script := ""
  anotherValueExists := false

  if len(body.LocationName) > 0 {

    script += fmt.Sprintf("location_name = '%s'", body.LocationName)
    anotherValueExists = true
    
  }

  if len(body.MapsLink) > 0 {

    if anotherValueExists {
      script += ", "
    }

    script += fmt.Sprintf("google_maps = '%s'", body.MapsLink) 
    anotherValueExists = true
  }

  if len(body.Comments) > 0 {

    if anotherValueExists {
      script += ", "
    }

    script += fmt.Sprintf("comments = '%s'", body.Comments) 
    anotherValueExists = true
  }
  
  if anotherValueExists {
    script += ", "
  }

  script += fmt.Sprintf("updated_at = '%s'", formattedTimestamp)

  // ------- REFACTOR CODE SECTION END ------- // 

  query := fmt.Sprintf(
    `
      Update %s
      Set
        %s
      Where id = '%s'
    `,
    lt.table,
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
    lt.table,
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
    lt.table,
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
    lt.table,
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
