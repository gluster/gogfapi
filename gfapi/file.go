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

// This file includes higher level operations on files, such as those provided by the 'os' package

// #cgo pkg-config: glusterfs-api
// #include "api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import (
	"os"
)

// The gluster file object.
type File struct {
        name string
	Fd
}

// Close() closes an open File.
// Close() is similar to os.Close() in its functioning.
//
// Returns an Error on failure.
func (f *File) Close() error {
	_, err := C.glfs_close(f.fd)

	return err
}

func (f *File) Chdir() error {
	return nil
}

func (f *File) Chmod(mode os.FileMode) error {
	return nil
}

func (f *File) Chown(uid, gid int) error {
	return nil
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Read(b []byte) (n int, err error) {
	return f.Fd.Read(b)
}

func (f *File) ReadAt(b []byte, off int64) (n int, err error) {
	return f.Fd.Pread(b, off)
}

func (f *File) Readdir(n int) (fi []os.FileInfo, err error) {
	return nil, nil
}

func (f *File) Readdirnames(n int) (names []string, err error) {
	return nil, nil
}

func (f *File) Seek(offset int64, whence int) (ret int64, err error) {
	return 0, nil
}

func (f *File) Stat() (fi os.FileInfo, err error) {
	return nil, nil
}

func (f *File) Sync() error {
        err := f.Fd.Fsync()
	return err
}

func (f *File) Truncate(size int64) error {
	return nil
}

func (f *File) Write(b []byte) (n int, err error) {
	return f.Fd.Write (b)
}

func (f *File) WriteAt(b []byte, off int64) (n int, err error) {
	return f.Fd.Pwrite(b, off)
}

func (f *File) WriteString(s string) (ret int, err error) {
	return f.Write([]byte(s))
}
