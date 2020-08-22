// +build linux

package units

import (
	"fmt"

	"golang.org/x/sys/unix"
	"testing"
)

func TestParseUlimitTooBig(t *testing.T) {
	var rlimit unix.Rlimit
	if err := unix.Getrlimit(rlimitNofile, &rlimit); err != nil {
		t.Fatalf("Failed to get rlimit of current process %q", err)
	}
	u, err := ParseUlimit(fmt.Sprintf("nofile=512:%d", rlimit.Max+1))
	if err != nil {
		t.Fatalf("expected valid value got %q", err)
	}
	if err := u.Verify(); err == nil {
		t.Fatalf("expected invalid hard limit")
	}
}
