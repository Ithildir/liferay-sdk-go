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

package layoutprototype

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLayoutPrototype(nameMap map[string]interface{}, description string, active bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/layoutprototype/add-layout-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddLayoutPrototype2(nameMap map[string]interface{}, description string, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutprototype/add-layout-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteLayoutPrototype(layoutPrototypeId int64) error {
	_params := make(map[string]interface{})

	_params["layoutPrototypeId"] = layoutPrototypeId

	_cmd := map[string]interface{}{
		"/layoutprototype/delete-layout-prototype": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetLayoutPrototype(layoutPrototypeId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutPrototypeId"] = layoutPrototypeId

	_cmd := map[string]interface{}{
		"/layoutprototype/get-layout-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, active *liferay.ObjectWrapper, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	liferay.MangleObjectWrapper(_params, "active", "java.lang.Boolean", active)
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/layoutprototype/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayoutPrototype(layoutPrototypeId int64, nameMap map[string]interface{}, description string, active bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutPrototypeId"] = layoutPrototypeId
	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/layoutprototype/update-layout-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayoutPrototype2(layoutPrototypeId int64, nameMap map[string]interface{}, description string, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutPrototypeId"] = layoutPrototypeId
	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutprototype/update-layout-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

