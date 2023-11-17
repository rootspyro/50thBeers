package handlers

import "50thbeers/models"

type UsersHandler struct {

}

func NewUsersHandler() *UsersHandler {
   return &UsersHandler{}
}

func( uh *UsersHandler ) GetItems() models.UserCollection {

   return models.UserCollection{
      ItemsFound: 1,
      Items: []models.User{
         {
            UserID: "hhtttpp",
            Username: "rootspyro",
            Email: "root.spyro@gmail.com",
            Password: "N0tAP4ssw0rdH4sh#",
            Status: "AVAIABLE",
         },
      },
   }
}
