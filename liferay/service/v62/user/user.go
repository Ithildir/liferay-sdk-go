// Copyright (c) 2014 Andrea Di Giorgi. All rights reserved.
//
// This library is free software; you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation; either version 2.1 of the License, or (at your option)
// any later version.
//
// This library is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.

package user

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddGroupUsers(groupId int64, userIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userIds"] = userIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/add-group-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddOrganizationUsers(organizationId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/add-organization-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddPasswordPolicyUsers(passwordPolicyId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/add-password-policy-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddRoleUsers(roleId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/add-role-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddTeamUsers(teamId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/add-team-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddUser(companyId int64, autoPassword bool, password1 string, password2 string, autoScreenName bool, screenName string, emailAddress string, facebookId int64, openId string, locale string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupIds []interface{}, sendEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["autoPassword"] = autoPassword
	_params["password1"] = password1
	_params["password2"] = password2
	_params["autoScreenName"] = autoScreenName
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["locale"] = locale
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupIds"] = userGroupIds
	_params["sendEmail"] = sendEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/add-user": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddUser2(companyId int64, autoPassword bool, password1 string, password2 string, autoScreenName bool, screenName string, emailAddress string, facebookId int64, openId string, locale string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupIds []interface{}, addresses []interface{}, emailAddresses []interface{}, phones []interface{}, websites []interface{}, announcementsDelivers []interface{}, sendEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["autoPassword"] = autoPassword
	_params["password1"] = password1
	_params["password2"] = password2
	_params["autoScreenName"] = autoScreenName
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["locale"] = locale
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupIds"] = userGroupIds
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["phones"] = phones
	_params["websites"] = websites
	_params["announcementsDelivers"] = announcementsDelivers
	_params["sendEmail"] = sendEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/add-user": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddUserGroupUsers(userGroupId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/add-user-group-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddUserWithWorkflow(companyId int64, autoPassword bool, password1 string, password2 string, autoScreenName bool, screenName string, emailAddress string, facebookId int64, openId string, locale string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupIds []interface{}, sendEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["autoPassword"] = autoPassword
	_params["password1"] = password1
	_params["password2"] = password2
	_params["autoScreenName"] = autoScreenName
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["locale"] = locale
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupIds"] = userGroupIds
	_params["sendEmail"] = sendEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/add-user-with-workflow": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddUserWithWorkflow2(companyId int64, autoPassword bool, password1 string, password2 string, autoScreenName bool, screenName string, emailAddress string, facebookId int64, openId string, locale string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupIds []interface{}, addresses []interface{}, emailAddresses []interface{}, phones []interface{}, websites []interface{}, announcementsDelivers []interface{}, sendEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["autoPassword"] = autoPassword
	_params["password1"] = password1
	_params["password2"] = password2
	_params["autoScreenName"] = autoScreenName
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["locale"] = locale
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupIds"] = userGroupIds
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["phones"] = phones
	_params["websites"] = websites
	_params["announcementsDelivers"] = announcementsDelivers
	_params["sendEmail"] = sendEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/add-user-with-workflow": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeletePortrait(userId int64) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/delete-portrait": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteRoleUser(roleId int64, userId int64) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/delete-role-user": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteUser(userId int64) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/delete-user": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCompanyUsers(companyId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/user/get-company-users": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyUsersCount(companyId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/user/get-company-users-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupUserIds(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/user/get-group-user-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupUsers(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/user/get-group-users": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationUserIds(organizationId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/user/get-organization-user-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationUsers(organizationId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/user/get-organization-users": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRoleUserIds(roleId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId

	_cmd := map[string]interface{}{
		"/user/get-role-user-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserByEmailAddress(companyId int64, emailAddress string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["emailAddress"] = emailAddress

	_cmd := map[string]interface{}{
		"/user/get-user-by-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserById(userId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/get-user-by-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserByScreenName(companyId int64, screenName string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["screenName"] = screenName

	_cmd := map[string]interface{}{
		"/user/get-user-by-screen-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupUsers(userGroupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId

	_cmd := map[string]interface{}{
		"/user/get-user-group-users": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserIdByEmailAddress(companyId int64, emailAddress string) (int64, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["emailAddress"] = emailAddress

	_cmd := map[string]interface{}{
		"/user/get-user-id-by-email-address": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64)
	}

	return v, err
}

func (s *Service) GetUserIdByScreenName(companyId int64, screenName string) (int64, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["screenName"] = screenName

	_cmd := map[string]interface{}{
		"/user/get-user-id-by-screen-name": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64)
	}

	return v, err
}

func (s *Service) HasGroupUser(groupId int64, userId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/has-group-user": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) HasRoleUser(roleId int64, userId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/user/has-role-user": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) HasRoleUser2(companyId int64, name string, userId int64, inherited bool) (bool, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name
	_params["userId"] = userId
	_params["inherited"] = inherited

	_cmd := map[string]interface{}{
		"/user/has-role-user": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) SetRoleUsers(roleId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/set-role-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetUserGroupUsers(userGroupId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/set-user-group-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetGroupTeamsUsers(groupId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-group-teams-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetGroupUsers(groupId int64, userIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userIds"] = userIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/unset-group-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetOrganizationUsers(organizationId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-organization-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetPasswordPolicyUsers(passwordPolicyId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-password-policy-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetRoleUsers(roleId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-role-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetTeamUsers(teamId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-team-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetUserGroupUsers(userGroupId int64, userIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId
	_params["userIds"] = userIds

	_cmd := map[string]interface{}{
		"/user/unset-user-group-users": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateAgreedToTermsOfUse(userId int64, agreedToTermsOfUse bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["agreedToTermsOfUse"] = agreedToTermsOfUse

	_cmd := map[string]interface{}{
		"/user/update-agreed-to-terms-of-use": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateEmailAddress(userId int64, password string, emailAddress1 string, emailAddress2 string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["password"] = password
	_params["emailAddress1"] = emailAddress1
	_params["emailAddress2"] = emailAddress2
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/update-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateIncompleteUser(companyId int64, autoPassword bool, password1 string, password2 string, autoScreenName bool, screenName string, emailAddress string, facebookId int64, openId string, locale string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, jobTitle string, updateUserInformation bool, sendEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["autoPassword"] = autoPassword
	_params["password1"] = password1
	_params["password2"] = password2
	_params["autoScreenName"] = autoScreenName
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["locale"] = locale
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["jobTitle"] = jobTitle
	_params["updateUserInformation"] = updateUserInformation
	_params["sendEmail"] = sendEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/update-incomplete-user": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLockoutById(userId int64, lockout bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["lockout"] = lockout

	_cmd := map[string]interface{}{
		"/user/update-lockout-by-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOpenId(userId int64, openId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["openId"] = openId

	_cmd := map[string]interface{}{
		"/user/update-open-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrganizations(userId int64, organizationIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["organizationIds"] = organizationIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/update-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdatePassword(userId int64, password1 string, password2 string, passwordReset bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["password1"] = password1
	_params["password2"] = password2
	_params["passwordReset"] = passwordReset

	_cmd := map[string]interface{}{
		"/user/update-password": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdatePortrait(userId int64, bytes []byte) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["bytes"] = bytes

	_cmd := map[string]interface{}{
		"/user/update-portrait": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateReminderQuery(userId int64, question string, answer string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["question"] = question
	_params["answer"] = answer

	_cmd := map[string]interface{}{
		"/user/update-reminder-query": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateScreenName(userId int64, screenName string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["screenName"] = screenName

	_cmd := map[string]interface{}{
		"/user/update-screen-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateStatus(userId int64, status int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/user/update-status": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateUser(userId int64, oldPassword string, newPassword1 string, newPassword2 string, passwordReset bool, reminderQueryQuestion string, reminderQueryAnswer string, screenName string, emailAddress string, facebookId int64, openId string, languageId string, timeZoneId string, greeting string, comments string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, smsSn string, aimSn string, facebookSn string, icqSn string, jabberSn string, msnSn string, mySpaceSn string, skypeSn string, twitterSn string, ymSn string, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupRoles []interface{}, userGroupIds []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["oldPassword"] = oldPassword
	_params["newPassword1"] = newPassword1
	_params["newPassword2"] = newPassword2
	_params["passwordReset"] = passwordReset
	_params["reminderQueryQuestion"] = reminderQueryQuestion
	_params["reminderQueryAnswer"] = reminderQueryAnswer
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["languageId"] = languageId
	_params["timeZoneId"] = timeZoneId
	_params["greeting"] = greeting
	_params["comments"] = comments
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["smsSn"] = smsSn
	_params["aimSn"] = aimSn
	_params["facebookSn"] = facebookSn
	_params["icqSn"] = icqSn
	_params["jabberSn"] = jabberSn
	_params["msnSn"] = msnSn
	_params["mySpaceSn"] = mySpaceSn
	_params["skypeSn"] = skypeSn
	_params["twitterSn"] = twitterSn
	_params["ymSn"] = ymSn
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupRoles"] = userGroupRoles
	_params["userGroupIds"] = userGroupIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/update-user": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateUser2(userId int64, oldPassword string, newPassword1 string, newPassword2 string, passwordReset bool, reminderQueryQuestion string, reminderQueryAnswer string, screenName string, emailAddress string, facebookId int64, openId string, languageId string, timeZoneId string, greeting string, comments string, firstName string, middleName string, lastName string, prefixId int, suffixId int, male bool, birthdayMonth int, birthdayDay int, birthdayYear int, smsSn string, aimSn string, facebookSn string, icqSn string, jabberSn string, msnSn string, mySpaceSn string, skypeSn string, twitterSn string, ymSn string, jobTitle string, groupIds []interface{}, organizationIds []interface{}, roleIds []interface{}, userGroupRoles []interface{}, userGroupIds []interface{}, addresses []interface{}, emailAddresses []interface{}, phones []interface{}, websites []interface{}, announcementsDelivers []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["oldPassword"] = oldPassword
	_params["newPassword1"] = newPassword1
	_params["newPassword2"] = newPassword2
	_params["passwordReset"] = passwordReset
	_params["reminderQueryQuestion"] = reminderQueryQuestion
	_params["reminderQueryAnswer"] = reminderQueryAnswer
	_params["screenName"] = screenName
	_params["emailAddress"] = emailAddress
	_params["facebookId"] = facebookId
	_params["openId"] = openId
	_params["languageId"] = languageId
	_params["timeZoneId"] = timeZoneId
	_params["greeting"] = greeting
	_params["comments"] = comments
	_params["firstName"] = firstName
	_params["middleName"] = middleName
	_params["lastName"] = lastName
	_params["prefixId"] = prefixId
	_params["suffixId"] = suffixId
	_params["male"] = male
	_params["birthdayMonth"] = birthdayMonth
	_params["birthdayDay"] = birthdayDay
	_params["birthdayYear"] = birthdayYear
	_params["smsSn"] = smsSn
	_params["aimSn"] = aimSn
	_params["facebookSn"] = facebookSn
	_params["icqSn"] = icqSn
	_params["jabberSn"] = jabberSn
	_params["msnSn"] = msnSn
	_params["mySpaceSn"] = mySpaceSn
	_params["skypeSn"] = skypeSn
	_params["twitterSn"] = twitterSn
	_params["ymSn"] = ymSn
	_params["jobTitle"] = jobTitle
	_params["groupIds"] = groupIds
	_params["organizationIds"] = organizationIds
	_params["roleIds"] = roleIds
	_params["userGroupRoles"] = userGroupRoles
	_params["userGroupIds"] = userGroupIds
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["phones"] = phones
	_params["websites"] = websites
	_params["announcementsDelivers"] = announcementsDelivers
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/user/update-user": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

