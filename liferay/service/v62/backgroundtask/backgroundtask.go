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

package backgroundtask

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetBackgroundTaskStatusJson(backgroundTaskId int64) (string, error) {
	_params := make(map[string]interface{})

	_params["backgroundTaskId"] = backgroundTaskId

	_cmd := map[string]interface{}{
		"/backgroundtask/get-background-task-status-json": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetBackgroundTasksCount(groupId int64, taskExecutorClassName string, completed string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["taskExecutorClassName"] = taskExecutorClassName
	_params["completed"] = completed

	_cmd := map[string]interface{}{
		"/backgroundtask/get-background-tasks-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

