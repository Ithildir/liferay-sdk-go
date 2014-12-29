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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
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
	url := getURL(s, "/invoke")

	r, w := io.Pipe()

	errChan := make(chan error, 1)

	go func() {
		defer w.Close()

		enc := json.NewEncoder(w)

		errChan <- enc.Encode(cmds)
	}()

	req, err := http.NewRequest("POST", url, r)

	if err != nil {
		return nil, err
	}

	v, err := doRequest(s, req)

	if err != nil {
		return nil, err
	}

	return v.([]interface{}), <-errChan
}

func upload(s Session, cmd map[string]interface{}) (interface{}, error) {
	path := getPath(cmd)
	params := cmd[path].(map[string]interface{})

	url := getURL(s, path)

	r, w := io.Pipe()

	multiWriter := multipart.NewWriter(w)

	errChan := make(chan error, 1)

	go func() {
		defer w.Close()

		for k, v := range params {
			switch v := v.(type) {
			case io.Reader:
				part, err := multiWriter.CreateFormFile(k, "")

				if err != nil {
					errChan <- err

					return
				}

				io.Copy(part, v)

			default:
				if err := multiWriter.WriteField(k, fmt.Sprintf("%v", v)); err != nil {
					errChan <- err

					return
				}
			}
		}

		errChan <- multiWriter.Close()
	}()

	req, err := http.NewRequest("POST", url, r)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", multiWriter.FormDataContentType())

	v, err := doRequest(s, req)

	if err != nil {
		return nil, err
	}

	return v, <-errChan
}
