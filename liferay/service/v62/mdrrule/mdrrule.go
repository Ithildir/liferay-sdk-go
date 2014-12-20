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

package mdrrule

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRule(ruleGroupId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, _type string, typeSettings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = _type
	_params["typeSettings"] = typeSettings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrule/add-rule": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteRule(ruleId int64) error {
	_params := make(map[string]interface{})

	_params["ruleId"] = ruleId

	_cmd := map[string]interface{}{
		"/mdrrule/delete-rule": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) FetchRule(ruleId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleId"] = ruleId

	_cmd := map[string]interface{}{
		"/mdrrule/fetch-rule": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRule(ruleId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleId"] = ruleId

	_cmd := map[string]interface{}{
		"/mdrrule/get-rule": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRule(ruleId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, _type string, typeSettings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleId"] = ruleId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = _type
	_params["typeSettings"] = typeSettings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrule/update-rule": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRule2(ruleId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, _type string, typeSettingsProperties *liferay.ObjectWrapper, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleId"] = ruleId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = _type
	liferay.MangleObjectWrapper(_params, "typeSettingsProperties", "com.liferay.portal.kernel.util.UnicodeProperties", typeSettingsProperties)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrule/update-rule": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

