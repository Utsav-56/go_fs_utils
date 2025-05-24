// Example of directory operations in fsutils
package main

import (
	"fmt"
	"github.com/go-fs-utils/fsutils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a temporary directory for our examples
	tempDir, err := os.MkdirTemp("", "fsutils-dir-example")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up at the end

	fmt.Printf("Working with temporary directory: %s\n", tempDir)

	// Example 1: Check if a directory exists
	testDir := filepath.Join(tempDir, "testdir")
	fmt.Printf("Directory exists before creation: %v\n", fsutils.DirExists(testDir))

	// Example 2: Create a directory
	fmt.Println("\n--- Creating directory ---")
	if err := fsutils.Mkdir(testDir); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	fmt.Printf("Directory exists after creation: %v\n", fsutils.DirExists(testDir))

	// Create some test files and subdirectories for later examples
	subDir := filepath.Join(testDir, "subdir")
	if err := fsutils.Mkdir(subDir); err != nil {
		log.Fatalf("Failed to create subdirectory: %v", err)
	}

	testFile1 := filepath.Join(testDir, "file1.txt")
	testFile2 := filepath.Join(testDir, "file2.txt")
	if err := fsutils.Touch(testFile1); err != nil {
		log.Fatalf("Failed to create test file: %v", err)
	}
	if err := fsutils.Touch(testFile2); err != nil {
		log.Fatalf("Failed to create test file: %v", err)
	}

	// Example 3: Get directory information
	fmt.Println("\n--- Getting directory information ---")
	dirInfo := fsutils.GetDirInfo(testDir)
	fmt.Printf("Directory name: %v\n", dirInfo["name"])
	fmt.Printf("Number of files: %v\n", dirInfo["numFiles"])
	fmt.Printf("Number of subdirectories: %v\n", dirInfo["numDirs"])

	// Example 4: List directories
	fmt.Println("\n--- Listing subdirectories ---")
	dirs, err := fsutils.GetDirList(testDir)
	if err != nil {
		log.Fatalf("Failed to list directories: %v", err)
	}
	fmt.Println("Subdirectories:")
	for _, dir := range dirs {
		fmt.Printf("- %s\n", dir)
	}

	// Example 5: List all entries (files and directories)
	fmt.Println("\n--- Listing all entries ---")
	entries, err := fsutils.GetList(testDir)
	if err != nil {
		log.Fatalf("Failed to list entries: %v", err)
	}
	fmt.Println("All entries:")
	for _, entry := range entries {
		fmt.Printf("- %s\n", entry)
	}

	// Example 6: Copy a directory
	fmt.Println("\n--- Copying directory ---")
	copiedDir := filepath.Join(tempDir, "copied-dir")
	if err := fsutils.CopyDir(testDir, copiedDir); err != nil {
		log.Fatalf("Failed to copy directory: %v", err)
	}
	fmt.Printf("Copied directory exists: %v\n", fsutils.DirExists(copiedDir))

	// Example 7: Move a directory
	fmt.Println("\n--- Moving directory ---")
	movedDir := filepath.Join(tempDir, "moved-dir")
	if err := fsutils.MoveDir(copiedDir, movedDir); err != nil {
		log.Fatalf("Failed to move directory: %v", err)
	}
	fmt.Printf("Original directory exists: %v\n", fsutils.DirExists(copiedDir))
	fmt.Printf("Moved directory exists: %v\n", fsutils.DirExists(movedDir))

	// Example 8: General path information
	fmt.Println("\n--- General path information ---")
	pathInfo := fsutils.PathInfo(testDir)
	fmt.Printf("Path: %v\n", pathInfo["absPath"])
	fmt.Printf("Is directory: %v\n", pathInfo["isDir"])
	fmt.Printf("Is file: %v\n", pathInfo["isFile"])
	fmt.Printf("Permission: %v\n", pathInfo["permissions"])

	fmt.Println("\nAll examples completed successfully!")
}
