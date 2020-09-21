// +build !glusterfs_legacy_api

package gfapi

// #cgo pkg-config: glusterfs-api
// #include "glusterfs/api/glfs.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// Fsync performs an fsync on the Fd
//
// Returns error on failure
func (fd *Fd) Fsync() error {
	ret, err := C.glfs_fsync(fd.fd, nil, nil)
	if ret < 0 {
		return err
	}
	return nil
}

// Ftruncate truncates the size of the Fd to the given size
//
// Returns error on failure
func (fd *Fd) Ftruncate(size int64) error {
	_, err := C.glfs_ftruncate(fd.fd, C.off_t(size), nil, nil)

	return err
}

// Pread reads at most len(b) bytes into b from offset off in Fd
//
// Returns number of bytes read on success and error on failure
func (fd *Fd) Pread(b []byte, off int64) (int, error) {
	n, err := C.glfs_pread(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0, nil)

	return int(n), err
}

// Pwrite writes len(b) bytes from b into the Fd from offset off
//
// Returns number of bytes written on success and error on failure
func (fd *Fd) Pwrite(b []byte, off int64) (int, error) {
	n, err := C.glfs_pwrite(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0, nil, nil)

	return int(n), err
}
