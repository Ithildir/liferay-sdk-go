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

package repository

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddRepository(groupId int64, classNameId int64, parentFolderId int64, name string, description string, portletId string, typeSettingsProperties *liferay.ObjectWrapper, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	_params["portletId"] = portletId
	liferay.MangleObjectWrapper(_params, "typeSettingsProperties", "com.liferay.portal.kernel.util.UnicodeProperties", typeSettingsProperties)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/repository/add-repository": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CheckRepository(repositoryId int64) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/check-repository": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteRepository(repositoryId int64) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/delete-repository": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetLocalRepositoryImpl(repositoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/get-local-repository-impl": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetLocalRepositoryImpl2(folderId int64, fileEntryId int64, fileVersionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["fileEntryId"] = fileEntryId
	_params["fileVersionId"] = fileVersionId

	_cmd := map[string]interface{}{
		"/repository/get-local-repository-impl": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRepository(repositoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/get-repository": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRepositoryImpl(repositoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/get-repository-impl": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetRepositoryImpl2(folderId int64, fileEntryId int64, fileVersionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["fileEntryId"] = fileEntryId
	_params["fileVersionId"] = fileVersionId

	_cmd := map[string]interface{}{
		"/repository/get-repository-impl": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetSupportedConfigurations(classNameId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNameId"] = classNameId

	_cmd := map[string]interface{}{
		"/repository/get-supported-configurations": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSupportedParameters(classNameId int64, configuration string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNameId"] = classNameId
	_params["configuration"] = configuration

	_cmd := map[string]interface{}{
		"/repository/get-supported-parameters": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTypeSettingsProperties(repositoryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId

	_cmd := map[string]interface{}{
		"/repository/get-type-settings-properties": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateRepository(repositoryId int64, name string, description string) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["name"] = name
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/repository/update-repository": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

