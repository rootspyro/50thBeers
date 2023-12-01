package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

type CountriesTable struct {
   db      *DB
   table   string
   Filters []models.Filter
}

func NewCountriesTable( db *DB ) *CountriesTable {
   return &CountriesTable{
      db: db,
      table: "countries",
      Filters: []models.Filter{
         {
            Name: "country_name",
            Type: models.FilterTypes.Like,
         },
         {
            Name: "status",
            Type: models.FilterTypes.EqualString,
         },
      },
   }
}

func (ct *CountriesTable) GetAllCountries( params url.Values ) ( []models.Country, int, error ) {

   var (
      country     models.Country
      updatedAt   sql.NullString

      countries   []models.Country
   )

   itemsFound := 0

   whereScript := ct.db.BuildWhere(params, ct.Filters)

   if whereScript == "" {
      whereScript = fmt.Sprintf("where status = '%s'", models.CountriesStatuses.Default)
   }

   query := fmt.Sprintf(
      `
         Select
            *
         From
            %s
         %s
      `,
      ct.table,
      whereScript,
   )

   rows, err := ct.db.Conn.Query(query)
   
   if err != nil {
      return countries, itemsFound, err
   }

   for rows.Next() {

      itemsFound++ 

      err := rows.Scan(
         &country.CountryId,
         &country.CountryName,
         &country.CreatedAt,
         &updatedAt,
         &country.Status,
      )

      if err != nil {
         return countries, itemsFound, err
      }

      country.UpdatedAt = updatedAt.String

      countries = append(countries, country)
   }

   return countries, itemsFound, err
}

func( ct *CountriesTable ) GetSingleCountry(countryId int) (models.Country, error) {

   var (
      data models.Country
      updatedAt sql.NullString
   )

   query := fmt.Sprintf(
      `
         Select
            *
         from %s
         Where
            id = %d and status = '%s'
      `,
      ct.table,
      countryId,
      models.CountriesStatuses.Default,
   )

   err := ct.db.Conn.QueryRow(query).Scan(
      &data.CountryId,
      &data.CountryName,
      &data.CreatedAt,
      &updatedAt,
      &data.Status,
   )

   data.UpdatedAt = updatedAt.String

   return data, err 

}

func( ct *CountriesTable ) SearchCountryByName(name string) (models.Country, error) {

   var (
      country   models.Country
      updatedAt sql.NullString
   )

   query := fmt.Sprintf(
      `
         Select
            *
         from %s
         Where
            country_name = '%s'
      `,
      ct.table,
      name,
   )

   err := ct.db.Conn.QueryRow(query).Scan(
      &country.CountryId,
      &country.CountryName,
      &country.CreatedAt,
      &updatedAt,
      &country.Status,
   )
   
   country.UpdatedAt = updatedAt.String

   return country, err
}

func( ct *CountriesTable ) CreateCountry(body models.CountryBody) ( models.Country, error ) {

   var(
      country   models.Country
      countryId int
   )

   query := fmt.Sprintf(
      `
         Insert 
         into %s
         (
            country_name
         )
         Values
         (
            '%s'
         )
         Returning id
      `,
      ct.table,
      body.CountryName,
   )

   err := ct.db.Conn.QueryRow(query).Scan(&countryId)

   if err != nil {
      return country, err
   }

   return ct.GetSingleCountry(countryId)
}

func( ct *CountriesTable ) UpdateCountry(body models.CountryBody, countryId int) ( models.Country, error ) {

   var (
      country models.Country
   )
  
   timestamp := time.Now()
   formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

   query := fmt.Sprintf(
      `
         Update
            %s
         Set
            country_name = '%s'
            updated_at = '%s'
         Where
            id = '%d'
      `,
      ct.table,
      body.CountryName,
      formattedTimestamp,
      countryId,
   )

   _, err := ct.db.Conn.Exec(query)

   if err != nil {
      return country, err
   }

   country, err = ct.GetSingleCountry(countryId)  

   return country, err
}

func( ct *CountriesTable ) DeleteCountry(countryId int) ( bool, error ) {

   query := fmt.Sprintf(
      `
         Delete from %s
         Where
            id = '%d'
      `,
      ct.table,
      countryId,
   )

   _, err := ct.db.Conn.Exec(query)

   if err != nil {
      return false, err
   }

   return true, nil
}
