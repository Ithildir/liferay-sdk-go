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

package layoutsetbranch

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLayoutSetBranch(groupId int64, privateLayout bool, name string, description string, master bool, copyLayoutSetBranchId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["name"] = name
	_params["description"] = description
	_params["master"] = master
	_params["copyLayoutSetBranchId"] = copyLayoutSetBranchId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutsetbranch/add-layout-set-branch": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteLayoutSetBranch(layoutSetBranchId int64) error {
	_params := make(map[string]interface{})

	_params["layoutSetBranchId"] = layoutSetBranchId

	_cmd := map[string]interface{}{
		"/layoutsetbranch/delete-layout-set-branch": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetLayoutSetBranches(groupId int64, privateLayout bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout

	_cmd := map[string]interface{}{
		"/layoutsetbranch/get-layout-set-branches": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) MergeLayoutSetBranch(layoutSetBranchId int64, mergeLayoutSetBranchId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["layoutSetBranchId"] = layoutSetBranchId
	_params["mergeLayoutSetBranchId"] = mergeLayoutSetBranchId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutsetbranch/merge-layout-set-branch": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayoutSetBranch(groupId int64, layoutSetBranchId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["layoutSetBranchId"] = layoutSetBranchId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutsetbranch/update-layout-set-branch": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

