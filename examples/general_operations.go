// Example of general operations in fsutils
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/utsav-56/go_fs_utils/fsutils"
)

func main() {
	// Create a temporary directory for our examples
	tempDir, err := os.MkdirTemp("", "fsutils-general-example")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up at the end

	fmt.Printf("Working with temporary directory: %s\n", tempDir)

	// Create test source directories and files
	sourceDir := filepath.Join(tempDir, "source")
	if err := fsutils.Mkdir(sourceDir); err != nil {
		log.Fatalf("Failed to create source directory: %v", err)
	}

	// Create a few files in the source directory
	file1 := filepath.Join(sourceDir, "file1.txt")
	file2 := filepath.Join(sourceDir, "file2.txt")
	subDir := filepath.Join(sourceDir, "subdir")

	if err := fsutils.Touch(file1); err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	if err := fsutils.Touch(file2); err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	if err := fsutils.Mkdir(subDir); err != nil {
		log.Fatalf("Failed to create subdirectory: %v", err)
	}
	if err := fsutils.Touch(filepath.Join(subDir, "subfile.txt")); err != nil {
		log.Fatalf("Failed to create file in subdirectory: %v", err)
	}

	// Example 1: Using Cp to copy a file
	fmt.Println("\n--- Copying a file using Cp ---")
	destFile := filepath.Join(tempDir, "copied-file1.txt")
	if err := fsutils.Cp(file1, destFile); err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}
	fmt.Printf("File copied successfully: %v\n", fsutils.FileExists(destFile))

	// Example 2: Using Cp to copy a directory
	fmt.Println("\n--- Copying a directory using Cp ---")
	destDir := filepath.Join(tempDir, "copied-source")
	if err := fsutils.Cp(sourceDir, destDir); err != nil {
		log.Fatalf("Failed to copy directory: %v", err)
	}
	fmt.Printf("Directory copied successfully: %v\n", fsutils.DirExists(destDir))

	// Check if subdirectories and files were copied correctly
	fmt.Printf("Subdirectory copied: %v\n", fsutils.DirExists(filepath.Join(destDir, "subdir")))
	fmt.Printf("File in subdirectory copied: %v\n", fsutils.FileExists(filepath.Join(destDir, "subdir", "subfile.txt")))

	// Example 3: Using Mv to move a file
	fmt.Println("\n--- Moving a file using Mv ---")
	movedFile := filepath.Join(tempDir, "moved-file2.txt")
	if err := fsutils.Mv(file2, movedFile); err != nil {
		log.Fatalf("Failed to move file: %v", err)
	}
	fmt.Printf("Original file exists: %v\n", fsutils.FileExists(file2))
	fmt.Printf("Moved file exists: %v\n", fsutils.FileExists(movedFile))

	// Example 4: Creating a symbolic link
	fmt.Println("\n--- Creating a symbolic link ---")
	linkPath := filepath.Join(tempDir, "link-to-file")
	if err := fsutils.Symlink(destFile, linkPath); err != nil {
		fmt.Printf("Note: Symbolic link creation failed: %v (may require admin privileges on Windows)\n", err)
	} else {
		fmt.Printf("Symbolic link created successfully\n")
	}

	fmt.Println("\nAll examples completed successfully!")
}
