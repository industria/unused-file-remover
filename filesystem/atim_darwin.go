//go:build darwin

package filesystem

import (
	"syscall"
	"time"
)

func Atim(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec).UTC()
}
