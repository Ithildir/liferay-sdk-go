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

package assetcategory

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCategory(title string, vocabularyId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["title"] = title
	_params["vocabularyId"] = vocabularyId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetcategory/add-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddCategory2(parentCategoryId int64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, vocabularyId int64, categoryProperties []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["vocabularyId"] = vocabularyId
	_params["categoryProperties"] = categoryProperties
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetcategory/add-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteCategories(categoryIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["categoryIds"] = categoryIds

	_cmd := map[string]interface{}{
		"/assetcategory/delete-categories": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteCategories2(categoryIds []interface{}, serviceContext *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryIds"] = categoryIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetcategory/delete-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) DeleteCategory(categoryId int64) error {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/assetcategory/delete-category": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategories(className string, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/assetcategory/get-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategory(categoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId

	_cmd := map[string]interface{}{
		"/assetcategory/get-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetChildCategories(parentCategoryId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId

	_cmd := map[string]interface{}{
		"/assetcategory/get-child-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetChildCategories2(parentCategoryId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-child-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetJsonSearch(groupId int64, name string, vocabularyIds []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["vocabularyIds"] = vocabularyIds
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assetcategory/get-json-search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetJsonVocabularyCategories(vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-json-vocabulary-categories": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetJsonVocabularyCategories2(groupId int64, name string, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-json-vocabulary-categories": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyCategories(vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyCategories2(parentCategoryId int64, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentCategoryId"] = parentCategoryId
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyCategories3(groupId int64, name string, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyCategoriesCount(groupId int64, vocabularyId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["vocabularyId"] = vocabularyId

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetVocabularyCategoriesCount2(groupId int64, name string, vocabularyId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["vocabularyId"] = vocabularyId

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetVocabularyCategoriesDisplay(vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyCategoriesDisplay2(groupId int64, name string, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-categories-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyRootCategories(vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-root-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyRootCategories2(groupId int64, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-root-categories": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularyRootCategoriesCount(groupId int64, vocabularyId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["vocabularyId"] = vocabularyId

	_cmd := map[string]interface{}{
		"/assetcategory/get-vocabulary-root-categories-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) MoveCategory(categoryId int64, parentCategoryId int64, vocabularyId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["parentCategoryId"] = parentCategoryId
	_params["vocabularyId"] = vocabularyId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetcategory/move-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search(groupId int64, name string, categoryProperties []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["categoryProperties"] = categoryProperties
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assetcategory/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(groupIds []interface{}, name string, vocabularyIds []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["name"] = name
	_params["vocabularyIds"] = vocabularyIds
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assetcategory/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search3(groupId int64, keywords string, vocabularyId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["keywords"] = keywords
	_params["vocabularyId"] = vocabularyId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetcategory/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateCategory(categoryId int64, parentCategoryId int64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, vocabularyId int64, categoryProperties []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["parentCategoryId"] = parentCategoryId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["vocabularyId"] = vocabularyId
	_params["categoryProperties"] = categoryProperties
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetcategory/update-category": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

