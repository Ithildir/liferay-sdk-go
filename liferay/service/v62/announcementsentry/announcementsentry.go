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

package announcementsentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddEntry(plid int64, classNameId int64, classPK int64, title string, content string, url string, type string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, priority int, alert bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["title"] = title
	_params["content"] = content
	_params["url"] = url
	_params["type"] = type
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["priority"] = priority
	_params["alert"] = alert

	_cmd := map[string]interface{}{
		"/announcementsentry/add-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddEntry2(plid int64, classNameId int64, classPK int64, title string, content string, url string, type string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, displayImmediately bool, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, priority int, alert bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["title"] = title
	_params["content"] = content
	_params["url"] = url
	_params["type"] = type
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["displayImmediately"] = displayImmediately
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["priority"] = priority
	_params["alert"] = alert

	_cmd := map[string]interface{}{
		"/announcementsentry/add-entry": _params,
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
		"/announcementsentry/delete-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetEntry(entryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId

	_cmd := map[string]interface{}{
		"/announcementsentry/get-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateEntry(entryId int64, title string, content string, url string, type string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, displayImmediately bool, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, priority int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["entryId"] = entryId
	_params["title"] = title
	_params["content"] = content
	_params["url"] = url
	_params["type"] = type
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["displayImmediately"] = displayImmediately
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["priority"] = priority

	_cmd := map[string]interface{}{
		"/announcementsentry/update-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

