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

package wikinode

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddNode(name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikinode/add-node": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteNode(nodeId int64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/delete-node": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetNode(nodeId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/get-node": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetNode2(groupId int64, name string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/wikinode/get-node": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetNodes(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetNodes2(groupId int64, status int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetNodes3(groupId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetNodes4(groupId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetNodesCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetNodesCount2(groupId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/wikinode/get-nodes-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) MoveNodeToTrash(nodeId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/move-node-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreNodeFromTrash(nodeId int64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/restore-node-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SubscribeNode(nodeId int64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/subscribe-node": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeNode(nodeId int64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikinode/unsubscribe-node": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateNode(nodeId int64, name string, description string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["name"] = name
	_params["description"] = description
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikinode/update-node": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

