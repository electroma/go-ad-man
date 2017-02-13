package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
	log "github.com/Sirupsen/logrus"
)

func ExampleWithRealAd() {
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

func TestMockedAD(t *testing.T) {
	ad := new(mockAD)
	SetLibAd(ad)

	ad.On("GetUsers").Return([]string {"u1", "u2"}, nil)
	ad.On("GetDisabledUsers").Return([]string {"u2"}, nil)
	ad.On("GetUserDisplayName", "u1").Return("u u1", nil)
	ad.On("GetUserDisplayName", "u2").Return("u u2", nil)

	users, err := GetUsers()
	ad.AssertExpectations(t)
	assert.Equal(t, map[string]UserInfo {
		"u1": UserInfo{Name: "u1", DisplayName:"u u1", Enabled: true},
		"u2": UserInfo{Name: "u2", DisplayName:"u u2", Enabled: false},
	}, users)
	assert.Nil(t, err)

}
