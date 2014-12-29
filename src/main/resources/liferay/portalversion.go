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
	"strconv"
	"strings"
)

const (
	UnknownVersion int = -1
	V62            int = 6200
)

func GetPortalVersion(url string) (int, error) {
	res, err := http.Head(url)

	if err != nil {
		return UnknownVersion, err
	}

	defer res.Body.Close()

	h := res.Header.Get("Liferay-Portal")

	if len(h) == 0 {
		return UnknownVersion, nil
	}

	i := strings.Index(h, "Build")

	if i == -1 {
		return UnknownVersion, nil
	}

	return strconv.Atoi(h[(i + 6):(i + 10)])
}
