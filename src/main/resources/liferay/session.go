package liferay

import (
	"net/http"
)

type Session interface {
	Client() *http.Client
	Invoke(cmd map[string]interface{}) (interface{}, error)
	Server() string
	SetAuth(req *http.Request)
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
