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

package expandocolumn

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddColumn(tableId int64, name string, _type int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["tableId"] = tableId
	_params["name"] = name
	_params["type"] = _type

	_cmd := map[string]interface{}{
		"/expandocolumn/add-column": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddColumn2(tableId int64, name string, _type int, defaultData map[string]interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["tableId"] = tableId
	_params["name"] = name
	_params["type"] = _type
	_params["defaultData"] = defaultData

	_cmd := map[string]interface{}{
		"/expandocolumn/add-column": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteColumn(columnId int64) error {
	_params := make(map[string]interface{})

	_params["columnId"] = columnId

	_cmd := map[string]interface{}{
		"/expandocolumn/delete-column": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateColumn(columnId int64, name string, _type int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["columnId"] = columnId
	_params["name"] = name
	_params["type"] = _type

	_cmd := map[string]interface{}{
		"/expandocolumn/update-column": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateColumn2(columnId int64, name string, _type int, defaultData map[string]interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["columnId"] = columnId
	_params["name"] = name
	_params["type"] = _type
	_params["defaultData"] = defaultData

	_cmd := map[string]interface{}{
		"/expandocolumn/update-column": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateTypeSettings(columnId int64, typeSettings string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["columnId"] = columnId
	_params["typeSettings"] = typeSettings

	_cmd := map[string]interface{}{
		"/expandocolumn/update-type-settings": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

