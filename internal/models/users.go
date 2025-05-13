package models

import (
	"database/sql"
	"time"
)

// Define a new User struct. Notice how the field names and types align
// with the columsn ins the database "users" table?
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Define a new UserModel struct which wraps a database connection pool.
type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert meethod to add a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) error {
	//stmt := `INSERT INTO users (name, email, hashed_password, created)
	//VALUES (?, ?, ?, UTC_TIMESTAMP())`
	//
	//hashedPassword := password
	//
	//_, err := m.DB.Exec(stmt, name, email, hashedPassword)
	//if err != nil {
	//	return err
	//}

	return nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. this will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Exists method to check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
