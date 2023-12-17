package models

type Drink struct {
  DrinkId      string   `json:"drinkId"`
  DrinkName    string   `json:"drinkName"`
  DrinkType    string   `json:"drinkType"`
  CountryName  string   `json:"countryName"`
  TastingDate  string   `json:"tastingDate"`
  ABV          float32  `json:"abv"`
  Rating       int      `json:"rating"`
  PictureUrl   string   `json:"pictureUrl"`
  LocationName string   `json:"locationName"`
  Tags         []string `json:"tags"`
  Appearance   string   `json:"appearance"`
  Aroma        string   `json:"aroma"`
  Taste        string   `json:"taste"`
  Comments     string   `json:"comments"`
  CreatedAt    string   `json:"createdAt"`
  PublicatedAt string   `json:"publicatedAt"`
  UpdatedAt    string   `json:"updatedAt"`
  Status       string   `json:"status"`
}

type DrinkCollection struct {
  ItemsFound int             `json:"itemsFound"`
  Items      []DrinkGeneral  `json:"items"`
  Filters    []Filter        `json:"avaiableFilters"`
}

type DrinkGeneral struct {
  DrinkId      string  `json:"drinkId"`
  DrinkName    string  `json:"drinkName"`
  DrinkType    string  `json:"drinkType"`
  CountryName  string  `json:"countryName"`
  TastingDate  string  `json:"tastingDate"`
  ABV          float32 `json:"abv"`
  Rating       int     `json:"rating"`
  PictureUrl   string  `json:"pictureUrl"`
  LocationName string   `json:"locationName"`
  CreatedAt    string  `json:"createdAt"`
  PublicatedAt string  `json:"publicatedAt"`
  UpdatedAt    string  `json:"updatedAt"`
  Status       string  `json:"status"`
}

type DrinkPostBody struct {
  DrinkName    string  `json:"drinkName" binding:"required,min=5"`
  DrinkType    string  `json:"drinkType" binding:"required,min=5"`
  CountryId    int     `json:"countryId"`
  TastingDate  string  `json:"tastingDate" binding:"required,min=10,max=10"`
  ABV          float32 `json:"abv"`
  Rating       int     `json:"rating" binding:"required"`
  PictureUrl   string  `json:"pictureUrl"`
  LocationId   string  `json:"locationId"`
  Tags         []int   `json:"tags"`
  Appearance   string  `json:"appearance" binding:"required,min=10"`
  Aroma        string  `json:"aroma" binding:"required,min=10"`
  Taste        string  `json:"taste" binding:"required,min=10"`
  Comments     string  `json:"comments"`
}

type DrinkPatchBody struct {
  DrinkName    *string  `json:"drinkName"`
  DrinkType    *string  `json:"drinkType"`
  CountryId    *int     `json:"countryId"`
  TastingDate  *string  `json:"tastingDate"`
  ABV          *float32 `json:"abv"`
  Rating       *int     `json:"rating"`
  PictureUrl   *string  `json:"pictureUrl"`
  LocationId   *string  `json:"locationId"`
  Appearance   *string  `json:"appearance"`
  Aroma        *string  `json:"aroma"`
  Taste        *string  `json:"taste"`
  Comments     *string  `json:"comments"`
}

var DrinksTable = []TableFields {
  {
    StructName: "DrinkId",
    FieldName:  "drink_id",
  },
  {
    StructName: "DrinkName",
    FieldName:  "drink_name",
  },
  {
    StructName: "DrinkType",
    FieldName:  "drink_type",
  },
  {
    StructName: "CountryId",
    FieldName:  "country_id",
  },
  {
    StructName: "TastingDate",
    FieldName:  "tasting_date",
  },
  {
    StructName: "ABV",
    FieldName:  "abv",
  },
  {
    StructName: "Rating",
    FieldName:  "rating",
  },
  {
    StructName: "PictureUrl",
    FieldName:  "picture_url",
  },
  {
    StructName: "LocationId",
    FieldName:  "location_id",
  },
  {
    StructName: "Appearance",
    FieldName:  "appearance",
  },
  {
    StructName: "Aroma",
    FieldName:  "aroma",
  },
  {
    StructName: "Taste",
    FieldName:  "taste",
  },
  {
    StructName: "Comments",
    FieldName:  "comments",
  },
}

type DrinkTags struct {
  TagId   int    `json:"tagId"`
  Tagname string `json:"tagname"`
}

type DrinkTagsCollection struct {
  ItemsFound int         `json:"itemsFound"`
  Items      []DrinkTags `json:"items"`
}

