package tests

import (
	"testing"

	"bitbucket.org/ithildir/liferay-sdk-go/liferay"
)

func TestPortalVersion(t *testing.T) {
	v, err := liferay.GetPortalVersion(server)

	if err != nil {
		t.Fatalf("GetPortalVersion returned error: %v", err)
	}

	if v < liferay.V62 {
		t.Fatalf("GetPortalVersion returned %d, want >= %d", v, liferay.V62)
	}
}
