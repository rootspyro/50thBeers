package models

type Location struct {
   LocationId   string `json:"locationId"`
   LocationName string `json:"locationName"`
   MapsLink     string `json:"mapsLink"`
   CreatedAt    string `json:"createdAt"`
   PublicatedAt string `json:"publicatedAt"`
   UpdatedAt    string `json:"updatedAt"`
   Comments     string `json:"comments"`
   Status       string `json:"status"`
}

type LocationsCollection struct {
   ItemsFound int        `json:"itemsFound"`
   Items      []Location `json:"items"`
   Filters    []Filter   `josn:"avaiableFilters"`
}

type LocationBody struct {
  LocationName string `json:"locationName"`
  MapsLink     string `json:"mapsLink"`
  Comments     string `json:"comments"`
}

