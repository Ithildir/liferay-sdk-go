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

package layoutsetprototype

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLayoutSetPrototype(nameMap map[string]interface{}, description string, active bool, layoutsUpdateable bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active
	_params["layoutsUpdateable"] = layoutsUpdateable
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutsetprototype/add-layout-set-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteLayoutSetPrototype(layoutSetPrototypeId int64) error {
	_params := make(map[string]interface{})

	_params["layoutSetPrototypeId"] = layoutSetPrototypeId

	_cmd := map[string]interface{}{
		"/layoutsetprototype/delete-layout-set-prototype": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetLayoutSetPrototype(layoutSetPrototypeId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutSetPrototypeId"] = layoutSetPrototypeId

	_cmd := map[string]interface{}{
		"/layoutsetprototype/get-layout-set-prototype": _params,
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
		"/layoutsetprototype/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayoutSetPrototype(layoutSetPrototypeId int64, settings string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutSetPrototypeId"] = layoutSetPrototypeId
	_params["settings"] = settings

	_cmd := map[string]interface{}{
		"/layoutsetprototype/update-layout-set-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayoutSetPrototype2(layoutSetPrototypeId int64, nameMap map[string]interface{}, description string, active bool, layoutsUpdateable bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutSetPrototypeId"] = layoutSetPrototypeId
	_params["nameMap"] = nameMap
	_params["description"] = description
	_params["active"] = active
	_params["layoutsUpdateable"] = layoutsUpdateable
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutsetprototype/update-layout-set-prototype": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

