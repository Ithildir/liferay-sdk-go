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

package socialactivity

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetActivities(className string, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetActivities2(classNameId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNameId"] = classNameId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetActivities3(mirrorActivityId int64, className string, classPK int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["mirrorActivityId"] = mirrorActivityId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetActivities4(mirrorActivityId int64, classNameId int64, classPK int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["mirrorActivityId"] = mirrorActivityId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetActivitiesCount(className string) (int, error) {
	_params := make(map[string]interface{})

	_params["className"] = className

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetActivitiesCount2(classNameId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["classNameId"] = classNameId

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetActivitiesCount3(mirrorActivityId int64, className string, classPK int64) (int, error) {
	_params := make(map[string]interface{})

	_params["mirrorActivityId"] = mirrorActivityId
	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetActivitiesCount4(mirrorActivityId int64, classNameId int64, classPK int64) (int, error) {
	_params := make(map[string]interface{})

	_params["mirrorActivityId"] = mirrorActivityId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/socialactivity/get-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetActivity(activityId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["activityId"] = activityId

	_cmd := map[string]interface{}{
		"/socialactivity/get-activity": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetActivitySetActivities(activitySetId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["activitySetId"] = activitySetId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-activity-set-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupActivities(groupId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-group-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupActivitiesCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/socialactivity/get-group-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupUsersActivities(groupId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-group-users-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupUsersActivitiesCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/socialactivity/get-group-users-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetMirrorActivity(mirrorActivityId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["mirrorActivityId"] = mirrorActivityId

	_cmd := map[string]interface{}{
		"/socialactivity/get-mirror-activity": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationActivities(organizationId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-organization-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationActivitiesCount(organizationId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/socialactivity/get-organization-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetOrganizationUsersActivities(organizationId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-organization-users-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetOrganizationUsersActivitiesCount(organizationId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["organizationId"] = organizationId

	_cmd := map[string]interface{}{
		"/socialactivity/get-organization-users-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetRelationActivities(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-relation-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRelationActivities2(userId int64, type int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["type"] = type
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-relation-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRelationActivitiesCount(userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/socialactivity/get-relation-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetRelationActivitiesCount2(userId int64, type int) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["type"] = type

	_cmd := map[string]interface{}{
		"/socialactivity/get-relation-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserActivities(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserActivitiesCount(userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserGroupsActivities(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-groups-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupsActivitiesCount(userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-groups-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserGroupsAndOrganizationsActivities(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-groups-and-organizations-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserGroupsAndOrganizationsActivitiesCount(userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-groups-and-organizations-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetUserOrganizationsActivities(userId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-organizations-activities": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserOrganizationsActivitiesCount(userId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/socialactivity/get-user-organizations-activities-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

