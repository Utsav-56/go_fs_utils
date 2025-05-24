// Package fsutils provides filesystem utility functions for Go applications.
package fsutils

import (
	"io/fs"
	"time"
)

// getCreatedTime returns the creation time of a file or directory.
// Since creation time is not available on all platforms (particularly Unix),
// this function falls back to modification time.
func getCreatedTime(path string, info fs.FileInfo) time.Time {
	// Unix doesn't support creation time, fallback to ModTime
	return info.ModTime()
}
