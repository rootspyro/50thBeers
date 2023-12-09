package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
)

type DrinksTable struct {
  db             *DB
  table          string
  drinkTagsTable string
  tagsTable      string
  countryTable   string 
  locationsTable string 
  Filters        []models.Filter
}

func NewDrinksTable( db *DB, tt, ct, lt string ) *DrinksTable {
  return &DrinksTable{
    db: db,
    table: "drinks",
    drinkTagsTable: "drink_tags",
    tagsTable: tt,
    countryTable: ct,
    locationsTable: lt,
    Filters: []models.Filter{
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
      {
        Name: "tag_id",
        Type: models.FilterTypes.Custom, 
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

    tagId        sql.NullInt16 // tagId var for filter
    itemsFound   int = 0
  )

  whereScript := dt.db.BuildWhere(params, dt.Filters)
  pagScript   := dt.db.BuildPagination(params, dt.Filters)

  // Verify if tag_id filter exist!
  for index := range params {

    if index == "tag_id" {
      
      strTagId, _ := strconv.Atoi(params.Get(index)) 
      tagId = sql.NullInt16{Int16: int16(strTagId), Valid: true}
      break
    }
  }

  countQuery := fmt.Sprintf(
    `
      Select distinct
        count(d.drink_id)
      From
        %s d
    `,
    dt.table,
  )

  if tagId.Int16 > 0 {
    countQuery += fmt.Sprintf(
      `
      Inner Join
        %s dt
      On dt.drink_id = d.drink_id and dt.tag_id = %d
      `,
      dt.drinkTagsTable,
      tagId.Int16,
    )
  }

  countQuery += whereScript
  err := dt.db.Conn.QueryRow(countQuery).Scan(&itemsFound)

  if err != nil {
    return drinks, 0, err
  }

  query := fmt.Sprintf(
    `
      Select distinct
        d.drink_id,
        d.drink_name,
        d.drink_type,
        c.country_name,
        d.tasting_date,
        d.abv,
        d.rating,
        d.picture_url,
        l.location_name,
        d.created_at,
        d.publicated_at,
        d.updated_at,
        d.status
      From
        %s d
      Inner join
        %s dt
      On dt.drink_id = d.drink_id and dt.tag_id = coalesce($1, dt.tag_id) 
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
    dt.drinkTagsTable,
    dt.countryTable,
    dt.locationsTable,
    whereScript,
    pagScript,
  )

  rows, err := dt.db.Conn.Query(query, tagId)

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

func( dt *DrinksTable ) GetSingleDrink( drinkId string ) ( models.Drink, error ) {

  var (
    drink        models.Drink // drink data object

    drinkName    string
    drinkType    sql.NullString
    countryName  sql.NullString
    tastingDate  string 
    abv          sql.NullFloat64
    rating       sql.NullInt16
    pictureUrl   sql.NullString
    locationName sql.NullString
    Tags         []string
    appearance   string
    aroma        string
    taste        string
    comments     sql.NullString
    createdAt    string 
    updatedAt    sql.NullString
    publicatedAt sql.NullString
    status       string

    tagname          string // to read drink tags
  )

  query := fmt.Sprintf(
    `
      Select distinct
        d.drink_id,
        d.drink_name,
        d.drink_type,
        c.country_name,
        d.tasting_date,
        d.abv,
        d.rating,
        d.picture_url,
        l.location_name,
        d.appearance,
        d.aroma,
        d.taste,
        d.comments,
        d.created_at,
        d.publicated_at,
        d.updated_at,
        d.status
      From 
        %s d
      Left Join %s c On d.country_id = c.id
      Left Join %s l On d.location_id = l.id
      Where drink_id = '%s'
    `,
    dt.table,
    dt.countryTable,
    dt.locationsTable,
    drinkId,
  )

  err := dt.db.Conn.QueryRow(query).Scan(
    &drinkId,
    &drinkName,
    &drinkType,
    &countryName,
    &tastingDate,
    &abv,
    &rating,
    &pictureUrl,
    &locationName,
    &appearance,
    &aroma,
    &taste,
    &comments,
    &createdAt,
    &publicatedAt,
    &updatedAt,
    &status,
  )

  if err != nil {
    return drink, err
  }

  // Get Drink Tags
  tagsQuery := fmt.Sprintf(
    `
      Select
        t.tagname
      From
        %s t 
      Join
        %s dt
      On 
        t.id = dt.tag_id and dt.drink_id = '%s'
    `,
    dt.tagsTable,
    dt.drinkTagsTable,
    drinkId,
  ) 

  rows, err := dt.db.Conn.Query(tagsQuery)

  if err != nil {

    return drink, err
  }

  for rows.Next() {
    
    err := rows.Scan(
      &tagname,
    )

    if err != nil {
      return drink, err
    }

    Tags = append(Tags, tagname)
  }

  // build the drink object
  drink.DrinkId = drinkId
  drink.DrinkName = drinkName
  drink.DrinkType = drinkType.String
  drink.CountryName = countryName.String
  drink.TastingDate = tastingDate
  drink.ABV = float32(abv.Float64)
  drink.Rating = int(rating.Int16)
  drink.PictureUrl = pictureUrl.String
  drink.LocationName = locationName.String
  drink.Tags = Tags
  drink.Appearance = appearance
  drink.Aroma = aroma
  drink.Taste = taste
  drink.Comments = comments.String
  drink.CreatedAt = createdAt
  drink.PublicatedAt = publicatedAt.String
  drink.UpdatedAt = updatedAt.String
  drink.Status = status

  return drink, nil
}
