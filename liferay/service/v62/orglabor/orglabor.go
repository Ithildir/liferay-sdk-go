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

package orglabor

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddOrgLabor(organizationId int64, typeId int, sunOpen int, sunClose int, monOpen int, monClose int, tueOpen int, tueClose int, wedOpen int, wedClose int, thuOpen int, thuClose int, friOpen int, friClose int, satOpen int, satClose int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["typeId"] = typeId
	_params["sunOpen"] = sunOpen
	_params["sunClose"] = sunClose
	_params["monOpen"] = monOpen
	_params["monClose"] = monClose
	_params["tueOpen"] = tueOpen
	_params["tueClose"] = tueClose
	_params["wedOpen"] = wedOpen
	_params["wedClose"] = wedClose
	_params["thuOpen"] = thuOpen
	_params["thuClose"] = thuClose
	_params["friOpen"] = friOpen
	_params["friClose"] = friClose
	_params["satOpen"] = satOpen
	_params["satClose"] = satClose

	_cmd := map[string]interface{}{
		"/orglabor/add-org-labor": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteOrgLabor(orgLaborId int64) error {
	_params := make(map[string]interface{})

	_params["orgLaborId"] = orgLaborId

	_cmd := map[string]interface{}{
		"/orglabor/delete-org-labor": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetOrgLabor(orgLaborId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["orgLaborId"] = orgLaborId

	_cmd := map[string]interface{}{
		"/orglabor/get-org-labor": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetOrgLabors(organizationId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/orglabor/get-org-labors": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrgLabor(orgLaborId int64, typeId int, sunOpen int, sunClose int, monOpen int, monClose int, tueOpen int, tueClose int, wedOpen int, wedClose int, thuOpen int, thuClose int, friOpen int, friClose int, satOpen int, satClose int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["orgLaborId"] = orgLaborId
	_params["typeId"] = typeId
	_params["sunOpen"] = sunOpen
	_params["sunClose"] = sunClose
	_params["monOpen"] = monOpen
	_params["monClose"] = monClose
	_params["tueOpen"] = tueOpen
	_params["tueClose"] = tueClose
	_params["wedOpen"] = wedOpen
	_params["wedClose"] = wedClose
	_params["thuOpen"] = thuOpen
	_params["thuClose"] = thuClose
	_params["friOpen"] = friOpen
	_params["friClose"] = friClose
	_params["satOpen"] = satOpen
	_params["satClose"] = satClose

	_cmd := map[string]interface{}{
		"/orglabor/update-org-labor": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

