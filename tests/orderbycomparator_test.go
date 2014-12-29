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
	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/bookmarksentry"
)

func TestGetEntriesDescending(t *testing.T) {
	f := func() {
		service := bookmarksentry.NewService(session)

		orderByComparator := &liferay.ObjectWrapper{
			ClassName: "com.liferay.portlet.bookmarks.util.comparator.EntryNameComparator",
		}

		a, err := service.GetEntries2(groupId, parentFolderId, -1, -1, orderByComparator)

		if err != nil {
			t.Fatalf("GetEntries returned error: %v", err)
		}

		first := a[0].(map[string]interface{})

		if first["name"] != "Z" {
			t.Fatalf("first has name %v, want Z", first["name"])
		}

		second := a[1].(map[string]interface{})

		if second["name"] != "A" {
			t.Fatalf("second has name %v, want A", second["name"])
		}
	}

	testWithBookmarkEntries(t, f)
}

func TestNullOrderByComparator(t *testing.T) {
	f := func() {
		service := bookmarksentry.NewService(session)

		a, err := service.GetEntries2(groupId, parentFolderId, -1, -1, nil)

		if err != nil {
			t.Fatalf("GetEntries returned error: %v", err)
		}

		if len(a) != 2 {
			t.Fatalf("GetEntries returned %d elements, want 2", len(a))
		}

		first := a[0].(map[string]interface{})

		if first["name"] != "A" {
			t.Fatalf("first has name %v, want A", first["name"])
		}

		second := a[1].(map[string]interface{})

		if second["name"] != "Z" {
			t.Fatalf("second has name %v, want Z", second["name"])
		}
	}

	testWithBookmarkEntries(t, f)
}

func testWithBookmarkEntries(t *testing.T, f func()) {
	entry1 := addBookmarksEntry(t, "A", nil)
	entry2 := addBookmarksEntry(t, "Z", nil)

	f()

	defer func() {
		deleteBookmarksEntry(t, entry1)
		deleteBookmarksEntry(t, entry2)
	}()
}
