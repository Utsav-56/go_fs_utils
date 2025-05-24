package fsutils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-fs-utils/fsutils"
)

func TestFileOperations(t *testing.T) {
	// Create a temporary directory for our tests
	tempDir, err := os.MkdirTemp("", "fsutils-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test file creation
	testFile := filepath.Join(tempDir, "test.txt")
	if err := fsutils.Touch(testFile); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test FileExists
	if !fsutils.FileExists(testFile) {
		t.Errorf("FileExists failed to detect file at %s", testFile)
	}

	// Test CopyFile
	copyDest := filepath.Join(tempDir, "test-copy.txt")
	if err := fsutils.CopyFile(testFile, copyDest); err != nil {
		t.Fatalf("Failed to copy file: %v", err)
	}

	if !fsutils.FileExists(copyDest) {
		t.Errorf("CopyFile didn't create destination file at %s", copyDest)
	}

	// Test MoveFile
	moveDest := filepath.Join(tempDir, "test-move.txt")
	if err := fsutils.MoveFile(copyDest, moveDest); err != nil {
		t.Fatalf("Failed to move file: %v", err)
	}

	if fsutils.FileExists(copyDest) {
		t.Errorf("MoveFile didn't remove source file at %s", copyDest)
	}

	if !fsutils.FileExists(moveDest) {
		t.Errorf("MoveFile didn't create destination file at %s", moveDest)
	}

	// Test GetFileInfo
	info := fsutils.GetFileInfo(testFile)
	if info["name"] != "test.txt" {
		t.Errorf("GetFileInfo returned wrong name: got %v, want %v", info["name"], "test.txt")
	}
}
