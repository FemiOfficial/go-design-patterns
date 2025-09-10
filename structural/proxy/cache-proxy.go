package proxy

import (
	"errors"
	"godesignpatterns/structural/proxy/interfaces"
	"godesignpatterns/structural/proxy/model"

)

type UserListProxy struct {
	Db interfaces.UserList
	StackCache interfaces.UserList
	StackSize int
	DidLastSearchUseCache bool
}

func (userProxy *UserListProxy) FindUser(userId int)  (model.User, error){

	user, err := userProxy.StackCache.FindUser(userId)
	if err == nil {
		userProxy.DidLastSearchUseCache = true
		return user, nil
	}

	// if user is not in cache then get from db  and cache it 
	user, err = userProxy.Db.FindUser(userId)
	if (err == nil) {
		userProxy.AddUserToStack(user)
		return user, nil
	}

	return model.User{}, errors.New("user not found")
}

func (userProxy *UserListProxy) AddUserToStack(user model.User) {
	if len(userProxy.StackCache) >= userProxy.StackSize {
		userProxy.StackCache = userProxy.StackCache[1:]
	}
	userProxy.StackCache = append(userProxy.StackCache, user)
}