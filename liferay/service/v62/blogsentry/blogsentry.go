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

package blogsentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) DeleteEntry(entryId int64) error {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/blogsentry/delete-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCompanyEntries(companyId int64, displayDate int64, status int, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/blogsentry/get-company-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyEntriesRss(companyId int64, displayDate int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/blogsentry/get-company-entries-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetEntry(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/blogsentry/get-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetEntry2(groupId int64, urlTitle string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["urlTitle"] = urlTitle

	_cmd := map[string]interface{}{
		"/blogsentry/get-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries(groupId int64, status int, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries2(groupId int64, displayDate int64, status int, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries3(groupId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntries4(groupId int64, displayDate int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupEntriesCount(groupId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupEntriesCount2(groupId int64, displayDate int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["displayDate"] = displayDate
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupEntriesRss(groupId int64, displayDate int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/blogsentry/get-group-entries-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetGroupsEntries(companyId int64, groupId int64, displayDate int64, status int, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/blogsentry/get-groups-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationEntries(organizationId int64, displayDate int64, status int, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/blogsentry/get-organization-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationEntriesRss(organizationId int64, displayDate int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["displayDate"] = displayDate
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/blogsentry/get-organization-entries-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) MoveEntryToTrash(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/blogsentry/move-entry-to-trash": _params,
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
		"/blogsentry/restore-entry-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Subscribe(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/blogsentry/subscribe": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Unsubscribe(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/blogsentry/unsubscribe": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

