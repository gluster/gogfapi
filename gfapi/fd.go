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

// Package gfapi provides a wrapper around gfapi, the GlusterFS api, which is used to access files/directories on a Gluster volume.
// The design tries to follow the default go file handling functions provided by the os package as much as possible.
package gfapi

// This file includes lower level operations on fd like the ones in the 'syscall' package

// #cgo pkg-config: glusterfs-api
// #include "api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import "unsafe"

type Fd struct {
	fd *C.glfs_fd_t
}

func (fd *Fd) Fsync() error {
	_, err := C.glfs_fsync(fd.fd)

	return err
}

func (fd *Fd) Pread(b []byte, off int64) (int, error) {
	n, err := C.glfs_pread(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0)

	return int(n), err
}

func (fd *Fd) Pwrite(b []byte, off int64) (int, error) {
	n, err := C.glfs_pwrite(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), C.off_t(off), 0)

	return int(n), err
}

func (fd *Fd) Read(b []byte) (int, error) {
	n, err := C.glfs_read(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), 0)

	return int(n), err
}

func (fd *Fd) Write(b []byte) (int, error) {

	n, err := C.glfs_write(fd.fd, unsafe.Pointer(&b[0]), C.size_t(len(b)), 0)
	return int(n), err

}
