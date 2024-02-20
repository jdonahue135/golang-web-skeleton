package dbrepo

import (
	"errors"

	"github.com/jdonahue135/golang-web-skeleton/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "jack@nimble.com" {
		return 0, "", errors.New("some error")
	}
	return 1, "", nil
}
