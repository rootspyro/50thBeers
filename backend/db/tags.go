package db

import (
	"50thbeers/models"
	"database/sql"
	"fmt"
	"strconv"
)

type TagsTable struct {
   db    *DB
   table string
}

func NewTagsTable( db *DB ) *TagsTable {
   return &TagsTable{
      db: db,
      table: "tags",
   }
}

func ( tt *TagsTable ) GetAllTags() ([]models.Tag, int, error) {

   var (
      tagId     string
      tagName   string
      createdAt string
      updatedAt sql.NullString 
      status    string
      data      []models.Tag
   )

   itemsFound := 0

   query := fmt.Sprintf("Select * from %s where status = 'PUBLIC' order by tagname", tt.table)
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
