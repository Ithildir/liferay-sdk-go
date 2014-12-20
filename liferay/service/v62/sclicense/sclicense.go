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

package sclicense

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLicense(name string, url string, openSource bool, active bool, recommended bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["url"] = url
	_params["openSource"] = openSource
	_params["active"] = active
	_params["recommended"] = recommended

	_cmd := map[string]interface{}{
		"/sclicense/add-license": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteLicense(licenseId int64) error {
	_params := make(map[string]interface{})

	_params["licenseId"] = licenseId

	_cmd := map[string]interface{}{
		"/sclicense/delete-license": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetLicense(licenseId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["licenseId"] = licenseId

	_cmd := map[string]interface{}{
		"/sclicense/get-license": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLicense(licenseId int64, name string, url string, openSource bool, active bool, recommended bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["licenseId"] = licenseId
	_params["name"] = name
	_params["url"] = url
	_params["openSource"] = openSource
	_params["active"] = active
	_params["recommended"] = recommended

	_cmd := map[string]interface{}{
		"/sclicense/update-license": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

