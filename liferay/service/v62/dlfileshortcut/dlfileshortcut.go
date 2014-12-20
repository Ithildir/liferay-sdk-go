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

package dlfileshortcut

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFileShortcut(groupId int64, folderId int64, toFileEntryId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["toFileEntryId"] = toFileEntryId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileshortcut/add-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFileShortcut(fileShortcutId int64) error {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlfileshortcut/delete-file-shortcut": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFileShortcut(fileShortcutId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlfileshortcut/get-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateFileShortcut(fileShortcutId int64, folderId int64, toFileEntryId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId
	_params["folderId"] = folderId
	_params["toFileEntryId"] = toFileEntryId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileshortcut/update-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

