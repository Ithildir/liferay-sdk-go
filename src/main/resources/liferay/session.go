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

import (
	"net/http"
)

type Session interface {
	Client() *http.Client
	Invoke(cmd map[string]interface{}) (interface{}, error)
	Server() string
	SetAuth(req *http.Request)
	Upload(cmd map[string]interface{}) (interface{}, error)
}

type session struct {
	client   *http.Client
	password string
	server   string
	username string
}

func NewSession(server string, username string, password string) Session {
	return &session{
		client:   http.DefaultClient,
		server:   server,
		password: password,
		username: username}
}

func (s *session) Client() *http.Client {
	return s.client
}

func (s *session) Invoke(cmd map[string]interface{}) (interface{}, error) {
	cmds := []map[string]interface{}{cmd}

	a, err := post(s, cmds)

	if err != nil {
		return nil, err
	}

	return a[0], nil
}

func (s *session) Server() string {
	return s.server
}

func (s *session) SetAuth(req *http.Request) {
	if len(s.username) > 0 && len(s.password) > 0 {
		req.SetBasicAuth(s.username, s.password)
	}
}

func (s *session) Upload(cmd map[string]interface{}) (interface{}, error) {
	return upload(s, cmd)
}
