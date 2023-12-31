package handlers

import (
	"50thbeers/db"
	"50thbeers/models"
	"50thbeers/utils"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)


type TagsHandler struct {
   tagsTable *db.TagsTable
}

func NewTagsHandler( table *db.TagsTable ) *TagsHandler {
   return &TagsHandler{
      tagsTable: table,
   }
}

func( th *TagsHandler ) GetItems( ctx *gin.Context ) {

  params := ctx.Request.URL.Query()

  data, itemsFound, err := th.tagsTable.GetAllTags(params)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  tags := models.TagCollection {
    ItemsFound: itemsFound,
    Items: data,
    Filters: th.tagsTable.Filters,
  }

  models.OK(ctx, tags)
}

func( th *TagsHandler ) GetItem( ctx *gin.Context ) {

  tagId, err := strconv.Atoi(ctx.Param("id"))

  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  tag, err := th.tagsTable.GetSingleTag(tagId) 

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, tag)

}

func( th *TagsHandler ) CreateItem( ctx *gin.Context ) {

  var body models.TagBody

  if err := ctx.ShouldBindJSON(&body); err != nil {
    models.InvalidRequest(ctx, err.Error())
  }

  // validate that the tag to create doesn't exist already 

  _, err := th.tagsTable.SearchTagByName(body.TagName)

  if err == nil {
    models.Conflict(ctx)
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  newTag, err := th.tagsTable.CreateTag(body) 

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.Created(ctx, newTag)
}

func( th *TagsHandler ) UpdateItem( ctx *gin.Context ) {

  var body models.TagBody

  tagId, err := strconv.Atoi(ctx.Param("id"))

  if err != nil {
    models.InvalidRequest(ctx, err.Error())
    return
  }

  if err := ctx.ShouldBindJSON(&body); err != nil {

    models.InvalidRequest(ctx, err.Error())
    return
  }

  // validate the tag actually exist
  _, err = th.tagsTable.GetSingleTag(tagId)

  if err != nil {

    if err == sql.ErrNoRows {
      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // validate that the tag is unique
  _, err = th.tagsTable.SearchTagByName(body.TagName)

  if err == nil {
    models.Conflict(ctx)
    return
  }

  if err != sql.ErrNoRows {
    utils.ServerError(ctx, err)
    return
  }

  // Updates the tag

  tag, err := th.tagsTable.UpdateTag(body, tagId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, tag)

}

func( th *TagsHandler ) DeleteItem( ctx *gin.Context ) {

  tagId, err := strconv.Atoi(ctx.Param("id"))

  if err != nil {
    models.InvalidRequest(ctx, err)
    return
  }

  // validate the tag exist

  _, err = th.tagsTable.GetSingleTag(tagId)

  if err != nil {

    if err == sql.ErrNoRows {

      models.NotFound(ctx)
      return
    }

    utils.ServerError(ctx, err)
    return
  }

  // Deletes the tag
  err = th.tagsTable.DeleteTag(tagId)

  if err != nil {
    utils.ServerError(ctx, err)
    return
  }

  models.OK(ctx, "Tag successfully deleted!")
}
