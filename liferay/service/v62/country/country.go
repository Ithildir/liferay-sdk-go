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

package country

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCountry(name string, a2 string, a3 string, number string, idd string, active bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["a2"] = a2
	_params["a3"] = a3
	_params["number"] = number
	_params["idd"] = idd
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/country/add-country": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) FetchCountry(countryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId

	_cmd := map[string]interface{}{
		"/country/fetch-country": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) FetchCountryByA2(a2 string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["a2"] = a2

	_cmd := map[string]interface{}{
		"/country/fetch-country-by-a2": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) FetchCountryByA3(a3 string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["a3"] = a3

	_cmd := map[string]interface{}{
		"/country/fetch-country-by-a3": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCountries() ([]interface{}, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/country/get-countries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCountries2(active bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["active"] = active

	_cmd := map[string]interface{}{
		"/country/get-countries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCountry(countryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["countryId"] = countryId

	_cmd := map[string]interface{}{
		"/country/get-country": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCountryByA2(a2 string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["a2"] = a2

	_cmd := map[string]interface{}{
		"/country/get-country-by-a2": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCountryByA3(a3 string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["a3"] = a3

	_cmd := map[string]interface{}{
		"/country/get-country-by-a3": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCountryByName(name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name

	_cmd := map[string]interface{}{
		"/country/get-country-by-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

