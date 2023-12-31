package models

type FILTER_TYPES struct {
   Like        string // name like '% ^'
   EqualString string // name = ''
   EqualNumber string // number = 0'
   Greater     string // age > 0
   Lower       string // age < 100
   Limit       string // Limit  8
   Offset      string // Offset 8
   OrderBy     string // OrderBy name
   Direction   string // ASC or DESC
   Custom      string
}

var FilterTypes = FILTER_TYPES {
   Like:          "Like",
   EqualString:   "Equal",
   EqualNumber:   "Equal",
   Greater:       "Greater Than",
   Lower:         "Lower than",
   Limit:         "Limit",
   Offset:        "Offset",
   OrderBy:       "Order By",
   Direction:     "ASC or DESC",
   Custom:        "Reserved for customized filters",
}

type Filter struct {
  Name        string `json:"name"`
  Type        string `json:"type"`
  DefaultVal  string `json:"defaultValue"`
}

// this object specify the relation between the Models structs and tables
type TableFields struct {
  StructName string
  FieldName  string
}
