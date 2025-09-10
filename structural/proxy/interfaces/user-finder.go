package interfaces

import (
	"godesignpatterns/structural/proxy/model"
)

type UserFinder interface {
	FindUser(id int) (model.User, error)
}