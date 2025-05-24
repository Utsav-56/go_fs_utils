/*
Package fsutils provides filesystem utility functions for Go applications.

FSUtils is a comprehensive library that makes working with files and directories
easier in Go. It provides functions for common operations such as checking
if files exist, creating directories, copying files, and getting file metadata.

# File Operations

File operations include functions for checking if files exist, copying and moving
files, and getting detailed file information:

	if fsutils.FileExists("myfile.txt") {
		// File exists
	}

	// Create an empty file
	err := fsutils.Touch("newfile.txt")

	// Copy a file
	err = fsutils.CopyFile("source.txt", "destination.txt")

	// Get file information
	info := fsutils.GetFileInfo("myfile.txt")
	fmt.Printf("File size: %d bytes\n", info["size"])

# Directory Operations

Directory operations include functions for checking if directories exist,
listing directory contents, and manipulating directories:

	// Create a directory
	err := fsutils.Mkdir("mynewdir")

	// List subdirectories
	dirs, err := fsutils.GetDirList("parent")

	// Copy a directory and its contents
	err = fsutils.CopyDir("sourcedir", "destdir")

	// Get directory information
	info := fsutils.GetDirInfo("mydir")
	fmt.Printf("Contains %d files\n", info["numFiles"])

# General Operations

General operations work on both files and directories:

	// Copy a file or directory
	err := fsutils.Cp("source", "destination")

	// Move a file or directory
	err = fsutils.Mv("source", "destination")

	// Get detailed path information
	info := fsutils.PathInfo("path")
*/
package fsutils
