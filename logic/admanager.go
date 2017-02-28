package logic

import (
	"github.com/paleg/libadclient"
	log "github.com/Sirupsen/logrus"
)

// UserInfo DTO containing user information.
// Used both as input and output
type UserInfo struct {
	Name        string
	DisplayName string
	Password    string
	Enabled     bool
	Admin       bool
}


// AdOprations provides high-level API to manage users in AD domain
type AdOprations interface {
	EnableUser(user string) (err error)
	DisableUser(user string) (err error)
	SetUserDisplayName(user string, displayName string) (err error)
	Login(params adclient.ADConnParams) (err error)
	GetUsers() ([]string, error)
	GetDisabledUsers() ([]string, error)
	CreateUser(cn string, container string, userShort string) (err error)
	DeleteDN(dn string) (err error)
	GetUserDisplayName(name string) (result string, err error)
	SearchBase() (result string)
	SetUserPassword(user string, password string) (err error)
	GroupAddUser(group string, user string) (err error)
	GetUserGroups(user string, nested bool) (result []string, err error)
	GroupRemoveUser(user string, group string) (err error)
	GetUsersInGroup(group string, nested bool) (result []string, err error)
	GetGroups() (groups []string, err error)
}

var _libad AdOprations
var admGroupName = "Domain Admins"

// SetLibAd allows to override low level AD API provider
func SetLibAd(ad AdOprations) {
	_libad = ad
}

// SetAdmGroupName sets domain-specific "Administrators" group
func SetAdmGroupName(adm string) {
	admGroupName = adm
}

// GetUsers returns list of users existing in AD
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

	adminUsers, err3 := _libad.GetUsersInGroup(admGroupName, false)

	if err3 != nil {
		log.Errorf("Failed to fetch admin users: %v", err3)
		groups, err4 := _libad.GetGroups()
		if err4 == nil {
			log.Errorf("Available groups are %v", groups)
		}
		return nil, err3
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

	for _, userName := range adminUsers {
		if userRecord, ok := userInfoIdx[userName]; ok {
			log.Debugf("User %v is admin account", userName)
			userRecord.Admin = true
			userInfoIdx[userName] = userRecord
		}
	}

	return userInfoIdx, nil

}

// CreateUser allows to add new user to AD
func CreateUser(info UserInfo) error {
	if err := _libad.CreateUser(info.Name, _libad.SearchBase(), info.Name); err != nil {
		return err
	}
	if err := _libad.SetUserDisplayName(info.Name, info.DisplayName); err != nil {
		return err
	}
	if err := _libad.SetUserPassword(info.Name, info.Password); err != nil {
		return err
	}
	if (info.Enabled) {
		if err := _libad.EnableUser(info.Name); err != nil {
			return err
		}
	}
	if (info.Admin) {

		if err := _libad.GroupAddUser(admGroupName, info.Name); err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser allows to delete user object from AD
// Important: It is destructive operation
func DeleteUser(name string) error {
	return _libad.DeleteDN("CN=" + name + "," + _libad.SearchBase())
}
