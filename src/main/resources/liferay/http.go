package liferay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

var ErrUnauthorized = errors.New("http: authentication failed")

type ServerError int

func (e ServerError) Error() string {
	return fmt.Sprintf("http: request failed, status code %d", int(e))
}

const jsonwsPath string = "api/jsonws"

func getPath(cmd map[string]interface{}) string {
	for k, _ := range cmd {
		return k
	}

	return ""
}

func getURL(session Session, path string) string {
	url := session.Server()

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	url += jsonwsPath
	url += path

	return url
}

func handleServerException(res *http.Response, b []byte) error {
	c := res.StatusCode

	if c == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if c != http.StatusOK {
		return ServerError(c)
	}

	var o map[string]interface{}

	if err := json.Unmarshal(b, &o); err == nil {
		msg, ok := o["exception"]

		if ok {
			return errors.New(msg.(string))
		}
	}

	return nil
}

func post(session Session, cmds []map[string]interface{}) ([]interface{}, error) {
	b, err := json.Marshal(cmds)

	if err != nil {
		return nil, err
	}

	url := getURL(session, "/invoke")

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	session.SetAuth(req)

	client := session.Client()

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := handleServerException(res, b); err != nil {
		return nil, err
	}

	var a []interface{}

	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}

	return a, nil
}

func upload(session Session, cmd map[string]interface{}) (map[string]interface{}, error) {
	path := getPath(cmd)
	params := cmd[path].(map[string]interface{})

	var body bytes.Buffer

	w := multipart.NewWriter(&body)

	w.Close()

	url := getURL(session, path)

	req, err := http.NewRequest("POST", url, &body)

	if err != nil {
		return nil, err
	}

	session.SetAuth(req)

	client := session.Client()

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := handleServerException(res, b); err != nil {
		return nil, err
	}

	var a []interface{}

	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}

	return a, nil
}
