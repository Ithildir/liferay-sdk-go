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

package membershiprequest

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddMembershipRequest(groupId int64, comments string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["comments"] = comments
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/membershiprequest/add-membership-request": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteMembershipRequests(groupId int64, statusId int) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["statusId"] = statusId

	_cmd := map[string]interface{}{
		"/membershiprequest/delete-membership-requests": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetMembershipRequest(membershipRequestId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["membershipRequestId"] = membershipRequestId

	_cmd := map[string]interface{}{
		"/membershiprequest/get-membership-request": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateStatus(membershipRequestId int64, reviewComments string, statusId int, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["membershipRequestId"] = membershipRequestId
	_params["reviewComments"] = reviewComments
	_params["statusId"] = statusId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/membershiprequest/update-status": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

