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

package liferay

import "errors"

type BatchSession struct {
	cmds []map[string]interface{}
	*session
}

func NewBatchSession(server string, username string, password string) *BatchSession {
	return &BatchSession{session: NewSession(server, username, password).(*session)}
}

func (s *BatchSession) Clear() {
	s.cmds = nil
}

func (s *BatchSession) Invoke(cmd map[string]interface{}) (interface{}, error) {
	s.cmds = append(s.cmds, cmd)

	return nil, nil
}

func (s *BatchSession) InvokeAll() ([]interface{}, error) {
	defer s.Clear()

	return post(s, s.cmds)
}

func (s *BatchSession) Upload(cmd map[string]interface{}) (interface{}, error) {
	return nil, errors.New("Can't batch upload requests")
}
