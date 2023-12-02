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
  CreatedAt    string  `json:"createdAt"`
  PublicatedAt string  `json:"publicatedAt"`
  UpdatedAt    string  `json:"updatedAt"`
  Status       string  `json:"status"`
}

type DrinkBody struct {
  DrinkName    string  `json:"drinkName"`
  DrinkType    string  `json:"drinkType"`
  CountryId    int     `json:"countryId"`
  TastingDate  string  `json:"tastingDate"`
  ABV          float32 `json:"abv"`
  Rating       int     `json:"rating"`
  PictureUrl   string  `json:"pictureUrl"`
  LocationId   int     `json:"locationId"`
  Appearance   string  `json:"appearance"`
  Aroma        string  `json:"aroma"`
  Taste        string  `json:"taste"`
  Comments     string  `json:"comments"`
}
