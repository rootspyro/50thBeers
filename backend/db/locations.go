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

   query := fmt.Sprintf(
      `
         Select
            *
         From
            %s
         %s
      `,
      lt.table,
      whereScript,
   )

   rows, err := lt.db.Conn.Query(query)

   if err != nil {
      return locations, itemsFound, err
   }

   for rows.Next() {

      itemsFound++

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
