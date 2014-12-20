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

package journalstructure

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddStructure(groupId int64, structureId string, autoStructureId bool, parentStructureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId
	_params["autoStructureId"] = autoStructureId
	_params["parentStructureId"] = parentStructureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalstructure/add-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyStructure(groupId int64, oldStructureId string, newStructureId string, autoStructureId bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["oldStructureId"] = oldStructureId
	_params["newStructureId"] = newStructureId
	_params["autoStructureId"] = autoStructureId

	_cmd := map[string]interface{}{
		"/journalstructure/copy-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteStructure(groupId int64, structureId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId

	_cmd := map[string]interface{}{
		"/journalstructure/delete-structure": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetStructure(groupId int64, structureId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId

	_cmd := map[string]interface{}{
		"/journalstructure/get-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetStructure2(groupId int64, structureId string, includeGlobalStructures bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId
	_params["includeGlobalStructures"] = includeGlobalStructures

	_cmd := map[string]interface{}{
		"/journalstructure/get-structure": _params,
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
		"/journalstructure/get-structures": _params,
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
		"/journalstructure/get-structures": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupIds []interface{}, keywords string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalstructure/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupIds []interface{}, structureId string, name string, description string, andOperator bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["structureId"] = structureId
	_params["name"] = name
	_params["description"] = description
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalstructure/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupIds []interface{}, keywords string) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords

	_cmd := map[string]interface{}{
		"/journalstructure/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupIds []interface{}, structureId string, name string, description string, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["structureId"] = structureId
	_params["name"] = name
	_params["description"] = description
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/journalstructure/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) UpdateStructure(groupId int64, structureId string, parentStructureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsd string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId
	_params["parentStructureId"] = parentStructureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsd"] = xsd
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalstructure/update-structure": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

