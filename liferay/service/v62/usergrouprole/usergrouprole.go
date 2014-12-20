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

package usergrouprole

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddUserGroupRoles(userId int64, groupId int64, roleIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["roleIds"] = roleIds

	_cmd := map[string]interface{}{
		"/usergrouprole/add-user-group-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddUserGroupRoles2(userIds []interface{}, groupId int64, roleId int64) error {
	_params := make(map[string]interface{})

	_params["userIds"] = userIds
	_params["groupId"] = groupId
	_params["roleId"] = roleId

	_cmd := map[string]interface{}{
		"/usergrouprole/add-user-group-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteUserGroupRoles(userId int64, groupId int64, roleIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["roleIds"] = roleIds

	_cmd := map[string]interface{}{
		"/usergrouprole/delete-user-group-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteUserGroupRoles2(userIds []interface{}, groupId int64, roleId int64) error {
	_params := make(map[string]interface{})

	_params["userIds"] = userIds
	_params["groupId"] = groupId
	_params["roleId"] = roleId

	_cmd := map[string]interface{}{
		"/usergrouprole/delete-user-group-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

