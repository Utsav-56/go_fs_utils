# FSUtils - Go Filesystem Utilities

[![Go Reference](https://pkg.go.dev/badge/github.com/utsav-56/go_fs_utils.svg)](https://pkg.go.dev/github.com/utsav-56/go_fs_utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/utsav-56/go_fs_utils)](https://goreportcard.com/report/github.com/utsav-56/go_fs_utils)

FSUtils is a Go library that provides simple and intuitive filesystem utility functions. It makes common file and directory operations easier with a clean and consistent API.

## Installation

```bash
go get github.com/utsav-56/go_fs_utils
```

## Import

Import the package to use it in a clean and intuitive way:

```go
import "github.com/utsav-56/go_fs_utils/fsutils"
```

## Features

-   File operations: check, copy, move, read, and create files
-   Directory operations: check, list, copy, move, and create directories
-   Path information: detailed metadata about files and directories
-   Cross-platform support for Windows, macOS, and Linux

## Quick Examples

### File Operations

```go
package main

import (
    "fmt"
    "github.com/utsav-56/go_fs_utils/fsutils"
    "log"
)

func main() {
    // Check if a file exists
    if fsutils.FileExists("myfile.txt") {
        fmt.Println("File exists!")
    }

    // Create a new file
    err := fsutils.Touch("newfile.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Copy a file
    err = fsutils.CopyFile("source.txt", "destination.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Get detailed file information
    info := fsutils.GetFileInfo("myfile.txt")
    fmt.Printf("File size: %d bytes\n", info["size"])
    fmt.Printf("Last modified: %v\n", info["dateModified"])
}
```

### Directory Operations

```go
package main

import (
    "fmt"
    "github.com/utsav-56/go_fs_utils/fsutils"
    "log"
)

func main() {
    // Create a directory
    err := fsutils.Mkdir("mynewdir")
    if err != nil {
        log.Fatal(err)
    }

    // List directories
    dirs, err := fsutils.GetDirList("parent")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Directories:")
    for _, dir := range dirs {
        fmt.Println("-", dir)
    }

    // Copy a directory and its contents
    err = fsutils.CopyDir("sourcedir", "destdir")
    if err != nil {
        log.Fatal(err)
    }

    // Get directory information
    info := fsutils.GetDirInfo("mydir")
    fmt.Printf("Contains %d files and %d subdirectories\n",
        info["numFiles"], info["numDirs"])
}
```

## API Reference

### File Operations

-   `FileExists(path string) bool` - Check if a file exists
-   `Touch(path string) error` - Create an empty file
-   `CopyFile(src, dst string) error` - Copy a file
-   `MoveFile(src, dst string) error` - Move a file
-   `GetFileInfo(path string) map[string]interface{}` - Get detailed file information

### Directory Operations

-   `DirExists(path string) bool` - Check if a directory exists
-   `Mkdir(path string) error` - Create a directory (and parent directories if needed)
-   `GetDirList(path string) ([]string, error)` - Get a list of subdirectories
-   `GetFileList(path string) ([]string, error)` - Get a list of files in a directory
-   `GetList(path string) ([]string, error)` - Get a list of all entries in a directory
-   `CopyDir(src, dst string) error` - Copy a directory and its contents
-   `MoveDir(src, dst string) error` - Move a directory
-   `RmDir(path string) error` - Remove a directory and its contents
-   `GetDirInfo(path string) map[string]interface{}` - Get detailed directory information

### General Operations

-   `Cp(src, dst string) error` - Copy a file or directory
-   `Mv(src, dst string) error` - Move a file or directory
-   `Symlink(target, linkName string) error` - Create a symbolic link
-   `PathInfo(path string) map[string]interface{}` - Get detailed information about a file or directory

## License

This project is licensed under the MIT License - see the LICENSE file for details.
