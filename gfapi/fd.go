// Copyright (c) 2013, Kaushal M <kshlmster at gmail dot com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package gfapi

// This file includes lower level operations on fd like the ones in the 'syscall' package

// #cgo pkg-config: glusterfs-api
// #include "api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import (
        "unsafe"
        "syscall"
)

// Fd is the glusterfs fd type
type Fd struct {
	fd *C.glfs_fd_t
}

// Fchmod changes the mode of the Fd to the given mode
//
// Returns error on failure
func (fd *Fd) Fchmod (mode uint32) error {
        _, err := C.glfs_fchmod(fd.fd, C.mode_t(mode))

        return err
}

// Fstat performs an fstat call on the Fd and saves stat details in the passed stat structure
//
// Returns error on failure
func (fd *Fd) Fstat(stat *syscall.Stat_t) error {
        _, err := C.glfs_fstat(fd.fd, (*C.struct_stat)(unsafe.Pointer(stat)))

        return err
}

// Fsync performs an fsync on the Fd
//
// Returns error on failure
func (fd *Fd) Fsync() error {
	_, err := C.glfs_fsync(fd.fd)

	return err
}

// Ftruncate truncates the size of the Fd to the given size
//
// Returns error on failure
func (fd *Fd) Ftruncate(size int64) error {
        _, err := C.glfs_ftruncate (fd.fd, C.off_t(size))

        return err
}

// Pread reads at most len(b) bytes into b from offset off in Fd
//
// Returns number of bytes read on success and error on failure
func (fd *Fd) Pread(b []byte, off int64) (int, error) {
	n, err := C.glfs_pread(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0)

	return int(n), err
}

// Pwrite writes len(b) bytes from b into the Fd from offset off
//
// Returns number of bytes written on success and error on failure
func (fd *Fd) Pwrite(b []byte, off int64) (int, error) {
	n, err := C.glfs_pwrite(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0)

	return int(n), err
}

// Read reads at most len(b) bytes into b from Fd
//
// Returns number of bytes read on success and error on failure
func (fd *Fd) Read(b []byte) (int, error) {
	n, err := C.glfs_read(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), 0)

	return int(n), err
}

// Write writes len(b) bytes from b into the Fd
//
// Returns number of bytes written on success and error on failure
func (fd *Fd) Write(b []byte) (int, error) {

	n, err := C.glfs_write(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), 0)
	return int(n), err

}

func (fd *Fd) lseek(offset int64, whence int) (int64, error) {
        ret, err := C.glfs_lseek(fd.fd, C.off_t(offset), C.int(whence))

        return int64(ret), err
}
