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
	"strings"
	"testing"

	"github.com/ithildir/liferay-sdk-go/liferay"
	"github.com/ithildir/liferay-sdk-go/liferay/service/v62/dlapp"
)

const (
	folderName     string = "test"
	folderName2    string = "test2"
	mimeType       string = "text/plain"
	parentFolderId int64  = 0
	sourceFileName string = "test.properties"
)

func TestAddFileEntryBytes(t *testing.T) {
	service := dlapp.NewService(session)

	b := []byte("Hello")

	entry, err := service.AddFileEntry(groupId, parentFolderId, sourceFileName, mimeType, sourceFileName, "", "", b, nil)

	if err != nil {
		t.Fatalf("AddFileEntry returned error: %v", err)
	}

	if entry["title"] != sourceFileName {
		t.Fatalf("entry has title %v, want %v", entry["title"], sourceFileName)
	}

	if err := service.DeleteFileEntry(int64(entry["fileEntryId"].(float64))); err != nil {
		t.Fatalf("DeleteFileEntry returned error: %v", err)
	}
}

func TestAddFileEntryReader(t *testing.T) {
	service := dlapp.NewService(session)

	r := strings.NewReader("Hello")

	entry, err := service.AddFileEntry2(groupId, parentFolderId, sourceFileName, mimeType, sourceFileName, "", "", r, nil)

	if err != nil {
		t.Fatalf("AddFileEntry2 returned error: %v", err)
	}

	if entry["title"] != sourceFileName {
		t.Fatalf("entry has title %v, want %v", entry["title"], sourceFileName)
	}

	if err := service.DeleteFileEntry(int64(entry["fileEntryId"].(float64))); err != nil {
		t.Fatalf("DeleteFileEntry returned error: %v", err)
	}
}

func TestAddFoldersBatch(t *testing.T) {
	batch := liferay.NewBatchSession(server, username, password)

	service := dlapp.NewService(batch)

	service.AddFolder(groupId, parentFolderId, folderName, "", nil)
	service.AddFolder(groupId, parentFolderId, folderName2, "", nil)

	a, err := batch.InvokeAll()

	if err != nil {
		t.Fatalf("InvokeAll returned error: %v", err)
	}

	folder1 := a[0].(map[string]interface{})

	if folder1["name"] != folderName {
		t.Fatalf("folder1 has name %v, want %v", folder1["name"], folderName)
	}

	folder2 := a[1].(map[string]interface{})

	if folder2["name"] != folderName2 {
		t.Fatalf("folder2 has name %v, want %v", folder1["name"], folderName2)
	}

	deleteFoldersBatch(t, batch)
}

func deleteFoldersBatch(t *testing.T, batch *liferay.BatchSession) {
	service := dlapp.NewService(batch)

	service.DeleteFolder2(groupId, parentFolderId, folderName)
	service.DeleteFolder2(groupId, parentFolderId, folderName2)

	a, err := batch.InvokeAll()

	if err != nil {
		t.Fatalf("InvokeAll returned error: %v", err)
	}

	if len(a) != 2 {
		t.Fatalf("InvokeAll returned %d elements, want 2", len(a))
	}
}
