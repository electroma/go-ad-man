package logic

import (
	"github.com/paleg/libadclient"
	"github.com/Sirupsen/logrus"
	"github.com/kataras/go-errors"
)

type remoteAdWrapper struct {
}

// Base sets base DN to be used by user operations
var Base string
// DCHost is AD DC hostname or IP
var DCHost string

// LoginToRemoteAd opens LDAP connection to remote AD server
func LoginToRemoteAd(bindDn string, bindPass string) (err error) {
	if len(DCHost) == 0 {
		return errors.New("Please specify DC host first")
	}

	if len(Base) == 0 {
		return errors.New("Please specify base-DN first")
	}


	adclient.New()
	params := adclient.DefaultADConnParams()
	params.Uries = append(params.Uries, adclient.LdapPrefix()+DCHost)
	params.Search_base = Base
	params.Binddn = bindDn;
	params.Bindpw = bindPass;
	params.Secured = false;
	params.UseGSSAPI = false
	params.Timelimit = 60
	params.Nettimeout = 60

	if err := adclient.Login(params); err != nil {
		logrus.Errorf("Failed to AD login: %v\n", err)
		return err
	}
	SetLibAd(remoteAdWrapper{})
	return
}

func (rad remoteAdWrapper) DisableUser(user string) (err error) {
	return adclient.DisableUser(user)
}

func (rad remoteAdWrapper) EnableUser(user string) (err error) {
	return adclient.EnableUser(user)
}


func (rad remoteAdWrapper) SetUserDisplayName(user string, displayName string) (err error) {
	return adclient.SetUserDisplayName(user, displayName)
}

func (rad remoteAdWrapper) SearchBase() (result string) {
	return adclient.SearchBase()
}

func (rad remoteAdWrapper) Login(_params adclient.ADConnParams) (err error) {
	return adclient.Login(_params)
}

func (rad remoteAdWrapper) GetUsers() ([]string, error) {
	return adclient.GetUsers()
}

func (rad remoteAdWrapper) CreateUser(cn string, container string, userShort string) (err error) {
	return adclient.CreateUser(cn, container, userShort)
}

func (rad remoteAdWrapper) GetUsersInGroup(group string, nested bool) (result []string, err error) {
	return adclient.GetUsersInGroup(group, nested)
}

func (rad remoteAdWrapper) GroupAddUser(group string, user string) (err error) {
	return adclient.GroupAddUser(group, user)
}

func (rad remoteAdWrapper) GetUserGroups(user string, nested bool) (result []string, err error) {
	return adclient.GetUserGroups(user, nested)
}

func (rad remoteAdWrapper) GroupRemoveUser(user string, group string) (err error)  {
	return adclient.GroupRemoveUser(user, group)
}

func (rad remoteAdWrapper) DeleteDN(dn string) (err error) {
	return adclient.DeleteDN(dn)
}

func (rad remoteAdWrapper) GetUserDisplayName(name string) (result string, err error) {
	return adclient.GetUserDisplayName(name)
}

func (rad remoteAdWrapper) GetDisabledUsers() ([]string, error) {
	return adclient.GetDisabledUsers()
}

func (rad remoteAdWrapper) SetUserPassword(user string, password string) (err error) {
	return adclient.SetUserPassword(user, password)
}

func (rad remoteAdWrapper) GetGroups() (groups []string, err error) {
	return adclient.GetGroups()
}

