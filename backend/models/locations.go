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
   Filters    []Filter   `json:"avaiableFilters"`
}

type LocationBody struct {
  LocationName string `json:"locationName" binding:"required"`
  MapsLink     string `json:"mapsLink" binding:"required"`
  Comments     string `json:"comments"`
}

type LocationPatchBody struct {
  LocationName *string `json:"locationName"`
  MapsLink     *string `json:"mapsLink"`
  Comments     *string `json:"comments"`
}

var LocationsTable = []TableFields {
  {
    StructName: "LocationId",
    FieldName: "id",
  },
  {
    StructName: "LocationName",
    FieldName: "location_name",
  },
  {
    StructName: "MapsLink",
    FieldName: "google_maps",
  },
  {
    StructName: "Comments",
    FieldName: "comments",
  },
}
