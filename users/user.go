package users

import (
	"github.com/renato-macedo/issuetracker_go/roles"
)

// User model
type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"-"`
	Role     *roles.Role `json:"role"`
}

// // UserDTO for request
// type UserDTO struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// LoginDTO for bind data da goes to /auth/login
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterDTO for bind data that goes to /auth/register
type RegisterDTO struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
