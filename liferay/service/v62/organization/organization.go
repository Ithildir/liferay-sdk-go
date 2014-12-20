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

package organization

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddGroupOrganizations(groupId int64, organizationIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["organizationIds"] = organizationIds

	_cmd := map[string]interface{}{
		"/organization/add-group-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddOrganization(parentOrganizationId int64, name string, _type string, recursable bool, regionId int64, countryId int64, statusId int, comments string, site bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["recursable"] = recursable
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/add-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddOrganization2(parentOrganizationId int64, name string, _type string, regionId int64, countryId int64, statusId int, comments string, site bool, addresses []interface{}, emailAddresses []interface{}, orgLabors []interface{}, phones []interface{}, websites []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["orgLabors"] = orgLabors
	_params["phones"] = phones
	_params["websites"] = websites
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/add-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddOrganization3(parentOrganizationId int64, name string, _type string, recursable bool, regionId int64, countryId int64, statusId int, comments string, site bool, addresses []interface{}, emailAddresses []interface{}, orgLabors []interface{}, phones []interface{}, websites []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["recursable"] = recursable
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["orgLabors"] = orgLabors
	_params["phones"] = phones
	_params["websites"] = websites
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/add-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddOrganization4(parentOrganizationId int64, name string, _type string, regionId int64, countryId int64, statusId int, comments string, site bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/add-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddPasswordPolicyOrganizations(passwordPolicyId int64, organizationIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["organizationIds"] = organizationIds

	_cmd := map[string]interface{}{
		"/organization/add-password-policy-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteLogo(organizationId int64) error {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/organization/delete-logo": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteOrganization(organizationId int64) error {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/organization/delete-organization": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetManageableOrganizations(actionId string, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["actionId"] = actionId
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/organization/get-manageable-organizations": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganization(organizationId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/organization/get-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationId(companyId int64, name string) (int64, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/organization/get-organization-id": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) GetOrganizations(companyId int64, parentOrganizationId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["parentOrganizationId"] = parentOrganizationId

	_cmd := map[string]interface{}{
		"/organization/get-organizations": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizations2(companyId int64, parentOrganizationId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["parentOrganizationId"] = parentOrganizationId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/organization/get-organizations": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationsCount(companyId int64, parentOrganizationId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["parentOrganizationId"] = parentOrganizationId

	_cmd := map[string]interface{}{
		"/organization/get-organizations-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserOrganizations(userId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/organization/get-user-organizations": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SetGroupOrganizations(groupId int64, organizationIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["organizationIds"] = organizationIds

	_cmd := map[string]interface{}{
		"/organization/set-group-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetGroupOrganizations(groupId int64, organizationIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["organizationIds"] = organizationIds

	_cmd := map[string]interface{}{
		"/organization/unset-group-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetPasswordPolicyOrganizations(passwordPolicyId int64, organizationIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["organizationIds"] = organizationIds

	_cmd := map[string]interface{}{
		"/organization/unset-password-policy-organizations": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateOrganization(organizationId int64, parentOrganizationId int64, name string, _type string, regionId int64, countryId int64, statusId int, comments string, site bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/update-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrganization2(organizationId int64, parentOrganizationId int64, name string, _type string, recursable bool, regionId int64, countryId int64, statusId int, comments string, site bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["recursable"] = recursable
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/update-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrganization3(organizationId int64, parentOrganizationId int64, name string, _type string, regionId int64, countryId int64, statusId int, comments string, site bool, addresses []interface{}, emailAddresses []interface{}, orgLabors []interface{}, phones []interface{}, websites []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["orgLabors"] = orgLabors
	_params["phones"] = phones
	_params["websites"] = websites
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/update-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrganization4(organizationId int64, parentOrganizationId int64, name string, _type string, recursable bool, regionId int64, countryId int64, statusId int, comments string, site bool, addresses []interface{}, emailAddresses []interface{}, orgLabors []interface{}, phones []interface{}, websites []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["parentOrganizationId"] = parentOrganizationId
	_params["name"] = name
	_params["type"] = _type
	_params["recursable"] = recursable
	_params["regionId"] = regionId
	_params["countryId"] = countryId
	_params["statusId"] = statusId
	_params["comments"] = comments
	_params["site"] = site
	_params["addresses"] = addresses
	_params["emailAddresses"] = emailAddresses
	_params["orgLabors"] = orgLabors
	_params["phones"] = phones
	_params["websites"] = websites
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/organization/update-organization": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

