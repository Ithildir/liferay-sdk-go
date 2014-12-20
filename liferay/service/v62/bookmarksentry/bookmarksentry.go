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

package bookmarksentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddEntry(groupId int64, folderId int64, name string, url string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["name"] = name
	_params["url"] = url
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/bookmarksentry/add-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteEntry(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/delete-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetEntries(groupId int64, folderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetEntries2(groupId int64, folderId int64, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetEntriesCount(groupId int64, folderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetEntriesCount2(groupId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetEntry(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersEntriesCount(groupId int64, folderIds []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderIds"] = folderIds

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-folders-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupEntries(groupId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries2(groupId int64, userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries3(groupId int64, userId int64, rootFolderId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntriesCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupEntriesCount2(groupId int64, userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupEntriesCount3(groupId int64, userId int64, rootFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId

	_cmd := map[string]interface{}{
		"/bookmarksentry/get-group-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) MoveEntry(entryId int64, parentFolderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/bookmarksentry/move-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveEntryFromTrash(entryId int64, parentFolderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["parentFolderId"] = parentFolderId

	_cmd := map[string]interface{}{
		"/bookmarksentry/move-entry-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveEntryToTrash(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/move-entry-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) OpenEntry(entry *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	liferay.MangleObjectWrapper(_params, "entry", "com.liferay.portlet.bookmarks.model.BookmarksEntry", entry)

	_cmd := map[string]interface{}{
		"/bookmarksentry/open-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) OpenEntry2(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/open-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreEntryFromTrash(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/restore-entry-from-trash": _params,
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
		"/bookmarksentry/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SubscribeEntry(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/subscribe-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeEntry(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/bookmarksentry/unsubscribe-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateEntry(entryId int64, groupId int64, folderId int64, name string, url string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["name"] = name
	_params["url"] = url
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/bookmarksentry/update-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

