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

package group

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddGroup(parentGroupId int64, liveGroupId int64, name string, description string, type int, manualMembership bool, membershipRestriction int, friendlyURL string, site bool, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentGroupId"] = parentGroupId
	_params["liveGroupId"] = liveGroupId
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["manualMembership"] = manualMembership
	_params["membershipRestriction"] = membershipRestriction
	_params["friendlyURL"] = friendlyURL
	_params["site"] = site
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/group/add-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddGroup2(name string, description string, type int, friendlyURL string, site bool, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["friendlyURL"] = friendlyURL
	_params["site"] = site
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/group/add-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddGroup3(parentGroupId int64, name string, description string, type int, friendlyURL string, site bool, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentGroupId"] = parentGroupId
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["friendlyURL"] = friendlyURL
	_params["site"] = site
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/group/add-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddRoleGroups(roleId int64, groupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/group/add-role-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CheckRemoteStagingGroup(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/check-remote-staging-group": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteGroup(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/delete-group": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DisableStaging(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/disable-staging": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) EnableStaging(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/enable-staging": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCompanyGroup(companyId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId

	_cmd := map[string]interface{}{
		"/group/get-company-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroup(groupId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/get-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroup2(companyId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/group/get-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroups(companyId int64, parentGroupId int64, site bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["parentGroupId"] = parentGroupId
	_params["site"] = site

	_cmd := map[string]interface{}{
		"/group/get-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetManageableSiteGroups(portlets map[string]interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["portlets"] = portlets
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-manageable-site-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetManageableSites(portlets map[string]interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["portlets"] = portlets
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-manageable-sites": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationsGroups(organizations []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizations"] = organizations

	_cmd := map[string]interface{}{
		"/group/get-organizations-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroup(companyId int64, userId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/group/get-user-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupsGroups(userGroups []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userGroups"] = userGroups

	_cmd := map[string]interface{}{
		"/group/get-user-groups-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserOrganizationsGroups(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/group/get-user-organizations-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserPlaces(classNames []interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNames"] = classNames
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-places": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserPlaces2(userId int64, classNames []interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["classNames"] = classNames
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-places": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserPlaces3(userId int64, classNames []interface{}, includeControlPanel bool, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["classNames"] = classNames
	_params["includeControlPanel"] = includeControlPanel
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-places": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserPlacesCount() (int, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/group/get-user-places-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserSites() ([]interface{}, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/group/get-user-sites": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserSitesGroups() ([]interface{}, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/group/get-user-sites-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserSitesGroups2(classNames []interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNames"] = classNames
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-sites-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserSitesGroups3(userId int64, classNames []interface{}, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["classNames"] = classNames
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-sites-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserSitesGroups4(userId int64, classNames []interface{}, includeControlPanel bool, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["classNames"] = classNames
	_params["includeControlPanel"] = includeControlPanel
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/group/get-user-sites-groups": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserSitesGroupsCount() (int, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/group/get-user-sites-groups-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) HasUserGroup(userId int64, groupId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/group/has-user-group": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) Search(companyId int64, name string, description string, params []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name
	_params["description"] = description
	_params["params"] = params
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/group/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, name string, description string, params []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["name"] = name
	_params["description"] = description
	_params["params"] = params

	_cmd := map[string]interface{}{
		"/group/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) SetRoleGroups(roleId int64, groupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/group/set-role-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsetRoleGroups(roleId int64, groupIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["roleId"] = roleId
	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/group/unset-role-groups": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateFriendlyUrl(groupId int64, friendlyURL string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["friendlyURL"] = friendlyURL

	_cmd := map[string]interface{}{
		"/group/update-friendly-url": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateGroup(groupId int64, parentGroupId int64, name string, description string, type int, manualMembership bool, membershipRestriction int, friendlyURL string, active bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["parentGroupId"] = parentGroupId
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["manualMembership"] = manualMembership
	_params["membershipRestriction"] = membershipRestriction
	_params["friendlyURL"] = friendlyURL
	_params["active"] = active
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/group/update-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateGroup2(groupId int64, typeSettings string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["typeSettings"] = typeSettings

	_cmd := map[string]interface{}{
		"/group/update-group": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateStagedPortlets(groupId int64, stagedPortletIds map[string]interface{}) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["stagedPortletIds"] = stagedPortletIds

	_cmd := map[string]interface{}{
		"/group/update-staged-portlets": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

