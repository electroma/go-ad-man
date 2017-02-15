package logic

import (
	"github.com/stretchr/testify/mock"
	"github.com/paleg/libadclient"
)

type mockAD struct {
	mock.Mock
}

func (m mockAD) DisableUser(user string) (err error) {
	return m.Called(user).Error(0)
}

func (m mockAD) EnableUser(user string) (err error) {
	return m.Called(user).Error(0)
}

func (m mockAD) SetUserDisplayName(user string, displayName string) (err error) {
	return m.Called(user, displayName).Error(0)
}

func (m mockAD) SearchBase() (result string) {
	return m.Called().String(0)
}

func (m mockAD) Login(_params adclient.ADConnParams) (err error) {
	return m.Called(_params).Error(0)
}

func (m mockAD) GetUsers() ([]string, error) {
	res := m.Called()
	return res.Get(0).([]string), res.Error(1)
}

func (m mockAD) CreateUser(cn string, container string, user_short string) (err error) {
	return m.Called(cn, container, user_short).Error(0)
}

func (m mockAD) DeleteDN(dn string) (err error) {
	return m.Called(dn).Error(0)
}

func (m mockAD) GetUserDisplayName(name string) (result string, err error) {
	res := m.Called(name)
	return res.String(0), res.Error(1)
}

func (m mockAD) GetDisabledUsers() ([]string, error) {
	res := m.Called()
	return res.Get(0).([]string), res.Error(1)
}

func (m mockAD) SetUserPassword(user string, password string) (err error) {
	return m.Called(user, password).Error(0)
}
