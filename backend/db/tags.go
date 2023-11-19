package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
)

type TagsTable struct {
   db      *DB
   table   string
   filters []models.Filter
}

func NewTagsTable( db *DB ) *TagsTable {
   return &TagsTable{
      db: db,
      table: "tags",
      filters: []models.Filter{
         {
            Name: "tagname",
            Type: models.FilterTypes.Like,
         },
         {
            Name: "status",
            Type: models.FilterTypes.EqualString,
         },
      },
   }
}

func ( tt *TagsTable ) GetAllTags( params url.Values ) ([]models.Tag, int, error) {

   var (
      tagId     string
      tagName   string
      createdAt string
      updatedAt sql.NullString 
      status    string
      data      []models.Tag
   )

   itemsFound := 0

   filters := tt.db.BuildWhere(params, tt.filters)

   query := fmt.Sprintf(
      `
         Select  
            *
         from %s 
         %s
         order by tagname
      `, 
      tt.table,
      filters,
   )
   rows, err := tt.db.Conn.Query(query)

   if err != nil {
      return  data, 0, err
   }

   for rows.Next() {

      err := rows.Scan(
         &tagId,
         &tagName,
         &createdAt,
         &updatedAt,
         &status,
      )

      if err != nil {

         return data, 0, err
      }

      itemsFound++

      tagIdInt, _ := strconv.Atoi(tagId)

      data = append(data, models.Tag{
         TagId:     tagIdInt,
         TagName:   tagName,
         CreatedAt: createdAt,
         UpdatedAt: updatedAt.String,
         Status:    status,
      })
   }

   return data, itemsFound, nil
}

func( tt *TagsTable ) GetSingleTag( tagId int ) (models.Tag, error) {

   var data models.Tag
   var updatedAt sql.NullString

   query := fmt.Sprintf(
      `
         Select 
            *
         from %s
         where
            id = %d
      `,
      tt.table,
      tagId,
   )

   err := tt.db.Conn.QueryRow(query).Scan(
      &data.TagId,
      &data.TagName,
      &data.CreatedAt,
      &updatedAt,
      &data.Status,
   )

   if err != nil {
      return data, err
   }

   data.UpdatedAt = updatedAt.String

   return data, nil
}
