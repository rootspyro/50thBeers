package models

type Country struct {
   CountryId   int    `json:"countryId"`
   CountryName string `json:"countryName"`
   CreatedAt   string `json:"createdAt"`
   UpdatedAt   string `json:"updatedAt"`
   Status      string `json:"status"`
}

type CountryCollection struct {
   ItemsFound int       `json:"itemsFound"`
   Items      []Country `json:"items"`
   Filters    []Filter  `json:"avaiableFilters"` 
}

type CountryBody struct {
   CountryName string `json:"countryName" binding:"required,min=3"`
}
