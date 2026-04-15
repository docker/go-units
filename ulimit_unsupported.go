// +build !linux

package units

// Verify that ulimit values work with current kernel
func (u *Ulimit) Verify() error {
	return nil
}
