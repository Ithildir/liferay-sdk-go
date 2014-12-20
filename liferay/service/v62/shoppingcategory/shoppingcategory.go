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

package shoppingcategory

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCategory(parentCategoryId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingcategory/add-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteCategory(categoryId int64) error {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/shoppingcategory/delete-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategories(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/shoppingcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories2(groupId int64, parentCategoryId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryId"] = parentCategoryId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/shoppingcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategoriesCount(groupId int64, parentCategoryId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryId"] = parentCategoryId

	_cmd := map[string]interface{}{
		"/shoppingcategory/get-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetCategory(categoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/shoppingcategory/get-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetSubcategoryIds(categoryIds []interface{}, groupId int64, categoryId int64) error {
	_params := make(map[string]interface{})

	_params["categoryIds"] = categoryIds
	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/shoppingcategory/get-subcategory-ids": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateCategory(categoryId int64, parentCategoryId int64, name string, description string, mergeWithParentCategory bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["parentCategoryId"] = parentCategoryId
	_params["name"] = name
	_params["description"] = description
	_params["mergeWithParentCategory"] = mergeWithParentCategory
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingcategory/update-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

