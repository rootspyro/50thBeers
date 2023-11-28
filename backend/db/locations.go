package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
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
