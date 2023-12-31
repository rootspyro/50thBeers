package models

type USER_STATUSES struct {
   Avaiable  string
   Deleted   string 
   Default   string
}

type TAGS_STATUSES struct {
   Public  string
   Created string
   Default string
}

type COUNTRIES_STATUSES struct {
   Public  string
   Created string
   Default string
}

type LOCATIONS_STATUSES struct {
   Public  string
   Created string
   Deleted string
   Default string
}

type DRINKS_STATUSES struct {
   Public  string
   Created string
   Deleted string
   Default string
}

var UserStatuses = USER_STATUSES {
   Avaiable: "AVAIABLE",
   Deleted:  "DELETED",
   Default:  "AVAIABLE",
}

var TagsStatuses = TAGS_STATUSES {
   Public:  "PUBLIC",
   Created: "CREATED",
   Default: "PUBLIC",
}

var CountriesStatuses = COUNTRIES_STATUSES {
   Public:  "PUBLIC",
   Created: "CREATED",
   Default: "PUBLIC",
}


var LocationsStatuses = LOCATIONS_STATUSES {
  Public:  "PUBLIC",
  Created: "CREATED",
  Deleted: "DELETED",
  Default: "CREATED",
}

var DrinksStatuses = DRINKS_STATUSES {
  Public:  "PUBLIC",
  Created: "CREATED",
  Deleted: "DELETED",
  Default: "CREATED",
}
