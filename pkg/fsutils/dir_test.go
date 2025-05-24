package fsutils_test

import (
	"os"
	"path/filepath"
	"testing"
	
	"github.com/go-fs-utils/fsutils"
)

func TestDirectoryOperations(t *testing.T) {
	// Create a temporary directory for our tests
	tempDir, err := os.MkdirTemp("", "fsutils-dir-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Test Mkdir
	testSubdir := filepath.Join(tempDir, "subdir")
	if err := fsutils.Mkdir(testSubdir); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	
	// Test DirExists
	if !fsutils.DirExists(testSubdir) {
		t.Errorf("DirExists failed to detect directory at %s", testSubdir)
	}
	
	// Create some files in the directory
	for i := 1; i <= 3; i++ {
		testFile := filepath.Join(testSubdir, "file-"+string(i))
		if err := fsutils.Touch(testFile); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}
	
	// Test CopyDir
	copyDest := filepath.Join(tempDir, "subdir-copy")
	if err := fsutils.CopyDir(testSubdir, copyDest); err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}
	
	if !fsutils.DirExists(copyDest) {
		t.Errorf("CopyDir didn't create destination directory at %s", copyDest)
	}
	
	// Test GetDirList
	subdirs, err := fsutils.GetDirList(tempDir)
	if err != nil {
		t.Fatalf("GetDirList failed: %v", err)
	}
	
	// We should have exactly 2 directories
	if len(subdirs) != 2 {
		t.Errorf("GetDirList returned wrong number of directories: got %d, want %d", len(subdirs), 2)
	}
	
	// Test GetDirInfo
	info := fsutils.GetDirInfo(testSubdir)
	if info["numFiles"] != 3 {
		t.Errorf("GetDirInfo returned wrong number of files: got %v, want %v", info["numFiles"], 3)
	}
}
