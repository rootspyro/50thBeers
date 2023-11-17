package models

type User struct {
   UserID   string `json:"id"`
   Username string `json:"username"`
   Email    string `json:"email"`
   Password string `json:"password"`
   Status   string `json:"status"`
}

type UserCollection struct {
   ItemsFound int    `json:"itemsFound"`
   Items      []User `json:"items"`
}
