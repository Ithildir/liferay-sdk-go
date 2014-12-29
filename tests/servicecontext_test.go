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
	"math/rand"
	"strconv"
	"testing"

	"github.com/ithildir/liferay-sdk-go/liferay"
	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/bookmarksentry"
)

func TestAddBookmarksEntry(t *testing.T) {
	uuid := strconv.Itoa(rand.Int())

	serviceContext := &liferay.ObjectWrapper{
		Object: map[string]interface{}{
			"uuid":                uuid,
			"addGroupPermissions": true,
			"addGuestPermissions": true,
		},
	}

	entry := addBookmarksEntry(t, "test", serviceContext)

	if entry["uuid"] != uuid {
		t.Fatalf("entry has uuid %v, want %v", entry["uuid"], uuid)
	}

	deleteBookmarksEntry(t, entry)
}

func addBookmarksEntry(t *testing.T, name string, serviceContext *liferay.ObjectWrapper) map[string]interface{} {
	service := bookmarksentry.NewService(session)

	entry, err := service.AddEntry(groupId, parentFolderId, name, "http://www.liferay.com", "", serviceContext)

	if err != nil {
		t.Fatalf("AddEntry returned error: %v", err)
	}

	return entry
}

func deleteBookmarksEntry(t *testing.T, entry map[string]interface{}) {
	entryId := int64(entry["entryId"].(float64))

	service := bookmarksentry.NewService(session)

	if err := service.DeleteEntry(entryId); err != nil {
		t.Fatalf("DeleteEntry returned error: %v", err)
	}
}
