//go:build linux

package filesystem

import (
	"syscall"
	"time"
)

func Atim(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec).UTC()
}
