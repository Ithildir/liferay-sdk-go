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

package mbthread

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) DeleteThread(threadId int64) error {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/delete-thread": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetGroupThreads(groupId int64, userId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupThreads2(groupId int64, userId int64, modifiedDate int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["modifiedDate"] = modifiedDate
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupThreads3(groupId int64, userId int64, status int, subscribed bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["subscribed"] = subscribed
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupThreads4(groupId int64, userId int64, status int, subscribed bool, includeAnonymous bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["subscribed"] = subscribed
	_params["includeAnonymous"] = includeAnonymous
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupThreadsCount(groupId int64, userId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupThreadsCount2(groupId int64, userId int64, modifiedDate int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["modifiedDate"] = modifiedDate
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupThreadsCount3(groupId int64, userId int64, status int, subscribed bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["subscribed"] = subscribed

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupThreadsCount4(groupId int64, userId int64, status int, subscribed bool, includeAnonymous bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["subscribed"] = subscribed
	_params["includeAnonymous"] = includeAnonymous

	_cmd := map[string]interface{}{
		"/mbthread/get-group-threads-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetThreads(groupId int64, categoryId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/get-threads": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetThreadsCount(groupId int64, categoryId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbthread/get-threads-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) LockThread(threadId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/lock-thread": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveThread(categoryId int64, threadId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/move-thread": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveThreadFromTrash(categoryId int64, threadId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/move-thread-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveThreadToTrash(threadId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/move-thread-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreThreadFromTrash(threadId int64) error {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/restore-thread-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Search(groupId int64, creatorUserId int64, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["creatorUserId"] = creatorUserId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search2(groupId int64, creatorUserId int64, startDate int64, endDate int64, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["creatorUserId"] = creatorUserId
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbthread/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SplitThread(messageId int64, subject string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId
	_params["subject"] = subject
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbthread/split-thread": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UnlockThread(threadId int64) error {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbthread/unlock-thread": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

