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

package dlfileversion

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetFileVersion(fileVersionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileVersionId"] = fileVersionId

	_cmd := map[string]interface{}{
		"/dlfileversion/get-file-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileVersions(fileEntryId int64, status int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfileversion/get-file-versions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileVersionsCount(fileEntryId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfileversion/get-file-versions-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetLatestFileVersion(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileversion/get-latest-file-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

