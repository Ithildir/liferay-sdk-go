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

package staging

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) CleanUpStagingRequest(stagingRequestId int64) error {
	_params := make(map[string]interface{})

	_params["stagingRequestId"] = stagingRequestId

	_cmd := map[string]interface{}{
		"/staging/clean-up-staging-request": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CreateStagingRequest(groupId int64, checksum string) (int64, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["checksum"] = checksum

	_cmd := map[string]interface{}{
		"/staging/create-staging-request": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64)
	}

	return v, err
}

func (s *Service) PublishStagingRequest(stagingRequestId int64, privateLayout bool, parameterMap map[string]interface{}) error {
	_params := make(map[string]interface{})

	_params["stagingRequestId"] = stagingRequestId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap

	_cmd := map[string]interface{}{
		"/staging/publish-staging-request": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateStagingRequest(stagingRequestId int64, fileName string, bytes []byte) error {
	_params := make(map[string]interface{})

	_params["stagingRequestId"] = stagingRequestId
	_params["fileName"] = fileName
	_params["bytes"] = bytes

	_cmd := map[string]interface{}{
		"/staging/update-staging-request": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ValidateStagingRequest(stagingRequestId int64, privateLayout bool, parameterMap map[string]interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["stagingRequestId"] = stagingRequestId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap

	_cmd := map[string]interface{}{
		"/staging/validate-staging-request": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

