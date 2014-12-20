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

package dlfileentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) CancelCheckOut(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/cancel-check-out": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CheckInFileEntry(fileEntryId int64, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfileentry/check-in-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckInFileEntry2(fileEntryId int64, lockUuid string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/check-in-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckInFileEntry3(fileEntryId int64, major bool, changeLog string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["major"] = major
	_params["changeLog"] = changeLog
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/check-in-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckOutFileEntry(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/check-out-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CheckOutFileEntry2(fileEntryId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/check-out-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CheckOutFileEntry3(fileEntryId int64, owner string, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["owner"] = owner
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlfileentry/check-out-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CheckOutFileEntry4(fileEntryId int64, owner string, expirationTime int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["owner"] = owner
	_params["expirationTime"] = expirationTime
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/check-out-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyFileEntry(groupId int64, repositoryId int64, fileEntryId int64, destFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["repositoryId"] = repositoryId
	_params["fileEntryId"] = fileEntryId
	_params["destFolderId"] = destFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/copy-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFileEntry(fileEntryId int64) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/delete-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFileEntry2(groupId int64, folderId int64, title string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/dlfileentry/delete-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFileVersion(fileEntryId int64, version string) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/dlfileentry/delete-file-version": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) FetchFileEntryByImageId(imageId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["imageId"] = imageId

	_cmd := map[string]interface{}{
		"/dlfileentry/fetch-file-entry-by-image-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries(groupId int64, folderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries2(groupId int64, folderId int64, fileEntryTypeId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries3(groupId int64, folderId int64, mimeTypes []interface{}, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["mimeTypes"] = mimeTypes
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries4(groupId int64, folderId int64, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntriesCount(groupId int64, folderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFileEntriesCount2(groupId int64, folderId int64, fileEntryTypeId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFileEntriesCount3(groupId int64, folderId int64, mimeTypes []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["mimeTypes"] = mimeTypes

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFileEntriesCount4(groupId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetFileEntry(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntry2(groupId int64, folderId int64, title string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntryByUuidAndGroupId(uuid string, groupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["uuid"] = uuid
	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entry-by-uuid-and-group-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntryLock(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-file-entry-lock": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersFileEntriesCount(groupId int64, folderIds []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfileentry/get-folders-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupFileEntries(groupId int64, userId int64, rootFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntries2(groupId int64, userId int64, rootFolderId int64, mimeTypes []interface{}, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["mimeTypes"] = mimeTypes
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlfileentry/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntriesCount(groupId int64, userId int64, rootFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId

	_cmd := map[string]interface{}{
		"/dlfileentry/get-group-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupFileEntriesCount2(groupId int64, userId int64, rootFolderId int64, mimeTypes []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["mimeTypes"] = mimeTypes
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlfileentry/get-group-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) HasFileEntryLock(fileEntryId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/has-file-entry-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) IsFileEntryCheckedOut(fileEntryId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlfileentry/is-file-entry-checked-out": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) MoveFileEntry(fileEntryId int64, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/move-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RefreshFileEntryLock(lockUuid string, companyId int64, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["lockUuid"] = lockUuid
	_params["companyId"] = companyId
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlfileentry/refresh-file-entry-lock": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RevertFileEntry(fileEntryId int64, version string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["version"] = version
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlfileentry/revert-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Search(groupId int64, creatorUserId int64, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["creatorUserId"] = creatorUserId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlfileentry/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search2(groupId int64, creatorUserId int64, folderId int64, mimeTypes []interface{}, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["creatorUserId"] = creatorUserId
	_params["folderId"] = folderId
	_params["mimeTypes"] = mimeTypes
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlfileentry/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) VerifyFileEntryCheckOut(fileEntryId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfileentry/verify-file-entry-check-out": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) VerifyFileEntryLock(fileEntryId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlfileentry/verify-file-entry-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

