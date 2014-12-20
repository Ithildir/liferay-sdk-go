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

package mdrrulegroup

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRuleGroup(groupId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrulegroup/add-rule-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyRuleGroup(ruleGroupId int64, groupId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId
	_params["groupId"] = groupId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrulegroup/copy-rule-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteRuleGroup(ruleGroupId int64) error {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId

	_cmd := map[string]interface{}{
		"/mdrrulegroup/delete-rule-group": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) FetchRuleGroup(ruleGroupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId

	_cmd := map[string]interface{}{
		"/mdrrulegroup/fetch-rule-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRuleGroup(ruleGroupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId

	_cmd := map[string]interface{}{
		"/mdrrulegroup/get-rule-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRuleGroup(ruleGroupId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupId"] = ruleGroupId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrulegroup/update-rule-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

