package db

import "50thbeers/models"


type UsersTable struct {
   db *DB
}

func NewUsersTable( db *DB ) *UsersTable {

   return &UsersTable{
      db: db,
   }
}

func( ut *UsersTable ) GetAllUsers() ([]models.User, int, error) {

   var(
      userId     string
      username   string
      email      string
      password   string
      status     string
      itemsFound int
      data       []models.User
   )

   itemsFound = 0;
   
   rows, err := ut.db.Conn.Query("Select user_id,username,email,password,status from bo_users");

   if err != nil {
      return data, itemsFound, err 
   }

   for rows.Next() {

      err := rows.Scan(
         &userId,
         &username,
         &email,
         &password,
         &status,
      )

      if err != nil {

         return data, itemsFound, err
      }

      itemsFound++
      data = append(data, models.User{
         UserID: userId,
         Username: username,
         Email: email,
         Password: password,
         Status: status,
      })
   }

   return data, itemsFound, nil
}


