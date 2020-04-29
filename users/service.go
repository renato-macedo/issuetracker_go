package users

import (
	"errors"
	"log"

	"github.com/labstack/echo"
	"github.com/renato-macedo/issuetracker_go/auth"
	"github.com/renato-macedo/issuetracker_go/database"
	"github.com/renato-macedo/issuetracker_go/roles"
)

// Service has use cases for the user
type Service struct {
	DB *database.DB
}

// NewService return a User Service with the given database
func NewService(db *database.DB) *Service {
	return &Service{
		DB: db,
	}
}

// FindUserByID filter users by the given id and join with roles table
func (s *Service) FindUserByID(id int) (*User, error) {

	//
	q := `SELECT idUsers, name, email, password, role, title, description 
				FROM USERS 
				INNER JOIN ROLES 
				ON ROLES.idRoles = USERS.role 
				WHERE idUsers = ?`

	row := s.DB.Conn.QueryRow(q, id)
	role := new(roles.Role)
	user := new(User)
	user.Role = role
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role.ID, &user.Role.Title, &user.Role.Description)

	return user, err
}

// FindUserByEmail filter users by the given email and join with roles table
func (s *Service) FindUserByEmail(email string) (*User, error) {

	//
	q := `SELECT idUsers, name, email, password, role, title, description 
				FROM USERS 
				INNER JOIN ROLES 
				ON ROLES.idRoles = USERS.role 
				WHERE email = ?`

	row := s.DB.Conn.QueryRow(q, email)
	role := new(roles.Role)
	user := new(User)
	user.Role = role
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role.ID, &user.Role.Title, &user.Role.Description)

	log.Println(err)
	return user, err
}

// CheckCredentials check if the given credentials are valid
func (s *Service) CheckCredentials(l echo.Logger, credentials *LoginDTO) error {
	var hashedPassword string

	err := s.DB.Conn.QueryRow("Select password from USERS where email = ?", credentials.Email).Scan(&hashedPassword)

	if err != nil {
		l.Printf("%v", err)
		return errors.New("user does not exist")
	}

	l.Printf("hashed %v plain %v", hashedPassword, credentials.Password)

	if auth.Compare(hashedPassword, credentials.Password) == true {
		return nil
	}

	return errors.New("Password did not match")

}

// Register insert user in the database
func (s *Service) Register(data *RegisterDTO) error {
	hashedPassword, err := auth.Hash(data.Password)
	if err != nil {
		return err
	}
	q := `INSERT INTO USERS (name, email, password, role) 
				values (?, ?, ?, 2)` // default role is employeee
	_, err = s.DB.Conn.Exec(q, data.Name, data.Email, hashedPassword)
	return err
}
