package webservice

import "errors"

type User struct {
	ID   int
	Name string
	Role string
}

type Users []User

const (
	NO_USER = "NO SUCH USER"
)

func (u Users) Exists(id int) bool {
	for _, user := range u {
		if user.ID == id {
			return true
		}
	}
	return false
}

func (u Users) FindByName(name string) (User, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New(NO_USER)
}
