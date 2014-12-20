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

package resourcepermission

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddResourcePermission(groupId int64, companyId int64, name string, scope int, primKey string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["scope"] = scope
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourcepermission/add-resource-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveResourcePermission(groupId int64, companyId int64, name string, scope int, primKey string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["scope"] = scope
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourcepermission/remove-resource-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveResourcePermissions(groupId int64, companyId int64, name string, scope int, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["scope"] = scope
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourcepermission/remove-resource-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetIndividualResourcePermissions(groupId int64, companyId int64, name string, primKey string, roleIdsToActionIds map[string]interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleIdsToActionIds"] = roleIdsToActionIds

	_cmd := map[string]interface{}{
		"/resourcepermission/set-individual-resource-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetIndividualResourcePermissions2(groupId int64, companyId int64, name string, primKey string, roleId int64, actionIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionIds"] = actionIds

	_cmd := map[string]interface{}{
		"/resourcepermission/set-individual-resource-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

