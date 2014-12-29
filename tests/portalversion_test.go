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

	"github.com/ithildir/liferay-sdk-go/liferay"
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
