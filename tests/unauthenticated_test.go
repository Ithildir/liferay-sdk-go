package tests

import (
	"testing"

	"github.com/ithildir/liferay-sdk-go/liferay"
	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/group"
)

func TestUnauthenticatedService(t *testing.T) {
	unauthenticated := liferay.NewSession(server, "", "")

	service := group.NewService(unauthenticated)

	_, err := service.GetUserSites()

	if err == nil {
		t.FailNow()
	}

	if err.Error() != "Authenticated access required" {
		t.Fatalf("GetUserSites returned error: %v", err)
	}
}
