//go:build linux
// +build linux

package git

import (
	"syscall"
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	fillSystemInfo = func(e *index.Entry, sys interface{}) {
		if os, ok := sys.(*syscall.Stat_t); ok {
			e.CreatedAt = time.Unix(int64(os.Ctim.Sec), int64(os.Ctim.Nsec))
			e.Dev = uint32(os.Dev)
			e.Inode = uint32(os.Ino)
			e.GID = os.Gid
			e.UID = os.Uid
		}
	}
}

func isSymlinkWindowsNonAdmin(err error) bool {
	return false
}
