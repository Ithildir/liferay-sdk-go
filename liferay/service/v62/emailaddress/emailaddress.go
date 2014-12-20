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

package emailaddress

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddEmailAddress(className string, classPK int64, address string, typeId int, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["address"] = address
	_params["typeId"] = typeId
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/emailaddress/add-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddEmailAddress2(className string, classPK int64, address string, typeId int, primary bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["address"] = address
	_params["typeId"] = typeId
	_params["primary"] = primary
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/emailaddress/add-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteEmailAddress(emailAddressId int64) error {
	_params := make(map[string]interface{})

	_params["emailAddressId"] = emailAddressId

	_cmd := map[string]interface{}{
		"/emailaddress/delete-email-address": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetEmailAddress(emailAddressId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["emailAddressId"] = emailAddressId

	_cmd := map[string]interface{}{
		"/emailaddress/get-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetEmailAddresses(className string, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/emailaddress/get-email-addresses": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateEmailAddress(emailAddressId int64, address string, typeId int, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["emailAddressId"] = emailAddressId
	_params["address"] = address
	_params["typeId"] = typeId
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/emailaddress/update-email-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

