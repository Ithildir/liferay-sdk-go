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

package journalfolder

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFolder(groupId int64, parentFolderId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfolder/add-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFolder(folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFolder2(folderId int64, includeTrashedEntries bool) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["includeTrashedEntries"] = includeTrashedEntries

	_cmd := map[string]interface{}{
		"/journalfolder/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFolder(folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFolderIds(groupId int64, folderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders2(groupId int64, parentFolderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders3(groupId int64, parentFolderId int64, status int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders4(groupId int64, parentFolderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders5(groupId int64, parentFolderId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndArticles(groupId int64, folderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-and-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndArticles2(groupId int64, folderId int64, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-and-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndArticlesCount(groupId int64, folderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-and-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFoldersAndArticlesCount2(groupId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-and-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFoldersAndArticlesCount3(groupId int64, folderIds []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-and-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFoldersCount(groupId int64, parentFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFoldersCount2(groupId int64, parentFolderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalfolder/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetSubfolderIds(folderIds []interface{}, groupId int64, folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderIds"] = folderIds
	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/get-subfolder-ids": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetSubfolderIds2(groupId int64, folderId int64, recurse bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["recurse"] = recurse

	_cmd := map[string]interface{}{
		"/journalfolder/get-subfolder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) MoveFolder(folderId int64, parentFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["parentFolderId"] = parentFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfolder/move-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFolderFromTrash(folderId int64, parentFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["parentFolderId"] = parentFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfolder/move-folder-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFolderToTrash(folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/move-folder-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreFolderFromTrash(folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalfolder/restore-folder-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateFolder(folderId int64, parentFolderId int64, name string, description string, mergeWithParentFolder bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	_params["mergeWithParentFolder"] = mergeWithParentFolder
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfolder/update-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

