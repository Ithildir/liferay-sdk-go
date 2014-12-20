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

package role

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRole(name string, titleMap map[string]interface{}, descriptionMap map[string]interface{}, _type int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = _type

	_cmd := map[string]interface{}{
		"/role/add-role": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddRole2(className string, classPK int64, name string, titleMap map[string]interface{}, descriptionMap map[string]interface{}, _type int, subtype string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["name"] = name
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = _type
	_params["subtype"] = subtype
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/role/add-role": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddUserRoles(userId int64, roleIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["roleIds"] = roleIds

	_cmd := map[string]interface{}{
		"/role/add-user-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteRole(roleId int64) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId

	_cmd := map[string]interface{}{
		"/role/delete-role": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetGroupRoles(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/role/get-group-roles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRole(roleId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId

	_cmd := map[string]interface{}{
		"/role/get-role": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRole2(companyId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/role/get-role": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupGroupRoles(userId int64, groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/role/get-user-group-group-roles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupRoles(userId int64, groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/role/get-user-group-roles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserRelatedRoles(userId int64, groups []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groups"] = groups

	_cmd := map[string]interface{}{
		"/role/get-user-related-roles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserRoles(userId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/role/get-user-roles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) HasUserRole(userId int64, companyId int64, name string, inherited bool) (bool, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["companyId"] = companyId
	_params["name"] = name
	_params["inherited"] = inherited

	_cmd := map[string]interface{}{
		"/role/has-user-role": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) HasUserRoles(userId int64, companyId int64, names []interface{}, inherited bool) (bool, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["companyId"] = companyId
	_params["names"] = names
	_params["inherited"] = inherited

	_cmd := map[string]interface{}{
		"/role/has-user-roles": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) UnsetUserRoles(userId int64, roleIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["roleIds"] = roleIds

	_cmd := map[string]interface{}{
		"/role/unset-user-roles": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateRole(roleId int64, name string, titleMap map[string]interface{}, descriptionMap map[string]interface{}, subtype string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["name"] = name
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["subtype"] = subtype
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/role/update-role": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

