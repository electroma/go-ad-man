package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
	log "github.com/Sirupsen/logrus"
)

func Example() {
	const userName = "testUser123"

	DCHost = "172.20.0.149"
	Base = "CN=Users,DC=corp,DC=riglet,DC=io"

	if err := LoginToRemoteAd("CN=Administrator,CN=Users,DC=corp,DC=riglet,DC=io", "XXX"); err != nil {
		log.Errorf("Failed to login %v", err)
	}
	users1, getErr := GetUsers()

	if getErr != nil {
		log.Errorf("Failed to get users %v", getErr)
	}
	log.Infoln(users1)

	if _, ok := users1[userName]; ok {
		DeleteUser(userName)
	}

	info := UserInfo{Name: userName, DisplayName: "Test User", Enabled: true}
	if err := CreateUser(info); err != nil {
		log.Errorf("Failed to create enabled user")
	}
	users2, err := GetUsers();
	if err != nil {
		log.Errorf("Failed to get users %v", err)
	}
	if users2[userName] != info {
		log.Errorf("User does not match %v", users2[userName])
	}

	if err := DeleteUser(userName); err != nil {
		log.Errorf("Failed to delete user %v", err)
	}
}

func TestListMockedAD(t *testing.T) {
	ad := new(mockAD)
	SetLibAd(ad)
	SetAdmGroupName("TestAdm")

	ad.On("GetUsers").Return([]string {"u1", "u2"}, nil)
	ad.On("GetDisabledUsers").Return([]string {"u2"}, nil)
	ad.On("GetUsersInGroup", "TestAdm", false).Return([]string {"u1"}, nil)
	ad.On("GetUserDisplayName", "u1").Return("u u1", nil)
	ad.On("GetUserDisplayName", "u2").Return("u u2", nil)

	users, err := GetUsers()
	ad.AssertExpectations(t)
	assert.Equal(t, map[string]UserInfo {
		"u1": UserInfo{Name: "u1", DisplayName:"u u1", Enabled: true, Admin: true},
		"u2": UserInfo{Name: "u2", DisplayName:"u u2", Enabled: false},
	}, users)
	assert.Nil(t, err)
}

func TestCreateMockedAD(t *testing.T) {
	ad := new(mockAD)
	SetLibAd(ad)
	SetAdmGroupName("TestAdm1")

	ad.On("SearchBase").Return("hz")
	ad.On("CreateUser", "u1", "hz", "u1").Return(nil)
	ad.On("SetUserDisplayName", "u1", "u1 name").Return(nil)
	ad.On("SetUserPassword", "u1", "pwd").Return( nil)
	ad.On("EnableUser", "u1").Return(nil)
	ad.On("GroupAddUser", "TestAdm1", "u1").Return(nil)

	err := CreateUser(UserInfo{"u1", "u1 name", "pwd", true, true})
	assert.Nil(t, err)
	ad.AssertExpectations(t)
}
