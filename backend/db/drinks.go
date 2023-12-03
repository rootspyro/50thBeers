package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
)

type DrinksTable struct {
  db             *DB
  table          string
  countryTable   string 
  locationsTable string 
  Filters        []models.Filter
}

func NewDrinkTable( db *DB, ct, lt string ) *DrinksTable {
  return &DrinksTable{
    db: db,
    table: "drinks",
    countryTable: ct,
    locationsTable: lt,
    Filters: []models.Filter{
      {
        Name: "limit",
        Type: models.FilterTypes.Like,
        DefaultVal: "10",
      },
      {
        Name: "offset",
        Type: models.FilterTypes.Offset,
        DefaultVal: "0",
      },
      {
        Name: "orderBy",
        Type: models.FilterTypes.Offset,
        DefaultVal: "drink_id",
      },
      {
        Name: "direction",
        Type: models.FilterTypes.Direction,
        DefaultVal: "ASC",
      },
      {
        Name: "drink_name",
        Type: models.FilterTypes.Like,
      },
      {
        Name: "status",
        Type: models.FilterTypes.EqualString,
      },
    },
  }
}

func( dt *DrinksTable ) GetAllDrinks( params url.Values ) ([]models.DrinkGeneral, int, error ) {

  var(

    drinks       []models.DrinkGeneral

    drinkId      string
    drinkName    string
    drinkType    sql.NullString
    countryName  sql.NullString
    tastingDate  string 
    abv          sql.NullFloat64
    rating       sql.NullInt16
    pictureUrl   sql.NullString
    locationName sql.NullString
    createdAt    string 
    updatedAt    sql.NullString
    publicatedAt sql.NullString
    status       string

    itemsFound   int = 0
  )

  whereScript := dt.db.BuildWhere(params, dt.Filters)
  pagScript   := dt.db.BuildPagination(params, dt.Filters)

  countQuery := fmt.Sprintf(
    `
      Select
        count(id)
      From
        %s
      %s
    `,
    dt.table,
    whereScript,
  )

  err := dt.db.Conn.QueryRow(countQuery).Scan(&itemsFound)

  if err != nil {
    return drinks, 0, err
  }

  query := fmt.Sprintf(
    `
      Select
        d.drink_id,
        d.drink_name,
        d.drink_type,
        c.country_name,
        d.tasting_date,
        d.abv,
        d.rating,
        l.location_name,
        d.picture_url,
        d.created_at,
        d.publicated_at,
        d.updated_at,
        d.status
      From
        %s d
      Left Join 
        %s c
      On d.country_id = c.id
      Left Join
        %s l
      On d.location_id = l.id
      %s
      %s
    `,
    dt.table,
    dt.countryTable,
    dt.locationsTable,
    whereScript,
    pagScript,
  )

  rows, err := dt.db.Conn.Query(query)

  if err != nil {
    return drinks, 0, nil
  }

  for rows.Next() {

    err := rows.Scan(
      &drinkId,
      &drinkName,
      &drinkType,
      &countryName,
      &tastingDate,
      &abv,
      &rating,
      &pictureUrl,
      &locationName,
      &createdAt,
      &publicatedAt,
      &updatedAt,
      &status,
    )

    if err != nil {
      return drinks, 0, err
    }

    // Fill the content
    drinks = append(drinks, models.DrinkGeneral{
      DrinkId: drinkId,
      DrinkName: drinkName,
      DrinkType: drinkType.String,
      CountryName: countryName.String,
      TastingDate: tastingDate,
      ABV: float32(abv.Float64),
      Rating: int(rating.Int16),
      PictureUrl: pictureUrl.String,
      LocationName: locationName.String,
      CreatedAt: createdAt,
      PublicatedAt: publicatedAt.String,
      UpdatedAt: updatedAt.String,
      Status: status,
    })
  }

  return drinks, itemsFound, nil
}
