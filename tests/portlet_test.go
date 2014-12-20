package tests

import (
	"testing"

	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/portlet"
)

func TestXYZ(t *testing.T) {
	service := portlet.NewService(session)

	a, err := service.GetWarPortlets()

	if err != nil {
		t.Fatalf("GetWarPortlets returned error: %v", err)
	}

	if a == nil || len(a) == 0 {
		t.Fatal("GetWarPortlets returned 0 elements, want > 0")
	}
}
