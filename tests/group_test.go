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
