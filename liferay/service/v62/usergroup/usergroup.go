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

package usergroup

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddGroupUserGroups(groupId int64, userGroupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userGroupIds"] = userGroupIds

	_cmd := map[string]interface{}{
		"/usergroup/add-group-user-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddTeamUserGroups(teamId int64, userGroupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId
	_params["userGroupIds"] = userGroupIds

	_cmd := map[string]interface{}{
		"/usergroup/add-team-user-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddUserGroup(name string, description string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/usergroup/add-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddUserGroup2(name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/usergroup/add-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteUserGroup(userGroupId int64) error {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId

	_cmd := map[string]interface{}{
		"/usergroup/delete-user-group": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetUserGroup(name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name

	_cmd := map[string]interface{}{
		"/usergroup/get-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroup2(userGroupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId

	_cmd := map[string]interface{}{
		"/usergroup/get-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserUserGroups(userId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/usergroup/get-user-user-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UnsetGroupUserGroups(groupId int64, userGroupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userGroupIds"] = userGroupIds

	_cmd := map[string]interface{}{
		"/usergroup/unset-group-user-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetTeamUserGroups(teamId int64, userGroupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId
	_params["userGroupIds"] = userGroupIds

	_cmd := map[string]interface{}{
		"/usergroup/unset-team-user-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateUserGroup(userGroupId int64, name string, description string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId
	_params["name"] = name
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/usergroup/update-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateUserGroup2(userGroupId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userGroupId"] = userGroupId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/usergroup/update-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

