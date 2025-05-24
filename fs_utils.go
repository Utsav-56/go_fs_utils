package fsutils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

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

func Touch(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

func Mkdir(path string) error {
	return os.MkdirAll(path, 0755)
}

func RmDir(path string) error {
	return os.RemoveAll(path)
}

func Symlink(target string, linkName string) error {
	return os.Symlink(target, linkName)
}

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

func Mv(src, dst string) error {
	return os.Rename(src, dst)
}

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

func getCreatedTime(path string, info fs.FileInfo) time.Time {
	// Unix doesn't support creation time, fallback to ModTime
	return info.ModTime()
}
