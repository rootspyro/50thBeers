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
               break
            
            case models.FilterTypes.EqualString:
               whereScript += fmt.Sprintf(" %s = '%s'", filter.Name, params.Get(index))
               break
            }


         }
      }

      counter++

   }

   if len(params) > 0 {
      whereScript = "Where " + whereScript
   }

   return whereScript
}
