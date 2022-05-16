package database

import (
	"errors"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()

	if err != nil {
		return User{}, err
	}

	if _, ok := db.Users[email]; ok {
		return User{}, errors.New("email already exists")
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}

	db.Users[email] = user
	err = c.updateDB(db)

	if err != nil {
		return User{}, err
	}

	return user, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()

	if err != nil {
		return User{}, err
	}

	user, ok := db.Users[email]

	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	user.Email = email
	user.Password = password
	user.Name = name
	user.Age = age

	db.Users[email] = user
	err = c.updateDB(db)

	if err != nil {
		return User{}, err
	}

	return user, err
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()

	if err != nil {
		return User{}, err
	}

	user, ok := db.Users[email]

	if !ok {
		return User{}, errors.New("User not found!")
	}

	return user, nil
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()

	if err != nil {
		return err
	}

	delete(db.Users, email)
	err = c.updateDB(db)

	if err != nil {
		return err
	}

	return err
}
