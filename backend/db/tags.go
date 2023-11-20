package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"time"
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

   whereScript := tt.db.BuildWhere(params, tt.filters)

   if whereScript == "" {
      whereScript = fmt.Sprintf("where status = '%s'", models.TagsStatuses.Public)
   }

   query := fmt.Sprintf(
      `
         Select  
            *
         from %s 
         %s
         order by tagname
      `, 
      tt.table,
      whereScript,
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

func( tt *TagsTable ) SearchTagByName(name string) (models.Tag, error) {

   var tag models.Tag
   var updatedAt sql.NullString

   query := fmt.Sprintf(
      `
         Select
            *
         from %s
         Where
            tagname = '%s'
      `,
      tt.table,
      name,
   )

   err := tt.db.Conn.QueryRow(query).Scan(
      &tag.TagId,
      &tag.TagName,
      &tag.CreatedAt,
      &updatedAt,
      &tag.Status,
   )

   return tag, err

}

func( tt *TagsTable ) CreateTag( data models.TagBody ) (models.Tag, error) {

   var response models.Tag

   query := fmt.Sprintf(
      `
         Insert into %s
         (
            tagname
         )
         Values
         (
            '%s'
         )
         Returning id
      `,
      tt.table,
      data.TagName,
   )

   var tagId string

   err := tt.db.Conn.QueryRow(query).Scan(&tagId)

   if err != nil {
      return response, err
   }

   tagIdInt, _ := strconv.Atoi(tagId)

   response, err = tt.GetSingleTag(tagIdInt) 

   return response, err 
}

func( tt *TagsTable ) UpdateTag( data models.TagBody, tagId int ) (models.Tag, error){

   var (
      tag models.Tag
   )

   timestamp := time.Now()
   formattedTimestamp := timestamp.Format("2006-01-02 15:04:05")

   query := fmt.Sprintf(
      `
         Update
            %s
         Set
            tagname = '%s',
            updated_at = '%s'
         Where
            id = %d
      `,
      tt.table,
      data.TagName,
      formattedTimestamp,
      tagId,
   )

   _, err := tt.db.Conn.Exec(query)

   if err != nil {
      return tag, err 
   }
   
   tag, err = tt.GetSingleTag(tagId)
   
   return tag, err 
} 

func(tt *TagsTable) DeleteTag(tagId int) ( bool, error ) {

   query := fmt.Sprintf(
      `
         Delete from
            %s
         Where
            id = '%d'
      `,
      tt.table,
      tagId,
   )

   _, err := tt.db.Conn.Exec(query)

   if err != nil {

      return false, err
   }

   return true, nil
}
