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

package ddmstructure

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddStructure(groupId int64, parentStructureId int64, classNameId int64, structureKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, storageType string, type int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentStructureId"] = parentStructureId
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	_params["storageType"] = storageType
	_params["type"] = type
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/add-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddStructure2(userId int64, groupId int64, parentStructureKey string, classNameId int64, structureKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, storageType string, type int, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["parentStructureKey"] = parentStructureKey
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	_params["storageType"] = storageType
	_params["type"] = type
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/add-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddStructure3(userId int64, groupId int64, classNameId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/add-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyStructure(structureId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["structureId"] = structureId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/copy-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyStructure2(structureId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["structureId"] = structureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/copy-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteStructure(structureId int64) error {
	_params := make(map[string]interface{})

	_params["structureId"] = structureId

	_cmd := map[string]interface{}{
		"/ddmstructure/delete-structure": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) FetchStructure(groupId int64, classNameId int64, structureKey string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey

	_cmd := map[string]interface{}{
		"/ddmstructure/fetch-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetStructure(structureId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["structureId"] = structureId

	_cmd := map[string]interface{}{
		"/ddmstructure/get-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetStructure2(groupId int64, classNameId int64, structureKey string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey

	_cmd := map[string]interface{}{
		"/ddmstructure/get-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetStructure3(groupId int64, classNameId int64, structureKey string, includeGlobalStructures bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey
	_params["includeGlobalStructures"] = includeGlobalStructures

	_cmd := map[string]interface{}{
		"/ddmstructure/get-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetStructures(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/ddmstructure/get-structures": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetStructures2(groupIds []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/ddmstructure/get-structures": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupIds []interface{}, classNameIds []interface{}, name string, description string, storageType string, type int, andOperator bool, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["name"] = name
	_params["description"] = description
	_params["storageType"] = storageType
	_params["type"] = type
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmstructure/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupIds []interface{}, classNameIds []interface{}, keywords string, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["keywords"] = keywords
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmstructure/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupIds []interface{}, classNameIds []interface{}, keywords string) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["keywords"] = keywords

	_cmd := map[string]interface{}{
		"/ddmstructure/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupIds []interface{}, classNameIds []interface{}, name string, description string, storageType string, type int, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["name"] = name
	_params["description"] = description
	_params["storageType"] = storageType
	_params["type"] = type
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/ddmstructure/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) UpdateStructure(structureId int64, parentStructureId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["structureId"] = structureId
	_params["parentStructureId"] = parentStructureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/update-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateStructure2(groupId int64, parentStructureId int64, classNameId int64, structureKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentStructureId"] = parentStructureId
	_params["classNameId"] = classNameId
	_params["structureKey"] = structureKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmstructure/update-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

