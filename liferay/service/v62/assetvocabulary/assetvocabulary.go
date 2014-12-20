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

package assetvocabulary

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddVocabulary(title string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["title"] = title
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/add-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddVocabulary2(titleMap map[string]interface{}, descriptionMap map[string]interface{}, settings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["settings"] = settings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/add-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddVocabulary3(title string, titleMap map[string]interface{}, descriptionMap map[string]interface{}, settings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["title"] = title
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["settings"] = settings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/add-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteVocabularies(vocabularyIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["vocabularyIds"] = vocabularyIds

	_cmd := map[string]interface{}{
		"/assetvocabulary/delete-vocabularies": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteVocabularies2(vocabularyIds []interface{}, serviceContext *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyIds"] = vocabularyIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/delete-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) DeleteVocabulary(vocabularyId int64) error {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId

	_cmd := map[string]interface{}{
		"/assetvocabulary/delete-vocabulary": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCompanyVocabularies(companyId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-company-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabularies(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabularies2(groupId int64, createDefaultVocabulary bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["createDefaultVocabulary"] = createDefaultVocabulary

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabularies3(groupId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabularies4(groupId int64, name string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabulariesCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupVocabulariesCount2(groupId int64, name string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupVocabulariesDisplay(groupId int64, name string, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroupVocabulariesDisplay2(groupId int64, name string, start int, end int, addDefaultVocabulary bool, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end
	_params["addDefaultVocabulary"] = addDefaultVocabulary
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-group-vocabularies-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroupsVocabularies(groupIds []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-groups-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupsVocabularies2(groupIds []interface{}, className string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["className"] = className

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-groups-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetJsonGroupVocabularies(groupId int64, name string, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-json-group-vocabularies": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetVocabularies(vocabularyIds []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyIds"] = vocabularyIds

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-vocabularies": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetVocabulary(vocabularyId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId

	_cmd := map[string]interface{}{
		"/assetvocabulary/get-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateVocabulary(vocabularyId int64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, settings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["settings"] = settings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/update-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateVocabulary2(vocabularyId int64, title string, titleMap map[string]interface{}, descriptionMap map[string]interface{}, settings string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["vocabularyId"] = vocabularyId
	_params["title"] = title
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["settings"] = settings
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assetvocabulary/update-vocabulary": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

