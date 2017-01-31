package logic

import (
	"github.com/paleg/libadclient"
	log "github.com/Sirupsen/logrus"
)

type UserInfo struct {
	Name        string
	DisplayName string
	Enabled     bool
}

type AdOprations interface {
	EnableUser(user string) (err error)
	DisableUser(user string) (err error)
	SetUserDisplayName(user string, displayName string) (err error)
	Login(_params adclient.ADConnParams) (err error)
	GetUsers() ([]string, error)
	GetDisabledUsers() ([]string, error)
	CreateUser(cn string, container string, user_short string) (err error)
	DeleteDN(dn string) (err error)
	GetUserDisplayName(name string) (result string, err error)
	SearchBase() (result string)
}

var _libad AdOprations

func SetLibAd(ad AdOprations) {
	_libad = ad
}

func GetUsers() (map[string]UserInfo, error) {

	users, err1 := _libad.GetUsers();
	if err1 != nil {
		log.Errorf("Failed to fetch users: %v", err1)
		return nil, err1
	}

	disabledUsers, err2 := _libad.GetDisabledUsers()

	if err2 != nil {
		log.Errorf("Failed to fetch disabled users: %v", err2)
		return nil, err2
	}

	userInfoIdx := map[string]UserInfo{}
	for _, userName := range users {
		userDisplayName, _ := _libad.GetUserDisplayName(userName)
		userInfoIdx[userName] = UserInfo{Name: userName, DisplayName: userDisplayName, Enabled: true }
	}

	for _, userName := range disabledUsers {
		if userRecord, ok := userInfoIdx[userName]; ok {
			log.Debugf("User %v is disabled", userName)
			userRecord.Enabled = false
			userInfoIdx[userName] = userRecord
		}
	}

	return userInfoIdx, nil

}

func CreateUser(info UserInfo) error {
	if err := _libad.CreateUser(info.Name, _libad.SearchBase(), info.Name); err != nil {
		return err
	}
	if err := _libad.SetUserDisplayName(info.Name, info.DisplayName); err != nil {
		return err
	}
	if (info.Enabled) {
		if err := _libad.EnableUser(info.Name); err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(name string) error {
	return _libad.DeleteDN("CN=" + name + "," + _libad.SearchBase())
}
