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

package mbcategory

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCategory(parentCategoryId int64, name string, description string, displayStyle string, emailAddress string, inProtocol string, inServerName string, inServerPort int, inUseSSL bool, inUserName string, inPassword string, inReadInterval int, outEmailAddress string, outCustom bool, outServerName string, outServerPort int, outUseSSL bool, outUserName string, outPassword string, mailingListActive bool, allowAnonymousEmail bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId
	_params["name"] = name
	_params["description"] = description
	_params["displayStyle"] = displayStyle
	_params["emailAddress"] = emailAddress
	_params["inProtocol"] = inProtocol
	_params["inServerName"] = inServerName
	_params["inServerPort"] = inServerPort
	_params["inUseSSL"] = inUseSSL
	_params["inUserName"] = inUserName
	_params["inPassword"] = inPassword
	_params["inReadInterval"] = inReadInterval
	_params["outEmailAddress"] = outEmailAddress
	_params["outCustom"] = outCustom
	_params["outServerName"] = outServerName
	_params["outServerPort"] = outServerPort
	_params["outUseSSL"] = outUseSSL
	_params["outUserName"] = outUserName
	_params["outPassword"] = outPassword
	_params["mailingListActive"] = mailingListActive
	_params["allowAnonymousEmail"] = allowAnonymousEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbcategory/add-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddCategory2(userId int64, parentCategoryId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["parentCategoryId"] = parentCategoryId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbcategory/add-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteCategory(categoryId int64, includeTrashedEntries bool) error {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["includeTrashedEntries"] = includeTrashedEntries

	_cmd := map[string]interface{}{
		"/mbcategory/delete-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteCategory2(groupId int64, categoryId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/delete-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategories(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories2(groupId int64, status int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories3(groupId int64, parentCategoryId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryId"] = parentCategoryId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories4(groupId int64, parentCategoryIds []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryIds"] = parentCategoryIds
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories5(groupId int64, parentCategoryId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryId"] = parentCategoryId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategories6(groupId int64, parentCategoryIds []interface{}, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryIds"] = parentCategoryIds
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories": _params,
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
		"/mbcategory/get-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetCategoriesCount2(groupId int64, parentCategoryIds []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryIds"] = parentCategoryIds

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetCategoriesCount3(groupId int64, parentCategoryId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryId"] = parentCategoryId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetCategoriesCount4(groupId int64, parentCategoryIds []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentCategoryIds"] = parentCategoryIds
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbcategory/get-categories-count": _params,
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
		"/mbcategory/get-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCategoryIds(groupId int64, categoryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/get-category-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSubcategoryIds(categoryIds []interface{}, groupId int64, categoryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryIds"] = categoryIds
	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/get-subcategory-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSubscribedCategories(groupId int64, userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbcategory/get-subscribed-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSubscribedCategoriesCount(groupId int64, userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/mbcategory/get-subscribed-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) MoveCategory(categoryId int64, parentCategoryId int64, mergeWithParentCategory bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["parentCategoryId"] = parentCategoryId
	_params["mergeWithParentCategory"] = mergeWithParentCategory

	_cmd := map[string]interface{}{
		"/mbcategory/move-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveCategoryFromTrash(categoryId int64, newCategoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["newCategoryId"] = newCategoryId

	_cmd := map[string]interface{}{
		"/mbcategory/move-category-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveCategoryToTrash(categoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/move-category-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreCategoryFromTrash(categoryId int64) error {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/restore-category-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SubscribeCategory(groupId int64, categoryId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/subscribe-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeCategory(groupId int64, categoryId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/mbcategory/unsubscribe-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateCategory(categoryId int64, parentCategoryId int64, name string, description string, displayStyle string, emailAddress string, inProtocol string, inServerName string, inServerPort int, inUseSSL bool, inUserName string, inPassword string, inReadInterval int, outEmailAddress string, outCustom bool, outServerName string, outServerPort int, outUseSSL bool, outUserName string, outPassword string, mailingListActive bool, allowAnonymousEmail bool, mergeWithParentCategory bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["parentCategoryId"] = parentCategoryId
	_params["name"] = name
	_params["description"] = description
	_params["displayStyle"] = displayStyle
	_params["emailAddress"] = emailAddress
	_params["inProtocol"] = inProtocol
	_params["inServerName"] = inServerName
	_params["inServerPort"] = inServerPort
	_params["inUseSSL"] = inUseSSL
	_params["inUserName"] = inUserName
	_params["inPassword"] = inPassword
	_params["inReadInterval"] = inReadInterval
	_params["outEmailAddress"] = outEmailAddress
	_params["outCustom"] = outCustom
	_params["outServerName"] = outServerName
	_params["outServerPort"] = outServerPort
	_params["outUseSSL"] = outUseSSL
	_params["outUserName"] = outUserName
	_params["outPassword"] = outPassword
	_params["mailingListActive"] = mailingListActive
	_params["allowAnonymousEmail"] = allowAnonymousEmail
	_params["mergeWithParentCategory"] = mergeWithParentCategory
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbcategory/update-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

