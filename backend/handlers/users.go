package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
)

type UsersHandler struct {
   usersTable *db.UsersTable
}

func NewUsersHandler( ut *db.UsersTable ) *UsersHandler {
   return &UsersHandler{
      usersTable: ut,
   }
}

func( uh *UsersHandler ) GetItems() (models.UserCollection, error) {

   data, itemsFound, err := uh.usersTable.GetAllUsers();

   userData := models.UserCollection{
      ItemsFound: 0,
      Items: []models.User{} ,
   }

   if err != nil {
      
      return userData, err
   }

   userData.ItemsFound = itemsFound
   userData.Items = data

   return userData, nil
}
