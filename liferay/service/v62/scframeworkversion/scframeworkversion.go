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

package scframeworkversion

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFrameworkVersion(name string, url string, active bool, priority int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["url"] = url
	_params["active"] = active
	_params["priority"] = priority
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/scframeworkversion/add-framework-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFrameworkVersion(frameworkVersionId int64) error {
	_params := make(map[string]interface{})

	_params["frameworkVersionId"] = frameworkVersionId

	_cmd := map[string]interface{}{
		"/scframeworkversion/delete-framework-version": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFrameworkVersion(frameworkVersionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["frameworkVersionId"] = frameworkVersionId

	_cmd := map[string]interface{}{
		"/scframeworkversion/get-framework-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFrameworkVersions(groupId int64, active bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/scframeworkversion/get-framework-versions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFrameworkVersions2(groupId int64, active bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["active"] = active
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/scframeworkversion/get-framework-versions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateFrameworkVersion(frameworkVersionId int64, name string, url string, active bool, priority int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["frameworkVersionId"] = frameworkVersionId
	_params["name"] = name
	_params["url"] = url
	_params["active"] = active
	_params["priority"] = priority

	_cmd := map[string]interface{}{
		"/scframeworkversion/update-framework-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

