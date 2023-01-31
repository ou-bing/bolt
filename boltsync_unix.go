//go:build !windows && !plan9 && !linux && !openbsd
// +build !windows,!plan9,!linux,!openbsd

package bolt

// fdatasync 将写入数据落盘。
func fdatasync(db *DB) error {
	return db.file.Sync()
}
