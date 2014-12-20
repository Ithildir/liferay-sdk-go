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

package classname

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) FetchClassName(value string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["value"] = value

	_cmd := map[string]interface{}{
		"/classname/fetch-class-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) FetchClassNameId(clazz map[string]interface{}) (int64, error) {
	_params := make(map[string]interface{})

	_params["clazz"] = clazz

	_cmd := map[string]interface{}{
		"/classname/fetch-class-name-id": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64)
	}

	return v, err
}

func (s *Service) FetchClassNameId2(value string) (int64, error) {
	_params := make(map[string]interface{})

	_params["value"] = value

	_cmd := map[string]interface{}{
		"/classname/fetch-class-name-id": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64)
	}

	return v, err
}

