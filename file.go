// Package fsutils provides filesystem utility functions for Go applications.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileExists checks if a file exists at the given path.
// Returns true if the file exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// GetFileList returns a list of all files (not directories) in the given directory.
// Returns a slice of filenames or an error if the directory cannot be read.
func GetFileList(path string) ([]string, error) {
	var filesList []string
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			filesList = append(filesList, f.Name())
		}
	}
	return filesList, nil
}

// GetFileInfo returns detailed information about a file as a map.
// Returns an empty map if the file doesn't exist or is a directory.
func GetFileInfo(path string) map[string]interface{} {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Warning: Invalid file path:", path)
		return map[string]interface{}{}
	}
	if info.IsDir() {
		fmt.Println("Warning: Path is a directory. Use mv or cp functions for folders:", path)
		return map[string]interface{}{}
	}

	absPath, _ := filepath.Abs(path)
	isExecutable := info.Mode().Perm()&0111 != 0
	ext := strings.ToLower(filepath.Ext(path))

	return map[string]interface{}{
		"name":         info.Name(),
		"absPath":      absPath,
		"ext":          ext,
		"dateCreated":  getCreatedTime(path, info),
		"dateModified": info.ModTime(),
		"isExecutable": isExecutable,
		"size":         info.Size(),
		"sizeKB":       float64(info.Size()) / 1024,
		"sizeMB":       float64(info.Size()) / (1024 * 1024),
		"sizeGB":       float64(info.Size()) / (1024 * 1024 * 1024),
		"permissions":  info.Mode().Perm().String(),
		"isHidden":     strings.HasPrefix(info.Name(), "."),
		"mode":         info.Mode().String(),
	}
}

// Touch creates an empty file at the given path if it doesn't exist,
// or updates the modification time if it does.
func Touch(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

// MoveFile moves a file from src to dst.
// Returns an error if src doesn't exist, is a directory, or if dst cannot be created.
func MoveFile(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("'%s' is a directory, use MoveDir or Mv", src)
	}
	return os.Rename(src, dst)
}

// CopyFile copies a file from src to dst.
// Returns an error if src doesn't exist, is a directory, or if dst cannot be created.
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("'%s' is a directory, use CopyDir or Cp", src)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// Symlink creates a symbolic link at linkName pointing to target.
// Returns an error if the link cannot be created.
func Symlink(target string, linkName string) error {
	return os.Symlink(target, linkName)
}

// Cp is a convenience function that copies either a file or directory from src to dst.
// It automatically determines whether to use CopyFile or CopyDir based on the src path.
func Cp(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return CopyDir(src, dst)
	}
	return CopyFile(src, dst)
}
