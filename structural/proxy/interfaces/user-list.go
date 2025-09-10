package interfaces

import (
	"errors"
	"godesignpatterns/structural/proxy/model"
)

type UserList []model.User

func (userList *UserList) FindUser (id int) (model.User, error) {

	for i := 0; i < len(*userList) ; i++ {
		if ((*userList)[i].Id == id) {
			return (*userList)[i], nil
		}
	}
		return model.User{}, errors.New("user not found")
}