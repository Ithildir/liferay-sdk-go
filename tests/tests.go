package tests

import "bitbucket.org/ithildir/liferay-sdk-go/liferay"

const (
	groupId  int64  = 10184
	password string = "test"
	server   string = "http://localhost:8080"
	username string = "test@liferay.com"
)

var session liferay.Session = liferay.NewSession(server, username, password)
