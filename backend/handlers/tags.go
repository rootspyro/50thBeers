package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"
)


type TagsHandler struct {
   tagsTable *db.TagsTable
}

func NewTagsHandler( table *db.TagsTable ) *TagsHandler {
   return &TagsHandler{
      tagsTable: table,
   }
}

func( th *TagsHandler ) GetItems() (models.TagCollection, error) {

   data, itemsFound, err := th.tagsTable.GetAllTags()
   
   var tags models.TagCollection

   if err != nil {
      log.Println(err)
      return tags, err 
   }

   tags = models.TagCollection{
      ItemsFound: itemsFound,
      Items: data,
   }

   return tags, nil
}
