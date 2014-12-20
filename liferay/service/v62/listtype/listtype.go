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

package listtype

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetListType(listTypeId int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["listTypeId"] = listTypeId

	_cmd := map[string]interface{}{
		"/listtype/get-list-type": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetListTypes(type string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["type"] = type

	_cmd := map[string]interface{}{
		"/listtype/get-list-types": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Validate(listTypeId int, type string) error {
	_params := make(map[string]interface{})

	_params["listTypeId"] = listTypeId
	_params["type"] = type

	_cmd := map[string]interface{}{
		"/listtype/validate": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Validate2(listTypeId int, classNameId int64, type string) error {
	_params := make(map[string]interface{})

	_params["listTypeId"] = listTypeId
	_params["classNameId"] = classNameId
	_params["type"] = type

	_cmd := map[string]interface{}{
		"/listtype/validate": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

