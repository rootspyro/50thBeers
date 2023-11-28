package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)


type DB struct {
   host     string
   username string
   password string
   dbname   string
   port     string 
   Conn     *sql.DB  
}

func NewDBConnection(host, username, password, dbname, port string) *DB {

   psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
      host,
      port,
      username,
      password,
      dbname,
   )

   log.Println(psqlInfo)

   conn, err := sql.Open("postgres", psqlInfo)
   
   if err != nil {
      conn.Close()
      log.Fatalln("Database connection error!")
   }

   return &DB{
      host: host, 
      username: username,
      password: password,
      dbname: dbname,
      port: port,
      Conn: conn,
   }
}

func ( db *DB ) BuildWhere( params url.Values, filters []models.Filter ) string {

  whereScript := ""
  counter := 0

  for index := range params {

    for _, filter := range filters {

      if index == filter.Name {

        if counter > 0 {
           whereScript += " and"
        }

        switch filter.Type {

          case models.FilterTypes.Like:
             whereScript += fmt.Sprintf(" %s like '%%%s%%'", filter.Name, params.Get(index))
             counter++
             break
          
          case models.FilterTypes.EqualString:
             whereScript += fmt.Sprintf(" %s = '%s'", filter.Name, params.Get(index))
             counter++
             break
        }
      }
    }
  }
   
  if whereScript != "" {
     whereScript = "Where " + whereScript
  }

  return whereScript
}

func (db *DB) BuildPagination( params url.Values, filters []models.Filter ) string {

  var script, limitScript, offsetScript, orderByScript, direction string

  for index := range params {
    for _, filter := range filters {
      if index == filter.Name {
        switch filter.Type {
          case models.FilterTypes.Limit:
            limitScript = fmt.Sprintf(" Limit %s", params.Get(index))
            break

          case models.FilterTypes.Offset:
            offsetScript = fmt.Sprintf(" Offset %s", params.Get(index))
            break

          case models.FilterTypes.OrderBy:
            orderByScript = fmt.Sprintf(" Order By %s", params.Get(index))
            break

          case models.FilterTypes.Direction:
            direction = fmt.Sprintf(" %s", params.Get(index))
            break
        } 
      }
    }
  }

  // Default Values

  for _, filter := range filters {

    switch filter.Type {
      case models.FilterTypes.Limit:

        if len(limitScript) == 0 {

          limitScript = fmt.Sprintf("Limit %s", filter.DefaultVal)
        }

        break

      case models.FilterTypes.Offset:

        if len(offsetScript) == 0 {

          offsetScript = fmt.Sprintf("Offset %s", filter.DefaultVal) 
        }
        break

      case models.FilterTypes.OrderBy:

        if len(orderByScript) == 0 {
          orderByScript = fmt.Sprintf("Order By %s", filter.DefaultVal)
        }
        break

      case models.FilterTypes.Direction:

        if len(direction) == 0 {
          direction = fmt.Sprintf(" %s", filter.DefaultVal)
        }
        break
    }
  }

  script = fmt.Sprintf("%s %s %s %s", orderByScript, direction, limitScript, offsetScript)
  return script
}
