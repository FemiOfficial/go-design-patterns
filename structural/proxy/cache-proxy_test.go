package proxy

import (
	"godesignpatterns/structural/proxy/interfaces"
	"godesignpatterns/structural/proxy/model"
	"math/rand"
	"testing"
)

func prepareMockData() UserListProxy {
	userDB := interfaces.UserList{}
	rand.Seed(20033994)
	for i := 0; i<= 1000000; i++ {
		id := rand.Int31()
		userDB = append(userDB, model.User{Id: int(id)})
	}

	userListProxy := UserListProxy{
		Db: userDB,
		StackCache: interfaces.UserList{},
		StackSize: 3,
		DidLastSearchUseCache: false,
	} 

	return userListProxy
}

func TestInvalidUserNotFound(t *testing.T) {

	userListProxy := prepareMockData()

	user, err := userListProxy.FindUser(1000001)
	if err == nil {
		t.Errorf("result should be a user not found error")
	}

	if user != (model.User{}) {
		t.Errorf("User should be an empty user")
	}

	if userListProxy.DidLastSearchUseCache != false {
		t.Errorf("DidLastSearchUseCache should be falsy")
	}
}

func TestUserFoundInCache(t *testing.T)	{
	userListProxy := prepareMockData()

	knowuser := userListProxy.Db[0]

	user, err := userListProxy.FindUser(knowuser.Id)

	if err != nil {
		t.Errorf("result should be found in db")
	}

	if (user.Id != knowuser.Id) {
		t.Errorf("retrieved user should be id = 0")
	}

	if userListProxy.DidLastSearchUseCache != false {
		t.Errorf("DidLastSearchUseCache should be falsy, as user was from db")
	}

	if (userListProxy.StackCache[0].Id != knowuser.Id) {
		t.Errorf("first user in cache should be id = 0")
	}

	user, err = userListProxy.FindUser(knowuser.Id)

	if err != nil {
		t.Errorf("result should be found in cache")
	}

	if userListProxy.DidLastSearchUseCache != true {
		t.Errorf("DidLastSearchUseCache should be truthy, as user was from cache")
	}

}