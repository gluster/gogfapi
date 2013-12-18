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

// This file includes higher level operations on files, such as those provided by the 'os' package

// #cgo pkg-config: glusterfs-api
// #include "api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import (
	"os"
        "syscall"
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

// Chmod() changes the mode of the file to the given mode
//
// Returns an error on failure
func (f *File) Chmod(mode os.FileMode) error {
	return f.Fd.Fchmod(posixMode(mode))
}

func (f *File) Chown(uid, gid int) error {
	return nil
}

// Name() returns the name of the opened file
func (f *File) Name() string {
	return f.name
}

// Read() reads atmost len(b) bytes into b
//
// Returns number of bytes read and an error if any
func (f *File) Read(b []byte) (int, error) {
	return f.Fd.Read(b)
}

// ReadAt() reads atmost len(b) bytes into b starting from offset off
//
// Returns number of bytes read and an error if any
func (f *File) ReadAt(b []byte, off int64) (int, error) {
	return f.Fd.Pread(b, off)
}

func (f *File) Readdir(n int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *File) Readdirnames(n int) ([]string, error) {
	return nil, nil
}

// Seek() sets the offset for the next read or write on the file based on whence,
// 0 - relative to beginning of file, 1 - relative to current offset, 2 - relative to end
//
// Returns new offset and an error if any
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return f.Fd.lseek(offset, whence)
}

func (f *File) Stat() (os.FileInfo, error) {
        var stat syscall.Stat_t
        err := f.Fd.Fstat(&stat)

        if err != nil {
                return nil, err
        } else {
                return fileInfoFromStat(&stat, f.name), nil
        }
}

// Sync() commits the file to the storage
//
// Returns error on failure
func (f *File) Sync() error {
	err := f.Fd.Fsync()
	return err
}

// Truncate() changes the size of the file
//
// Returns error on failure
func (f *File) Truncate(size int64) error {
	return f.Fd.Ftruncate(size)
}

// Write() writes len(b) bytes to the file
//
// Returns number of bytes written and an error if any
func (f *File) Write(b []byte) (int, error) {
	return f.Fd.Write(b)
}

// Write() writes len(b) bytes to the file starting at offset off
//
// Returns number of bytes written and an error if any
func (f *File) WriteAt(b []byte, off int64) (int, error) {
	return f.Fd.Pwrite(b, off)
}

// WriteString() writes the contents of string s to the file
//
// Returns number of bytes written and an error if any
func (f *File) WriteString(s string) (int, error) {
	return f.Write([]byte(s))
}
