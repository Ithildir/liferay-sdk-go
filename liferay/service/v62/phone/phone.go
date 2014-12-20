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

package phone

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddPhone(className string, classPK int64, number string, extension string, typeId int, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["number"] = number
	_params["extension"] = extension
	_params["typeId"] = typeId
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/phone/add-phone": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddPhone2(className string, classPK int64, number string, extension string, typeId int, primary bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["number"] = number
	_params["extension"] = extension
	_params["typeId"] = typeId
	_params["primary"] = primary
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/phone/add-phone": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeletePhone(phoneId int64) error {
	_params := make(map[string]interface{})

	_params["phoneId"] = phoneId

	_cmd := map[string]interface{}{
		"/phone/delete-phone": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetPhone(phoneId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["phoneId"] = phoneId

	_cmd := map[string]interface{}{
		"/phone/get-phone": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetPhones(className string, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/phone/get-phones": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdatePhone(phoneId int64, number string, extension string, typeId int, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["phoneId"] = phoneId
	_params["number"] = number
	_params["extension"] = extension
	_params["typeId"] = typeId
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/phone/update-phone": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

