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
  PictureUrl   string  `json:"pictureUrl" binding:"min=10"`
  LocationId   int     `json:"locationId"`
  Appearance   string  `json:"appearance" binding:"required,min=10"`
  Aroma        string  `json:"aroma" binding:"required,min=10"`
  Taste        string  `json:"taste" binding:"required,min=10"`
  Comments     string  `json:"comments" binding:"required,min=10"`
}

type DrinkPatchBody struct {
  DrinkName    string  `json:"drinkName" binding:"min=5"`
  DrinkType    string  `json:"drinkType" binding:"min=5"`
  CountryId    int     `json:"countryId"`
  TastingDate  string  `json:"tastingDate" binding:"min=10,max=10"`
  ABV          float32 `json:"abv"`
  Rating       int     `json:"rating"`
  PictureUrl   string  `json:"pictureUrl" binding:"min=10"`
  LocationId   int     `json:"locationId"`
  Appearance   string  `json:"appearance" binding:"min=10"`
  Aroma        string  `json:"aroma" binding:"min=10"`
  Taste        string  `json:"taste" binding:"min=10"`
  Comments     string  `json:"comments" binding:"min=10"`
}
