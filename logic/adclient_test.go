package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWithRealAd(t *testing.T) {
	const userName = "testUser123"

	if err := LoginToRemoteAd("172.20.0.149", "CN=Users,DC=corp,DC=riglet,DC=io", "CN=Administrator,CN=Users,DC=corp,DC=riglet,DC=io", "XXX"); err != nil {
		t.Errorf("Failed to login %v", err)
	}
	users1, getErr := GetUsers()

	if getErr != nil {
		t.Errorf("Failed to get users %v", getErr)
	}
	t.Log(users1)

	if _, ok := users1[userName]; ok {
		DeleteUser(userName)
	}

	info := UserInfo{Name: userName, DisplayName: "Test User", Enabled: true}
	if err := CreateUser(info); err != nil {
		t.Errorf("Failed to create enabled user")
	}
	users2, err := GetUsers();
	if err != nil {
		t.Errorf("Failed to get users %v", err)
	}
	if users2[userName] != info {
		t.Errorf("User does not match %v", users2[userName])
	}

	if err := DeleteUser(userName); err != nil {
		t.Errorf("Failed to delete user %v", err)
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
