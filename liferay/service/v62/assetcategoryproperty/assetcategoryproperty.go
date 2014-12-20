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

package assetcategoryproperty

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCategoryProperty(entryId int64, key string, value string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["key"] = key
	_params["value"] = value

	_cmd := map[string]interface{}{
		"/assetcategoryproperty/add-category-property": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteCategoryProperty(categoryPropertyId int64) error {
	_params := make(map[string]interface{})

	_params["categoryPropertyId"] = categoryPropertyId

	_cmd := map[string]interface{}{
		"/assetcategoryproperty/delete-category-property": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategoryProperties(entryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/assetcategoryproperty/get-category-properties": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategoryPropertyValues(companyId int64, key string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["key"] = key

	_cmd := map[string]interface{}{
		"/assetcategoryproperty/get-category-property-values": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateCategoryProperty(categoryPropertyId int64, key string, value string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryPropertyId"] = categoryPropertyId
	_params["key"] = key
	_params["value"] = value

	_cmd := map[string]interface{}{
		"/assetcategoryproperty/update-category-property": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

