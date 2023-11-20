package models

type USER_STATUSES struct {
   Avaiable  string
   Deleted   string 
}

type TAGS_STATUSES struct {
   Public  string
   Created string
   Deleted string
}

type COUNTRIES_STATUSES struct {
   Public  string
   Created string
   Deleted string
}

var UserStatuses = USER_STATUSES {
   Avaiable: "AVAIABLE",
   Deleted:  "DELETED",
}

var TagsStatuses = TAGS_STATUSES {
   Public:  "PUBLIC",
   Created: "CREATED",
}

var CountriesStatuses = COUNTRIES_STATUSES {
   Public:  "PUBLIC",
   Created: "CREATED",
}
