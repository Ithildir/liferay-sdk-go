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

package trashentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) DeleteEntries(entryIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["entryIds"] = entryIds

	_cmd := map[string]interface{}{
		"/trashentry/delete-entries": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteEntries2(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/trashentry/delete-entries": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteEntry(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/trashentry/delete-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteEntry2(className string, classPK int64) error {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/trashentry/delete-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetEntries(groupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/trashentry/get-entries": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetEntries2(groupId int64, start int, end int, obc *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/trashentry/get-entries": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveEntry(className string, classPK int64, destinationContainerModelId int64, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["destinationContainerModelId"] = destinationContainerModelId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/trashentry/move-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreEntry(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/trashentry/restore-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreEntry2(entryId int64, overrideClassPK int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["overrideClassPK"] = overrideClassPK
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/trashentry/restore-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

