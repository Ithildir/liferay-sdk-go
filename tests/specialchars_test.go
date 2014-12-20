package tests

import (
	"testing"
)

func TestAddBookmarksEntrySpecial(t *testing.T) {
	name := "entry áéíòúñ"

	entry := addBookmarksEntry(t, name, nil)

	if entry["name"] != name {
		t.Fatalf("entry has name %v, want %v", entry["name"], name)
	}

	deleteBookmarksEntry(t, entry)
}
