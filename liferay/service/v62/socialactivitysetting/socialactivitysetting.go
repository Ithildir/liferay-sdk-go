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

package socialactivitysetting

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetActivityDefinition(groupId int64, className string, activityType int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["activityType"] = activityType

	_cmd := map[string]interface{}{
		"/socialactivitysetting/get-activity-definition": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetActivityDefinitions(groupId int64, className string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className

	_cmd := map[string]interface{}{
		"/socialactivitysetting/get-activity-definitions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetActivitySettings(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/socialactivitysetting/get-activity-settings": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetJsonActivityDefinitions(groupId int64, className string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className

	_cmd := map[string]interface{}{
		"/socialactivitysetting/get-json-activity-definitions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateActivitySetting(groupId int64, className string, enabled bool) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["enabled"] = enabled

	_cmd := map[string]interface{}{
		"/socialactivitysetting/update-activity-setting": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateActivitySetting2(groupId int64, className string, activityType int, activityCounterDefinition *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["activityType"] = activityType
	liferay.MangleObjectWrapper(_params, "activityCounterDefinition", "com.liferay.portlet.social.model.SocialActivityCounterDefinition", activityCounterDefinition)

	_cmd := map[string]interface{}{
		"/socialactivitysetting/update-activity-setting": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateActivitySettings(groupId int64, className string, activityType int, activityCounterDefinitions []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["activityType"] = activityType
	_params["activityCounterDefinitions"] = activityCounterDefinitions

	_cmd := map[string]interface{}{
		"/socialactivitysetting/update-activity-settings": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

