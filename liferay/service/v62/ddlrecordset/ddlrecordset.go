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

package ddlrecordset

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRecordSet(groupId int64, ddmStructureId int64, recordSetKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, minDisplayRows int, scope int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["ddmStructureId"] = ddmStructureId
	_params["recordSetKey"] = recordSetKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["minDisplayRows"] = minDisplayRows
	_params["scope"] = scope
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecordset/add-record-set": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteRecordSet(recordSetId int64) error {
	_params := make(map[string]interface{})

	_params["recordSetId"] = recordSetId

	_cmd := map[string]interface{}{
		"/ddlrecordset/delete-record-set": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetRecordSet(recordSetId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordSetId"] = recordSetId

	_cmd := map[string]interface{}{
		"/ddlrecordset/get-record-set": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupId int64, keywords string, scope int, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["keywords"] = keywords
	_params["scope"] = scope
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddlrecordset/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupId int64, name string, description string, scope int, andOperator bool, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["description"] = description
	_params["scope"] = scope
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddlrecordset/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupId int64, keywords string, scope int) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["keywords"] = keywords
	_params["scope"] = scope

	_cmd := map[string]interface{}{
		"/ddlrecordset/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupId int64, name string, description string, scope int, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["description"] = description
	_params["scope"] = scope
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/ddlrecordset/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) UpdateMinDisplayRows(recordSetId int64, minDisplayRows int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordSetId"] = recordSetId
	_params["minDisplayRows"] = minDisplayRows
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecordset/update-min-display-rows": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRecordSet(recordSetId int64, ddmStructureId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, minDisplayRows int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["recordSetId"] = recordSetId
	_params["ddmStructureId"] = ddmStructureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["minDisplayRows"] = minDisplayRows
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecordset/update-record-set": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRecordSet2(groupId int64, ddmStructureId int64, recordSetKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, minDisplayRows int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["ddmStructureId"] = ddmStructureId
	_params["recordSetKey"] = recordSetKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["minDisplayRows"] = minDisplayRows
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddlrecordset/update-record-set": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

