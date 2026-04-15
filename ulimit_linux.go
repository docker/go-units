package units

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// Verify that ulimit values work with current kernel
func (u *Ulimit) Verify() error {
	var rlimit unix.Rlimit
	if err := unix.Getrlimit(ulimitNameMapping[u.Name], &rlimit); err == nil {
		if u.Hard > int64(rlimit.Max) {
			return fmt.Errorf("ulimit hard limit (%d) must be less than or equal to hard limit for the current process %d", u.Hard, rlimit.Max)
		}
	}
	return nil
}
