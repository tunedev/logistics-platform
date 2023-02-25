package auth

import (
	"net/http"
	"time"
)

type User struct {
   ID        int
   Username  string
   Email     string
   Password  string
   CreatedAt time.Time
   UpdatedAt time.Time
}

// func Authenticate(username, password string) (*User, error) {
//    // Authenticate the user and return a User struct if successful
// }

func HandleLogin(w http.ResponseWriter, r *http.Request) {
   // Get username and password from request body
   // Call Authenticate function from internal/auth package
   // Return error or user token in response
}


