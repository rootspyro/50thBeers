package models

type Tag struct {
   TagId int        `json:"id"`
   TagName string   `json:"tagname"`
   CreatedAt string `json:"createdAt"` 
   UpdatedAt string `json:"updatedAt"`
   Status    string `json:"status"`
}

type TagCollection struct {
   ItemsFound int   `json:"itemsFound"`
   Items      []Tag `json:"items"`
}

type NewTagBody struct {
   TagName string  `json:"tagname"`
}
