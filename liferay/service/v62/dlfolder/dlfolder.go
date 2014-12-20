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

package dlfolder

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFolder(groupId int64, repositoryId int64, mountPoint bool, parentFolderId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["repositoryId"] = repositoryId
	_params["mountPoint"] = mountPoint
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfolder/add-folder": _params,
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
		"/dlfolder/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFolder2(folderId int64, includeTrashedEntries bool) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["includeTrashedEntries"] = includeTrashedEntries

	_cmd := map[string]interface{}{
		"/dlfolder/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFolder3(groupId int64, parentFolderId int64, name string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/dlfolder/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFileEntriesAndFileShortcuts(groupId int64, folderId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlfolder/get-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntriesAndFileShortcutsCount(groupId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfolder/get-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFileEntriesAndFileShortcutsCount2(groupId int64, folderId int64, status int, mimeTypes []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes

	_cmd := map[string]interface{}{
		"/dlfolder/get-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFolder(folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/get-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFolder2(groupId int64, parentFolderId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/dlfolder/get-folder": _params,
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
		"/dlfolder/get-folder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders(groupId int64, parentFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders2(groupId int64, parentFolderId int64, status int, includeMountfolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status
	_params["includeMountfolders"] = includeMountfolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcuts(groupId int64, folderId int64, status int, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-and-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcuts2(groupId int64, folderId int64, status int, mimeTypes []interface{}, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-and-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcutsCount(groupId int64, folderId int64, status int, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-and-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcutsCount2(groupId int64, folderId int64, status int, mimeTypes []interface{}, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-and-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersCount(groupId int64, parentFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersCount2(groupId int64, parentFolderId int64, status int, includeMountfolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status
	_params["includeMountfolders"] = includeMountfolders

	_cmd := map[string]interface{}{
		"/dlfolder/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetMountFolders(groupId int64, parentFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfolder/get-mount-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetMountFoldersCount(groupId int64, parentFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlfolder/get-mount-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetSubfolderIds(folderIds []interface{}, groupId int64, folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderIds"] = folderIds
	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/get-subfolder-ids": _params,
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
		"/dlfolder/get-subfolder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) HasFolderLock(folderId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/has-folder-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) HasInheritableLock(folderId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/has-inheritable-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) IsFolderLocked(folderId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/is-folder-locked": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) LockFolder(folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfolder/lock-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) LockFolder2(folderId int64, owner string, inheritable bool, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["owner"] = owner
	_params["inheritable"] = inheritable
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlfolder/lock-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFolder(folderId int64, parentFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["parentFolderId"] = parentFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfolder/move-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RefreshFolderLock(lockUuid string, companyId int64, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["lockUuid"] = lockUuid
	_params["companyId"] = companyId
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlfolder/refresh-folder-lock": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UnlockFolder(folderId int64, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfolder/unlock-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnlockFolder2(groupId int64, parentFolderId int64, name string, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfolder/unlock-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateFolder(folderId int64, name string, description string, defaultFileEntryTypeId int64, fileEntryTypeIds []interface{}, overrideFileEntryTypes bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["name"] = name
	_params["description"] = description
	_params["defaultFileEntryTypeId"] = defaultFileEntryTypeId
	_params["fileEntryTypeIds"] = fileEntryTypeIds
	_params["overrideFileEntryTypes"] = overrideFileEntryTypes
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfolder/update-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) VerifyInheritableLock(folderId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfolder/verify-inheritable-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

