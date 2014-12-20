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

package dlfileentrytype

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFileEntryType(groupId int64, name string, description string, ddmStructureIds []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["description"] = description
	_params["ddmStructureIds"] = ddmStructureIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentrytype/add-file-entry-type": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddFileEntryType2(groupId int64, fileEntryTypeKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, ddmStructureIds []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["fileEntryTypeKey"] = fileEntryTypeKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["ddmStructureIds"] = ddmStructureIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentrytype/add-file-entry-type": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFileEntryType(fileEntryTypeId int64) error {
	_params := make(map[string]interface{})

	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlfileentrytype/delete-file-entry-type": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFileEntryType(fileEntryTypeId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlfileentrytype/get-file-entry-type": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntryTypes(groupIds []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/dlfileentrytype/get-file-entry-types": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntryTypes2(groupIds []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlfileentrytype/get-file-entry-types": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntryTypesCount(groupIds []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/dlfileentrytype/get-file-entry-types-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFolderFileEntryTypes(groupIds []interface{}, folderId int64, inherited bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["folderId"] = folderId
	_params["inherited"] = inherited

	_cmd := map[string]interface{}{
		"/dlfileentrytype/get-folder-file-entry-types": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupIds []interface{}, keywords string, includeBasicFileEntryType bool, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords
	_params["includeBasicFileEntryType"] = includeBasicFileEntryType
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/dlfileentrytype/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupIds []interface{}, keywords string, includeBasicFileEntryType bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords
	_params["includeBasicFileEntryType"] = includeBasicFileEntryType

	_cmd := map[string]interface{}{
		"/dlfileentrytype/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) UpdateFileEntryType(fileEntryTypeId int64, name string, description string, ddmStructureIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryTypeId"] = fileEntryTypeId
	_params["name"] = name
	_params["description"] = description
	_params["ddmStructureIds"] = ddmStructureIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentrytype/update-file-entry-type": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateFileEntryType2(fileEntryTypeId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, ddmStructureIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryTypeId"] = fileEntryTypeId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["ddmStructureIds"] = ddmStructureIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentrytype/update-file-entry-type": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

