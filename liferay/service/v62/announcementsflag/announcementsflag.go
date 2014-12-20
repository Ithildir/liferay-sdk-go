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

package announcementsflag

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFlag(entryId int64, value int) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["value"] = value

	_cmd := map[string]interface{}{
		"/announcementsflag/add-flag": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFlag(flagId int64) error {
	_params := make(map[string]interface{})

	_params["flagId"] = flagId

	_cmd := map[string]interface{}{
		"/announcementsflag/delete-flag": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFlag(entryId int64, value int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["value"] = value

	_cmd := map[string]interface{}{
		"/announcementsflag/get-flag": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

