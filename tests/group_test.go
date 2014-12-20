package tests

import (
	"testing"

	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/group"
)

func TestGetUserSites(t *testing.T) {
	service := group.NewService(session)

	a, err := service.GetUserSites()

	if err != nil {
		t.Fatalf("GetUserSites returned error: %v", err)
	}

	first := a[0].(map[string]interface{})

	if first["friendlyURL"] != "/test" {
		t.Fatalf("first has friendlyURL %v, want /test", first["friendlyURL"])
	}

	second := a[1].(map[string]interface{})

	if second["friendlyURL"] != "/guest" {
		t.Fatalf("second has friendlyURL %v, want /guest", second["friendlyURL"])
	}
}
