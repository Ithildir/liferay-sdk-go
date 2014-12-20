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

package company

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) DeleteLogo(companyId int64) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/company/delete-logo": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCompanyById(companyId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/company/get-company-by-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyByLogoId(logoId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["logoId"] = logoId

	_cmd := map[string]interface{}{
		"/company/get-company-by-logo-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyByMx(mx string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["mx"] = mx

	_cmd := map[string]interface{}{
		"/company/get-company-by-mx": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyByVirtualHost(virtualHost string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["virtualHost"] = virtualHost

	_cmd := map[string]interface{}{
		"/company/get-company-by-virtual-host": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetCompanyByWebId(webId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["webId"] = webId

	_cmd := map[string]interface{}{
		"/company/get-company-by-web-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateCompany(companyId int64, virtualHost string, mx string, homeURL string, name string, legalName string, legalId string, legalType string, sicCode string, tickerSymbol string, industry string, type string, size string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["virtualHost"] = virtualHost
	_params["mx"] = mx
	_params["homeURL"] = homeURL
	_params["name"] = name
	_params["legalName"] = legalName
	_params["legalId"] = legalId
	_params["legalType"] = legalType
	_params["sicCode"] = sicCode
	_params["tickerSymbol"] = tickerSymbol
	_params["industry"] = industry
	_params["type"] = type
	_params["size"] = size

	_cmd := map[string]interface{}{
		"/company/update-company": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateCompany2(companyId int64, virtualHost string, mx string, maxUsers int, active bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["virtualHost"] = virtualHost
	_params["mx"] = mx
	_params["maxUsers"] = maxUsers
	_params["active"] = active

	_cmd := map[string]interface{}{
		"/company/update-company": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateDisplay(companyId int64, languageId string, timeZoneId string) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["languageId"] = languageId
	_params["timeZoneId"] = timeZoneId

	_cmd := map[string]interface{}{
		"/company/update-display": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateLogo(companyId int64, bytes []byte) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["bytes"] = bytes

	_cmd := map[string]interface{}{
		"/company/update-logo": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

