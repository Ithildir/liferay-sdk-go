package tests

import (
	"testing"

	"bitbucket.org/ithildir/liferay-sdk-go/liferay"
	"bitbucket.org/ithildir/liferay-sdk-go/liferay/service/v62/dlapp"
)

const (
	folderName     string = "test"
	folderName2    string = "test2"
	parentFolderId int64  = 0
)

func TestAddFoldersBatch(t *testing.T) {
	batch := liferay.NewBatchSession(server, username, password)

	service := dlapp.NewService(batch)

	service.AddFolder(groupId, parentFolderId, folderName, "")
	service.AddFolder(groupId, parentFolderId, folderName2, "")

	array, err := batch.InvokeAll()

	if err != nil {
		t.Fatalf("InvokeAll returned error: %v", err)
	}

	folder1 := array[0].(map[string]interface{})

	if folder1["name"] != folderName {
		t.Fatalf("folder1 has name %v, want %v", folder1["name"], folderName)
	}

	folder2 := array[1].(map[string]interface{})

	if folder2["name"] != folderName2 {
		t.Fatalf("folder2 has name %v, want %v", folder1["name"], folderName2)
	}

	deleteFoldersBatch(t, batch)
}

func deleteFoldersBatch(t *testing.T, batch *liferay.BatchSession) {
	service := dlapp.NewService(batch)

	service.DeleteFolder(groupId, parentFolderId, folderName)
	service.DeleteFolder(groupId, parentFolderId, folderName2)

	a, err := batch.InvokeAll()

	if err != nil {
		t.Fatalf("InvokeAll returned error: %v", err)
	}

	if len(a) != 2 {
		t.Fatalf("InvokeAll returned %d elements, want 2", len(a))
	}
}
