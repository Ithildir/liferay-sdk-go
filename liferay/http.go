package liferay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var ErrUnauthorized = errors.New("http: authentication failed")

type ServerError int

func (e ServerError) Error() string {
	return fmt.Sprintf("http: request failed, status code %d", int(e))
}

const jsonwsPath string = "api/jsonws"

func ToJSONString(a []byte) string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	for i, b := range a {
		buffer.WriteString(strconv.Itoa(int(b)))

		if i < (len(a) - 1) {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}

func doRequest(s Session, req *http.Request) (interface{}, error) {
	s.SetAuth(req)

	client := s.Client()

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	var v interface{}

	if err := dec.Decode(&v); err != nil {
		return nil, err
	}

	if err := handleServerException(res, v); err != nil {
		return nil, err
	}

	return v, nil
}

func getPath(cmd map[string]interface{}) string {
	for k, _ := range cmd {
		return k
	}

	return ""
}

func getURL(s Session, path string) string {
	url := s.Server()

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	url += jsonwsPath
	url += path

	return url
}

func handleServerException(res *http.Response, v interface{}) error {
	c := res.StatusCode

	if c == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if c != http.StatusOK {
		return ServerError(c)
	}

	m, ok := v.(map[string]interface{})

	if ok {
		msg, ok := m["exception"]

		if ok {
			return errors.New(msg.(string))
		}
	}

	return nil
}

func post(s Session, cmds []map[string]interface{}) ([]interface{}, error) {
	b, err := json.Marshal(cmds)

	if err != nil {
		return nil, err
	}

	url := getURL(s, "/invoke")

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	v, err := doRequest(s, req)

	if err != nil {
		return nil, err
	}

	return v.([]interface{}), nil
}

func upload(s Session, cmd map[string]interface{}) (interface{}, error) {
	return nil, nil
}
