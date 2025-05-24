// Package fsutils provides filesystem utility functions for Go applications.
package fsutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DirExists checks if a directory exists at the given path.
// Returns true if the directory exists.
func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// GetDirList returns a list of all directories (not files) in the given directory.
// Returns a slice of directory names or an error if the directory cannot be read.
func GetDirList(path string) ([]string, error) {
	var dirs []string
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs, nil
}

// GetList returns a list of all entries (both files and directories) in the given directory.
// Returns a slice of names or an error if the directory cannot be read.
func GetList(path string) ([]string, error) {
	var entries []string
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		entries = append(entries, f.Name())
	}
	return entries, nil
}

// PathInfo returns detailed information about a file or directory as a map.
// Returns an empty map if the path doesn't exist.
func PathInfo(path string) map[string]interface{} {
	infoMap := make(map[string]interface{})
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Warning: Invalid path:", path)
		return map[string]interface{}{}
	}

	absPath, _ := filepath.Abs(path)
	isDir := info.IsDir()
	isFile := !isDir
	isExecutable := info.Mode().Perm()&0111 != 0
	childCount := 0
	if isDir {
		files, _ := os.ReadDir(path)
		childCount = len(files)
	}

	infoMap["name"] = info.Name()
	infoMap["absPath"] = absPath
	infoMap["isDir"] = isDir
	infoMap["isFile"] = isFile
	infoMap["isExecutable"] = isExecutable
	infoMap["size"] = info.Size()
	infoMap["sizeKb"] = float64(info.Size()) / 1024
	infoMap["sizeMB"] = float64(info.Size()) / (1024 * 1024)
	infoMap["sizeGB"] = float64(info.Size()) / (1024 * 1024 * 1024)
	infoMap["dateCreated"] = getCreatedTime(path, info)
	infoMap["dateModified"] = info.ModTime()
	infoMap["numChilds"] = childCount
	infoMap["mode"] = info.Mode().String()
	infoMap["permissions"] = info.Mode().Perm().String()
	infoMap["isHidden"] = strings.HasPrefix(info.Name(), ".")

	return infoMap
}

// GetDirInfo returns detailed information about a directory as a map.
// Returns an empty map if the path doesn't exist or is not a directory.
func GetDirInfo(path string) map[string]interface{} {
	infoMap := PathInfo(path)
	if len(infoMap) == 0 || !infoMap["isDir"].(bool) {
		fmt.Println("Warning: Not a valid directory:", path)
		return map[string]interface{}{}
	}

	numFiles := 0
	numDirs := 0
	files, err := os.ReadDir(path)
	if err != nil {
		return map[string]interface{}{}
	}
	for _, f := range files {
		if f.IsDir() {
			numDirs++
		} else {
			numFiles++
		}
	}
	infoMap["numFiles"] = numFiles
	infoMap["numDirs"] = numDirs
	return infoMap
}

// Mkdir creates a directory and all necessary parent directories at the given path.
// Uses 0755 (rwxr-xr-x) permissions by default.
func Mkdir(path string) error {
	return os.MkdirAll(path, 0755)
}

// RmDir removes a directory and all its contents recursively.
// Use with caution as this will delete all files and subdirectories.
func RmDir(path string) error {
	return os.RemoveAll(path)
}

// MoveDir moves a directory from src to dst.
// Returns an error if src doesn't exist, is not a directory, or if dst cannot be created.
func MoveDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("'%s' is not a directory, use MoveFile or Mv", src)
	}
	return os.Rename(src, dst)
}

// CopyDir copies a directory recursively from src to dst.
// It preserves the file permissions and copies all contents.
func CopyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		} else {
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			return err
		}
	})
}
