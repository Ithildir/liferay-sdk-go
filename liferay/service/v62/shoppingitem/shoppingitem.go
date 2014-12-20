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

package shoppingitem

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddBookItems(groupId int64, categoryId int64, isbns []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["isbns"] = isbns

	_cmd := map[string]interface{}{
		"/shoppingitem/add-book-items": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddItem(groupId int64, categoryId int64, sku string, name string, description string, properties string, fieldsQuantities string, requiresShipping bool, stockQuantity int, featured bool, sale *liferay.ObjectWrapper, smallImage bool, smallImageURL string, smallFile io.Reader, mediumImage bool, mediumImageURL string, mediumFile io.Reader, largeImage bool, largeImageURL string, largeFile io.Reader, itemFields []interface{}, itemPrices []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["sku"] = sku
	_params["name"] = name
	_params["description"] = description
	_params["properties"] = properties
	_params["fieldsQuantities"] = fieldsQuantities
	_params["requiresShipping"] = requiresShipping
	_params["stockQuantity"] = stockQuantity
	_params["featured"] = featured
	liferay.MangleObjectWrapper(_params, "sale", "java.lang.Boolean", sale)
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	_params["mediumImage"] = mediumImage
	_params["mediumImageURL"] = mediumImageURL
	_params["mediumFile"] = mediumFile
	_params["largeImage"] = largeImage
	_params["largeImageURL"] = largeImageURL
	_params["largeFile"] = largeFile
	_params["itemFields"] = itemFields
	_params["itemPrices"] = itemPrices
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingitem/add-item": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteItem(itemId int64) error {
	_params := make(map[string]interface{})

	_params["itemId"] = itemId

	_cmd := map[string]interface{}{
		"/shoppingitem/delete-item": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategoriesItemsCount(groupId int64, categoryIds []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryIds"] = categoryIds

	_cmd := map[string]interface{}{
		"/shoppingitem/get-categories-items-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetItem(itemId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["itemId"] = itemId

	_cmd := map[string]interface{}{
		"/shoppingitem/get-item": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetItems(groupId int64, categoryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/shoppingitem/get-items": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetItems2(groupId int64, categoryId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/shoppingitem/get-items": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetItemsCount(groupId int64, categoryId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/shoppingitem/get-items-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetItemsPrevAndNext(itemId int64, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["itemId"] = itemId
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/shoppingitem/get-items-prev-and-next": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateItem(itemId int64, groupId int64, categoryId int64, sku string, name string, description string, properties string, fieldsQuantities string, requiresShipping bool, stockQuantity int, featured bool, sale *liferay.ObjectWrapper, smallImage bool, smallImageURL string, smallFile io.Reader, mediumImage bool, mediumImageURL string, mediumFile io.Reader, largeImage bool, largeImageURL string, largeFile io.Reader, itemFields []interface{}, itemPrices []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["itemId"] = itemId
	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["sku"] = sku
	_params["name"] = name
	_params["description"] = description
	_params["properties"] = properties
	_params["fieldsQuantities"] = fieldsQuantities
	_params["requiresShipping"] = requiresShipping
	_params["stockQuantity"] = stockQuantity
	_params["featured"] = featured
	liferay.MangleObjectWrapper(_params, "sale", "java.lang.Boolean", sale)
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	_params["mediumImage"] = mediumImage
	_params["mediumImageURL"] = mediumImageURL
	_params["mediumFile"] = mediumFile
	_params["largeImage"] = largeImage
	_params["largeImageURL"] = largeImageURL
	_params["largeFile"] = largeFile
	_params["itemFields"] = itemFields
	_params["itemPrices"] = itemPrices
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingitem/update-item": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

