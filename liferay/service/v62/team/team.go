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

package team

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddTeam(groupId int64, name string, description string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/team/add-team": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteTeam(teamId int64) error {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId

	_cmd := map[string]interface{}{
		"/team/delete-team": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetGroupTeams(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/team/get-group-teams": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTeam(teamId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId

	_cmd := map[string]interface{}{
		"/team/get-team": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTeam2(groupId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/team/get-team": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetUserTeams(userId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId

	_cmd := map[string]interface{}{
		"/team/get-user-teams": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetUserTeams2(userId int64, groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/team/get-user-teams": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) HasUserTeam(userId int64, teamId int64) (bool, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["teamId"] = teamId

	_cmd := map[string]interface{}{
		"/team/has-user-team": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

func (s *Service) UpdateTeam(teamId int64, name string, description string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["teamId"] = teamId
	_params["name"] = name
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/team/update-team": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

