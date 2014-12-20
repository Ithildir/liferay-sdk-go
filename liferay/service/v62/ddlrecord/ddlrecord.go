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

package ddlrecord

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRecord(groupId int64, recordSetId int64, displayIndex int, fields *liferay.ObjectWrapper, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["recordSetId"] = recordSetId
	_params["displayIndex"] = displayIndex
	liferay.MangleObjectWrapper(_params, "fields", "com.liferay.portlet.dynamicdatamapping.storage.Fields", fields)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecord/add-record": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddRecord2(groupId int64, recordSetId int64, displayIndex int, fieldsMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["recordSetId"] = recordSetId
	_params["displayIndex"] = displayIndex
	_params["fieldsMap"] = fieldsMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecord/add-record": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteRecordLocale(recordId int64, locale string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordId"] = recordId
	_params["locale"] = locale
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecord/delete-record-locale": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRecord(recordId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordId"] = recordId

	_cmd := map[string]interface{}{
		"/ddlrecord/get-record": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRecord(recordId int64, displayIndex int, fieldsMap map[string]interface{}, mergeFields bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordId"] = recordId
	_params["displayIndex"] = displayIndex
	_params["fieldsMap"] = fieldsMap
	_params["mergeFields"] = mergeFields
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecord/update-record": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRecord2(recordId int64, majorVersion bool, displayIndex int, fields *liferay.ObjectWrapper, mergeFields bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordId"] = recordId
	_params["majorVersion"] = majorVersion
	_params["displayIndex"] = displayIndex
	liferay.MangleObjectWrapper(_params, "fields", "com.liferay.portlet.dynamicdatamapping.storage.Fields", fields)
	_params["mergeFields"] = mergeFields
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecord/update-record": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

