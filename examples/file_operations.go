// Example of file operations in fsutils
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
	tempDir, err := os.MkdirTemp("", "fsutils-file-example")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up at the end

	fmt.Printf("Working with temporary directory: %s\n", tempDir)

	// Example 1: Check if a file exists
	testFile := filepath.Join(tempDir, "test.txt")
	fmt.Printf("File exists before creation: %v\n", fsutils.FileExists(testFile))

	// Example 2: Create an empty file with Touch
	fmt.Println("\n--- Creating file with Touch ---")
	if err := fsutils.Touch(testFile); err != nil {
		log.Fatalf("Failed to touch file: %v", err)
	}
	fmt.Printf("File exists after Touch: %v\n", fsutils.FileExists(testFile))

	// Example 3: Get file information
	fmt.Println("\n--- Getting file information ---")
	fileInfo := fsutils.GetFileInfo(testFile)
	fmt.Printf("File name: %v\n", fileInfo["name"])
	fmt.Printf("File size: %v bytes\n", fileInfo["size"])
	fmt.Printf("File mode: %v\n", fileInfo["mode"])

	// Example 4: Create a copy of the file
	fmt.Println("\n--- Copying file ---")
	copiedFile := filepath.Join(tempDir, "copy.txt")
	if err := fsutils.CopyFile(testFile, copiedFile); err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}
	fmt.Printf("Copy exists: %v\n", fsutils.FileExists(copiedFile))

	// Example 5: Write some content to the original file
	fmt.Println("\n--- Writing content to file ---")
	content := []byte("Hello, fsutils!")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// Example 6: Move the file
	fmt.Println("\n--- Moving file ---")
	movedFile := filepath.Join(tempDir, "moved.txt")
	if err := fsutils.MoveFile(testFile, movedFile); err != nil {
		log.Fatalf("Failed to move file: %v", err)
	}
	fmt.Printf("Original file exists: %v\n", fsutils.FileExists(testFile))
	fmt.Printf("Moved file exists: %v\n", fsutils.FileExists(movedFile))

	// Example 7: Get list of files in directory
	fmt.Println("\n--- Listing files in directory ---")
	files, err := fsutils.GetFileList(tempDir)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}
	fmt.Println("Files in directory:")
	for _, file := range files {
		fmt.Printf("- %s\n", file)
	}

	fmt.Println("\nAll examples completed successfully!")
}
