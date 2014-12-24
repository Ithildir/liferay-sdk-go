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

package dlapp

import (
	"io"

	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFileEntry(repositoryId int64, folderId int64, sourceFileName string, mimeType string, title string, description string, changeLog string, bytes []byte, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["sourceFileName"] = sourceFileName
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["changeLog"] = changeLog
	_params["bytes"] = liferay.ToJSONString(bytes)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/add-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddFileEntry2(repositoryId int64, folderId int64, sourceFileName string, mimeType string, title string, description string, changeLog string, file io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["sourceFileName"] = sourceFileName
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["changeLog"] = changeLog
	_params["file"] = file
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/add-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddFileShortcut(repositoryId int64, folderId int64, toFileEntryId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["toFileEntryId"] = toFileEntryId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/add-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddFolder(repositoryId int64, parentFolderId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/add-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddTempFileEntry(groupId int64, folderId int64, fileName string, tempFolderName string, file io.Reader, mimeType string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["fileName"] = fileName
	_params["tempFolderName"] = tempFolderName
	_params["file"] = file
	_params["mimeType"] = mimeType

	_cmd := map[string]interface{}{
		"/dlapp/add-temp-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CancelCheckOut(fileEntryId int64) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/cancel-check-out": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckInFileEntry(fileEntryId int64, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/check-in-file-entry": _params,
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
		"/dlapp/check-in-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckInFileEntry3(fileEntryId int64, majorVersion bool, changeLog string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["majorVersion"] = majorVersion
	_params["changeLog"] = changeLog
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/check-in-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckOutFileEntry(fileEntryId int64, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/check-out-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckOutFileEntry2(fileEntryId int64, owner string, expirationTime int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["owner"] = owner
	_params["expirationTime"] = expirationTime
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/check-out-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyFolder(repositoryId int64, sourceFolderId int64, parentFolderId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["sourceFolderId"] = sourceFolderId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/copy-folder": _params,
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
		"/dlapp/delete-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFileEntryByTitle(repositoryId int64, folderId int64, title string) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/dlapp/delete-file-entry-by-title": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFileShortcut(fileShortcutId int64) error {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlapp/delete-file-shortcut": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFileVersion(fileEntryId int64, version string) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/dlapp/delete-file-version": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFolder(folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFolder2(repositoryId int64, parentFolderId int64, name string) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/dlapp/delete-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteTempFileEntry(groupId int64, folderId int64, fileName string, tempFolderName string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["fileName"] = fileName
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/dlapp/delete-temp-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFileEntries(repositoryId int64, folderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries2(repositoryId int64, folderId int64, fileEntryTypeId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries3(repositoryId int64, folderId int64, mimeTypes []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["mimeTypes"] = mimeTypes

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries4(repositoryId int64, folderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries5(repositoryId int64, folderId int64, fileEntryTypeId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries6(repositoryId int64, folderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntries7(repositoryId int64, folderId int64, fileEntryTypeId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntriesAndFileShortcuts(repositoryId int64, folderId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFileEntriesAndFileShortcutsCount(repositoryId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFileEntriesAndFileShortcutsCount2(repositoryId int64, folderId int64, status int, mimeTypes []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFileEntriesCount(repositoryId int64, folderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFileEntriesCount2(repositoryId int64, folderId int64, fileEntryTypeId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFileEntry(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-entry": _params,
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
		"/dlapp/get-file-entry": _params,
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
		"/dlapp/get-file-entry-by-uuid-and-group-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFileShortcut(fileShortcutId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlapp/get-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFolder(folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/get-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFolder2(repositoryId int64, parentFolderId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/dlapp/get-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFolders(repositoryId int64, parentFolderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders2(repositoryId int64, parentFolderId int64, includeMountFolders bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders3(repositoryId int64, parentFolderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders4(repositoryId int64, parentFolderId int64, includeMountFolders bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders5(repositoryId int64, parentFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders6(repositoryId int64, parentFolderId int64, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFolders7(repositoryId int64, parentFolderId int64, status int, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcuts(repositoryId int64, folderId int64, status int, includeMountFolders bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-and-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcuts2(repositoryId int64, folderId int64, status int, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-and-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcuts3(repositoryId int64, folderId int64, status int, mimeTypes []interface{}, includeMountFolders bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes
	_params["includeMountFolders"] = includeMountFolders
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-and-file-entries-and-file-shortcuts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcutsCount(repositoryId int64, folderId int64, status int, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-and-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersAndFileEntriesAndFileShortcutsCount2(repositoryId int64, folderId int64, status int, mimeTypes []interface{}, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["status"] = status
	_params["mimeTypes"] = mimeTypes
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-and-file-entries-and-file-shortcuts-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersCount(repositoryId int64, parentFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersCount2(repositoryId int64, parentFolderId int64, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersCount3(repositoryId int64, parentFolderId int64, status int, includeMountFolders bool) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["status"] = status
	_params["includeMountFolders"] = includeMountFolders

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetFoldersFileEntriesCount(repositoryId int64, folderIds []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderIds"] = folderIds
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlapp/get-folders-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupFileEntries(groupId int64, userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntries2(groupId int64, userId int64, rootFolderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntries3(groupId int64, userId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntries4(groupId int64, userId int64, rootFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntries5(groupId int64, userId int64, rootFolderId int64, mimeTypes []interface{}, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
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
		"/dlapp/get-group-file-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupFileEntriesCount(groupId int64, userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupFileEntriesCount2(groupId int64, userId int64, rootFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupFileEntriesCount3(groupId int64, userId int64, rootFolderId int64, mimeTypes []interface{}, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["mimeTypes"] = mimeTypes
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/dlapp/get-group-file-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetMountFolders(repositoryId int64, parentFolderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlapp/get-mount-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetMountFolders2(repositoryId int64, parentFolderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/get-mount-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetMountFolders3(repositoryId int64, parentFolderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/dlapp/get-mount-folders": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetMountFoldersCount(repositoryId int64, parentFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/dlapp/get-mount-folders-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetSubfolderIds(repositoryId int64, folderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/get-subfolder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSubfolderIds2(repositoryId int64, folderId int64, recurse bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["recurse"] = recurse

	_cmd := map[string]interface{}{
		"/dlapp/get-subfolder-ids": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetSubfolderIds3(repositoryId int64, folderIds []interface{}, folderId int64) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderIds"] = folderIds
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/get-subfolder-ids": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetTempFileEntryNames(groupId int64, folderId int64, tempFolderName string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/dlapp/get-temp-file-entry-names": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) LockFileEntry(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/lock-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) LockFileEntry2(fileEntryId int64, owner string, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["owner"] = owner
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlapp/lock-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) LockFolder(repositoryId int64, folderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/lock-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) LockFolder2(repositoryId int64, folderId int64, owner string, inheritable bool, expirationTime int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["owner"] = owner
	_params["inheritable"] = inheritable
	_params["expirationTime"] = expirationTime

	_cmd := map[string]interface{}{
		"/dlapp/lock-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFileEntry(fileEntryId int64, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/move-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFileEntryFromTrash(fileEntryId int64, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/move-file-entry-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFileEntryToTrash(fileEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/move-file-entry-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFileShortcutFromTrash(fileShortcutId int64, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/move-file-shortcut-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveFileShortcutToTrash(fileShortcutId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlapp/move-file-shortcut-to-trash": _params,
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
		"/dlapp/move-folder": _params,
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
		"/dlapp/move-folder-from-trash": _params,
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
		"/dlapp/move-folder-to-trash": _params,
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
		"/dlapp/refresh-file-entry-lock": _params,
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
		"/dlapp/refresh-folder-lock": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreFileEntryFromTrash(fileEntryId int64) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/restore-file-entry-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreFileShortcutFromTrash(fileShortcutId int64) error {
	_params := make(map[string]interface{})

	_params["fileShortcutId"] = fileShortcutId

	_cmd := map[string]interface{}{
		"/dlapp/restore-file-shortcut-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreFolderFromTrash(folderId int64) error {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/restore-folder-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RevertFileEntry(fileEntryId int64, version string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["version"] = version
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/revert-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Search(repositoryId int64, searchContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	liferay.MangleObjectWrapper(_params, "searchContext", "com.liferay.portal.kernel.search.SearchContext", searchContext)

	_cmd := map[string]interface{}{
		"/dlapp/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search2(repositoryId int64, searchContext *liferay.ObjectWrapper, query *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	liferay.MangleObjectWrapper(_params, "searchContext", "com.liferay.portal.kernel.search.SearchContext", searchContext)
	liferay.MangleObjectWrapper(_params, "query", "com.liferay.portal.kernel.search.Query", query)

	_cmd := map[string]interface{}{
		"/dlapp/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search3(repositoryId int64, creatorUserId int64, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["creatorUserId"] = creatorUserId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search4(repositoryId int64, creatorUserId int64, folderId int64, mimeTypes []interface{}, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["creatorUserId"] = creatorUserId
	_params["folderId"] = folderId
	_params["mimeTypes"] = mimeTypes
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/dlapp/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SubscribeFileEntryType(groupId int64, fileEntryTypeId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlapp/subscribe-file-entry-type": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SubscribeFolder(groupId int64, folderId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/subscribe-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnlockFileEntry(fileEntryId int64) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId

	_cmd := map[string]interface{}{
		"/dlapp/unlock-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnlockFileEntry2(fileEntryId int64, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/unlock-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnlockFolder(repositoryId int64, folderId int64, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/unlock-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnlockFolder2(repositoryId int64, parentFolderId int64, name string, lockUuid string) error {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["parentFolderId"] = parentFolderId
	_params["name"] = name
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/unlock-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeFileEntryType(groupId int64, fileEntryTypeId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["fileEntryTypeId"] = fileEntryTypeId

	_cmd := map[string]interface{}{
		"/dlapp/unsubscribe-file-entry-type": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeFolder(groupId int64, folderId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/dlapp/unsubscribe-folder": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateFileEntry(fileEntryId int64, sourceFileName string, mimeType string, title string, description string, changeLog string, majorVersion bool, bytes []byte, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["sourceFileName"] = sourceFileName
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["changeLog"] = changeLog
	_params["majorVersion"] = majorVersion
	_params["bytes"] = liferay.ToJSONString(bytes)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/update-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateFileEntry2(fileEntryId int64, sourceFileName string, mimeType string, title string, description string, changeLog string, majorVersion bool, file io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["sourceFileName"] = sourceFileName
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["changeLog"] = changeLog
	_params["majorVersion"] = majorVersion
	_params["file"] = file
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/update-file-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateFileEntryAndCheckIn(fileEntryId int64, sourceFileName string, mimeType string, title string, description string, changeLog string, majorVersion bool, file io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["fileEntryId"] = fileEntryId
	_params["sourceFileName"] = sourceFileName
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["changeLog"] = changeLog
	_params["majorVersion"] = majorVersion
	_params["file"] = file
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/update-file-entry-and-check-in": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

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
		"/dlapp/update-file-shortcut": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateFolder(folderId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["folderId"] = folderId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/dlapp/update-folder": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) VerifyFileEntryCheckOut(repositoryId int64, fileEntryId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/verify-file-entry-check-out": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) VerifyFileEntryLock(repositoryId int64, fileEntryId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["fileEntryId"] = fileEntryId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/verify-file-entry-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) VerifyInheritableLock(repositoryId int64, folderId int64, lockUuid string) (bool, error) {
	_params := make(map[string]interface{})

	_params["repositoryId"] = repositoryId
	_params["folderId"] = folderId
	_params["lockUuid"] = lockUuid

	_cmd := map[string]interface{}{
		"/dlapp/verify-inheritable-lock": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

