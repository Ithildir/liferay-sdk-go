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

package expandovalue

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddValue(companyId int64, className string, tableName string, columnName string, classPK int64, data string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["className"] = className
	_params["tableName"] = tableName
	_params["columnName"] = columnName
	_params["classPK"] = classPK
	_params["data"] = data

	_cmd := map[string]interface{}{
		"/expandovalue/add-value": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddValues(companyId int64, className string, tableName string, classPK int64, attributeValues map[string]interface{}) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["className"] = className
	_params["tableName"] = tableName
	_params["classPK"] = classPK
	_params["attributeValues"] = attributeValues

	_cmd := map[string]interface{}{
		"/expandovalue/add-values": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetData(companyId int64, className string, tableName string, columnName string, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["className"] = className
	_params["tableName"] = tableName
	_params["columnName"] = columnName
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/expandovalue/get-data": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetData2(companyId int64, className string, tableName string, columnNames map[string]interface{}, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["className"] = className
	_params["tableName"] = tableName
	_params["columnNames"] = columnNames
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/expandovalue/get-data": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetJsonData(companyId int64, className string, tableName string, columnName string, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["className"] = className
	_params["tableName"] = tableName
	_params["columnName"] = columnName
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/expandovalue/get-json-data": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

