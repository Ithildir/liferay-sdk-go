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

package mdrrulegroupinstance

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRuleGroupInstance(groupId int64, className string, classPK int64, ruleGroupId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["ruleGroupId"] = ruleGroupId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/add-rule-group-instance": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddRuleGroupInstance2(groupId int64, className string, classPK int64, ruleGroupId int64, priority int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["ruleGroupId"] = ruleGroupId
	_params["priority"] = priority
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/add-rule-group-instance": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteRuleGroupInstance(ruleGroupInstanceId int64) error {
	_params := make(map[string]interface{})

	_params["ruleGroupInstanceId"] = ruleGroupInstanceId

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/delete-rule-group-instance": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetRuleGroupInstances(className string, classPK int64, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/get-rule-group-instances": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRuleGroupInstancesCount(className string, classPK int64) (int, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/get-rule-group-instances-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) UpdateRuleGroupInstance(ruleGroupInstanceId int64, priority int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["ruleGroupInstanceId"] = ruleGroupInstanceId
	_params["priority"] = priority

	_cmd := map[string]interface{}{
		"/mdrrulegroupinstance/update-rule-group-instance": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

