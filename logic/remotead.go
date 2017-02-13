package logic

import (
	"github.com/paleg/libadclient"
	"github.com/Sirupsen/logrus"
	"github.com/kataras/go-errors"
)

type remoteAdWrapper struct {
}

var Base, DCHost string

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

func (_ remoteAdWrapper) DisableUser(user string) (err error) {
	return adclient.DisableUser(user)
}

func (_ remoteAdWrapper) EnableUser(user string) (err error) {
	return adclient.EnableUser(user)
}


func (_ remoteAdWrapper) SetUserDisplayName(user string, displayName string) (err error) {
	return adclient.SetUserDisplayName(user, displayName)
}

func (_ remoteAdWrapper) SearchBase() (result string) {
	return adclient.SearchBase()
}

func (_ remoteAdWrapper) Login(_params adclient.ADConnParams) (err error) {
	return adclient.Login(_params)
}

func (_ remoteAdWrapper) GetUsers() ([]string, error) {
	return adclient.GetUsers()
}

func (_ remoteAdWrapper) CreateUser(cn string, container string, user_short string) (err error) {
	return adclient.CreateUser(cn, container, user_short)
}

func (_ remoteAdWrapper) DeleteDN(dn string) (err error) {
	return adclient.DeleteDN(dn)
}

func (_ remoteAdWrapper) GetUserDisplayName(name string) (result string, err error) {
	return adclient.GetUserDisplayName(name)
}

func (_ remoteAdWrapper) GetDisabledUsers() ([]string, error) {
	return adclient.GetDisabledUsers()
}

