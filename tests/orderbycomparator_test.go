package tests

import (
	"testing"

	"bitbucket.org/ithildir/liferay-sdk-go/liferay"
	"bitbucket.org/ithildir/liferay-sdk-go/liferay/service/v62/bookmarksentry"
)

func TestGetEntriesDescending(t *testing.T) {
	f := func() {
		service := bookmarksentry.NewService(session)

		orderByComparator := &liferay.ObjectWrapper{
			ClassName: "com.liferay.portlet.bookmarks.util.comparator.EntryNameComparator",
		}

		entries, err := service.GetEntries(groupId, parentFolderId, -1, -1, orderByComparator)

		if err != nil {
			t.Fatalf("GetEntries returned error: %v", err)
		}

		first := entries[0]

		if first["name"] != "Z" {
			t.Fatalf("first has name %v, want Z", first["name"])
		}

		second := entries[1]

		if second["name"] != "A" {
			t.Fatalf("second has name %v, want A", second["name"])
		}
	}

	testWithBookmarkEntries(t, f)
}

func TestNullOrderByComparator(t *testing.T) {
	f := func() {
		service := bookmarksentry.NewService(session)

		entries, err := service.GetEntries(groupId, parentFolderId, -1, -1, nil)

		if err != nil {
			t.Fatalf("GetEntries returned error: %v", err)
		}

		if len(entries) != 2 {
			t.Fatalf("GetEntries returned %d elements, want 2", len(entries))
		}

		first := entries[0]

		if first["name"] != "A" {
			t.Fatalf("first has name %v, want A", first["name"])
		}

		second := entries[1]

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
