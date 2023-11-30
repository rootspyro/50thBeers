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
}

var FilterTypes = FILTER_TYPES {
   Like:          "like",
   EqualString:   "equal to a string",
   EqualNumber:   "equal to a number",
   Greater:       "greater than",
   Lower:         "lower than",
   Limit:         "limit",
   Offset:        "offset",
   OrderBy:       "orderBy",
   Direction:     "ASC or DESC",
}

type Filter struct {
   Name        string 
   Type        string
   DefaultVal  string 
}
