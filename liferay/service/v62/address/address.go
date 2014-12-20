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

package address

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddAddress(className string, classPK int64, street1 string, street2 string, street3 string, city string, zip string, regionId int64, countryId int64, typeId int, mailing bool, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["street1"] = street1
	_params["street2"] = street2
	_params["street3"] = street3
	_params["city"] = city
	_params["zip"] = zip
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["typeId"] = typeId
	_params["mailing"] = mailing
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/address/add-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddAddress2(className string, classPK int64, street1 string, street2 string, street3 string, city string, zip string, regionId int64, countryId int64, typeId int, mailing bool, primary bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["street1"] = street1
	_params["street2"] = street2
	_params["street3"] = street3
	_params["city"] = city
	_params["zip"] = zip
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["typeId"] = typeId
	_params["mailing"] = mailing
	_params["primary"] = primary
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/address/add-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteAddress(addressId int64) error {
	_params := make(map[string]interface{})

	_params["addressId"] = addressId

	_cmd := map[string]interface{}{
		"/address/delete-address": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetAddress(addressId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["addressId"] = addressId

	_cmd := map[string]interface{}{
		"/address/get-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetAddresses(className string, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/address/get-addresses": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateAddress(addressId int64, street1 string, street2 string, street3 string, city string, zip string, regionId int64, countryId int64, typeId int, mailing bool, primary bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["addressId"] = addressId
	_params["street1"] = street1
	_params["street2"] = street2
	_params["street3"] = street3
	_params["city"] = city
	_params["zip"] = zip
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["typeId"] = typeId
	_params["mailing"] = mailing
	_params["primary"] = primary

	_cmd := map[string]interface{}{
		"/address/update-address": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

