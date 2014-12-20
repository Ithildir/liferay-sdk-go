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

package resourceblock

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCompanyScopePermission(scopeGroupId int64, companyId int64, name string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/add-company-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddGroupScopePermission(scopeGroupId int64, companyId int64, groupId int64, name string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/add-group-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddIndividualScopePermission(companyId int64, groupId int64, name string, primKey int64, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/add-individual-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveAllGroupScopePermissions(scopeGroupId int64, companyId int64, name string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/remove-all-group-scope-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveCompanyScopePermission(scopeGroupId int64, companyId int64, name string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/remove-company-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveGroupScopePermission(scopeGroupId int64, companyId int64, groupId int64, name string, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/remove-group-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveIndividualScopePermission(companyId int64, groupId int64, name string, primKey int64, roleId int64, actionId string) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionId"] = actionId

	_cmd := map[string]interface{}{
		"/resourceblock/remove-individual-scope-permission": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetCompanyScopePermissions(scopeGroupId int64, companyId int64, name string, roleId int64, actionIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionIds"] = actionIds

	_cmd := map[string]interface{}{
		"/resourceblock/set-company-scope-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetGroupScopePermissions(scopeGroupId int64, companyId int64, groupId int64, name string, roleId int64, actionIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["scopeGroupId"] = scopeGroupId
	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["roleId"] = roleId
	_params["actionIds"] = actionIds

	_cmd := map[string]interface{}{
		"/resourceblock/set-group-scope-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetIndividualScopePermissions(companyId int64, groupId int64, name string, primKey int64, roleIdsToActionIds map[string]interface{}) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleIdsToActionIds"] = roleIdsToActionIds

	_cmd := map[string]interface{}{
		"/resourceblock/set-individual-scope-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetIndividualScopePermissions2(companyId int64, groupId int64, name string, primKey int64, roleId int64, actionIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["primKey"] = primKey
	_params["roleId"] = roleId
	_params["actionIds"] = actionIds

	_cmd := map[string]interface{}{
		"/resourceblock/set-individual-scope-permissions": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

