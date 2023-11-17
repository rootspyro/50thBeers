package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
)

type UsersHandler struct {
   db *db.DB
}

func NewUsersHandler( db *db.DB ) *UsersHandler {
   return &UsersHandler{
      db: db,
   }
}

func( uh *UsersHandler ) GetItems() (models.UserCollection, error) {

   userData := models.UserCollection{
      ItemsFound: 0,
      Items: []models.User{} ,
   }

   var(
      userId   string
      username string
      email    string
      password string
      status   string
   )

   rows, err := uh.db.Conn.Query("Select user_id,username,email,password,status from bo_users");
   
   if err != nil {
      return userData, err
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
         return userData, nil
      }

      user := models.User{
         UserID: userId,
         Username: username,
         Email: email,
         Password: password,
         Status: status,
      }

      userData.ItemsFound++
      userData.Items = append(userData.Items, user)

   }
   //
   // data := models.UserCollection{
   //    ItemsFound: 1,
   //    Items: []models.User{
   //       {
   //          UserID: "hhtttpp",
   //          Username: "rootspyro",
   //          Email: "root.spyro@gmail.com",
   //          Password: "N0tAP4ssw0rdH4sh#",
   //          Status: "AVAIABLE",
   //       },
   //    },
   // }

   return userData, nil
}
