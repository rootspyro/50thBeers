package db

import (
    _ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
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
