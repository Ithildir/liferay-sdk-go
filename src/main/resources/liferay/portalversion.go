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
