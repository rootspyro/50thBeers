package models

type FILTER_TYPES struct {
   Like        string // name like '% ^'
   EqualString string // name = ''
   EqualNumber string // number = 0'
   Greater     string // age > 0
   Lower       string // age < 100
}

var FilterTypes = FILTER_TYPES {
   Like:          "like",
   EqualString:   "equal to a string",
   EqualNumber:   "equal to a number",
   Greater:       "greater than",
   Lower:         "lower than",
}

type Filter struct {
   Name string 
   Type string
}
