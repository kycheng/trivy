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
		if os, ok := sys.(*syscall.Dir); ok {
			// Plan 9 doesn't have a CreatedAt field.
			e.CreatedAt = time.Unix(int64(os.Mtime), 0)

			e.Dev = uint32(os.Dev)

			// Plan 9 has no Inode.
			// ext2srv(4) appears to store Inode in Qid.Path.
			e.Inode = uint32(os.Qid.Path)

			// Plan 9 has string UID/GID
			e.GID = 0
			e.UID = 0
		}
	}
}

func isSymlinkWindowsNonAdmin(err error) bool {
	return true
}
