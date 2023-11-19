package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"log"
	"net/url"
)


type TagsHandler struct {
   tagsTable *db.TagsTable
}

func NewTagsHandler( table *db.TagsTable ) *TagsHandler {
   return &TagsHandler{
      tagsTable: table,
   }
}

func( th *TagsHandler ) GetItems(params url.Values) (models.TagCollection, error) {

   data, itemsFound, err := th.tagsTable.GetAllTags( params )
   
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
