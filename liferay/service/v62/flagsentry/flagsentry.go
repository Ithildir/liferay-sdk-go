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

package flagsentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddEntry(className string, classPK int64, reporterEmailAddress string, reportedUserId int64, contentTitle string, contentURL string, reason string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["reporterEmailAddress"] = reporterEmailAddress
	_params["reportedUserId"] = reportedUserId
	_params["contentTitle"] = contentTitle
	_params["contentURL"] = contentURL
	_params["reason"] = reason
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/flagsentry/add-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

