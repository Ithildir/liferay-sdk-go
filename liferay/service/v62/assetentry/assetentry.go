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

package assetentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetCompanyEntries(companyId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assetentry/get-company-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyEntriesCount(companyId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/assetentry/get-company-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetEntries(entryQuery *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	liferay.MangleObjectWrapper(_params, "entryQuery", "com.liferay.portlet.asset.service.persistence.AssetEntryQuery", entryQuery)

	_cmd := map[string]interface{}{
		"/assetentry/get-entries": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetEntriesCount(entryQuery *liferay.ObjectWrapper) (int, error) {
	_params := make(map[string]interface{})

	liferay.MangleObjectWrapper(_params, "entryQuery", "com.liferay.portlet.asset.service.persistence.AssetEntryQuery", entryQuery)

	_cmd := map[string]interface{}{
		"/assetentry/get-entries-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetEntry(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/assetentry/get-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) IncrementViewCounter(className string, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/assetentry/increment-view-counter": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateEntry(groupId int64, className string, classPK int64, classUuid string, classTypeId int64, categoryIds []interface{}, tagNames []interface{}, visible bool, startDate int64, endDate int64, expirationDate int64, mimeType string, title string, description string, summary string, url string, layoutUuid string, height int, width int, priority *liferay.ObjectWrapper, sync bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["classUuid"] = classUuid
	_params["classTypeId"] = classTypeId
	_params["categoryIds"] = categoryIds
	_params["tagNames"] = tagNames
	_params["visible"] = visible
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["expirationDate"] = expirationDate
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["summary"] = summary
	_params["url"] = url
	_params["layoutUuid"] = layoutUuid
	_params["height"] = height
	_params["width"] = width
	liferay.MangleObjectWrapper(_params, "priority", "java.lang.Integer", priority)
	_params["sync"] = sync

	_cmd := map[string]interface{}{
		"/assetentry/update-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateEntry2(groupId int64, className string, classPK int64, classUuid string, classTypeId int64, categoryIds []interface{}, tagNames []interface{}, visible bool, startDate int64, endDate int64, publishDate int64, expirationDate int64, mimeType string, title string, description string, summary string, url string, layoutUuid string, height int, width int, priority *liferay.ObjectWrapper, sync bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["classUuid"] = classUuid
	_params["classTypeId"] = classTypeId
	_params["categoryIds"] = categoryIds
	_params["tagNames"] = tagNames
	_params["visible"] = visible
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["publishDate"] = publishDate
	_params["expirationDate"] = expirationDate
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["summary"] = summary
	_params["url"] = url
	_params["layoutUuid"] = layoutUuid
	_params["height"] = height
	_params["width"] = width
	liferay.MangleObjectWrapper(_params, "priority", "java.lang.Integer", priority)
	_params["sync"] = sync

	_cmd := map[string]interface{}{
		"/assetentry/update-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateEntry3(groupId int64, createDate int64, modifiedDate int64, className string, classPK int64, classUuid string, classTypeId int64, categoryIds []interface{}, tagNames []interface{}, visible bool, startDate int64, endDate int64, expirationDate int64, mimeType string, title string, description string, summary string, url string, layoutUuid string, height int, width int, priority *liferay.ObjectWrapper, sync bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["createDate"] = createDate
	_params["modifiedDate"] = modifiedDate
	_params["className"] = className
	_params["classPK"] = classPK
	_params["classUuid"] = classUuid
	_params["classTypeId"] = classTypeId
	_params["categoryIds"] = categoryIds
	_params["tagNames"] = tagNames
	_params["visible"] = visible
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["expirationDate"] = expirationDate
	_params["mimeType"] = mimeType
	_params["title"] = title
	_params["description"] = description
	_params["summary"] = summary
	_params["url"] = url
	_params["layoutUuid"] = layoutUuid
	_params["height"] = height
	_params["width"] = width
	liferay.MangleObjectWrapper(_params, "priority", "java.lang.Integer", priority)
	_params["sync"] = sync

	_cmd := map[string]interface{}{
		"/assetentry/update-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

