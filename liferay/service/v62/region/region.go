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

package region

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRegion(countryId int64, regionCode string, name string, active bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId
	_params["regionCode"] = regionCode
	_params["name"] = name
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/region/add-region": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) FetchRegion(countryId int64, regionCode string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId
	_params["regionCode"] = regionCode

	_cmd := map[string]interface{}{
		"/region/fetch-region": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRegion(regionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["regionId"] = regionId

	_cmd := map[string]interface{}{
		"/region/get-region": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRegion2(countryId int64, regionCode string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId
	_params["regionCode"] = regionCode

	_cmd := map[string]interface{}{
		"/region/get-region": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRegions() ([]interface{}, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/region/get-regions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRegions2(active bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["active"] = active

	_cmd := map[string]interface{}{
		"/region/get-regions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRegions3(countryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId

	_cmd := map[string]interface{}{
		"/region/get-regions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRegions4(countryId int64, active bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/region/get-regions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

