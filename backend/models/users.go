package models

type User struct {
   UserID    string `json:"id"`
   Username  string `json:"username"`
   Email     string `json:"email"`
   Password  string `json:"password"`
   CreatedAt string `json:"createdAt"`
   UpdatedAt string `json:"updatedAt"`
   Status    string `json:"status"`
}

type UserCollection struct {
   ItemsFound int    `json:"itemsFound"`
   Items      []User `json:"items"`
}

type LoginResponse struct {
   Token string `json:"token"`
   Name  string `json:"name"`
   Sub   string `json:"sub"`
   Exp   int64    `json:"exp"`
}

type LoginBody struct {
   User     string `json:"user"` // it can be the username or email
   Password string `json:"password"`
}
