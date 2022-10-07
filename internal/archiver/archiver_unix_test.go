//go:build !windows
// +build !windows

package archiver

import (
	"os"
	"syscall"
	"testing"
)

type wrappedFileInfo struct {
	os.FileInfo
	sys  any
	mode os.FileMode
}

func (fi wrappedFileInfo) Sys() any {
	return fi.sys
}

func (fi wrappedFileInfo) Mode() os.FileMode {
	return fi.mode
}

// wrapFileInfo returns a new os.FileInfo with the mode, owner, and group fields changed.
func wrapFileInfo(t testing.TB, fi os.FileInfo) os.FileInfo {
	// get the underlying stat_t and modify the values
	stat := fi.Sys().(*syscall.Stat_t)
	stat.Mode = mockFileInfoMode
	stat.Uid = mockFileInfoUID
	stat.Gid = mockFileInfoGID

	// wrap the os.FileInfo so we can return a modified stat_t
	res := wrappedFileInfo{
		FileInfo: fi,
		sys:      stat,
		mode:     mockFileInfoMode,
	}

	return res
}
