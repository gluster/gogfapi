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

// #cgo pkg-config: glusterfs-api
// #include "api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import (
	"os"
	"unsafe"
)

// The gluster filesystem object. Represents the virtual filesystem.
type Volume struct {
	fs *C.glfs_t
}

// The gluster file object.
type File struct {
	fd *C.glfs_fd_t
}

// Init() initializes the Volume.
// This must be performed before calling Mount().
//
// host is the hostname/ip of a gluster server.
// volname is the name of a volume that you want to access.
//
// Return value is 0 for success and non 0 for failure.
func (v *Volume) Init(host string, volname string) int {
	cvolname := C.CString(volname)
	chost := C.CString(host)
	ctrans := C.CString("tcp")
	defer C.free(unsafe.Pointer(cvolname))
	defer C.free(unsafe.Pointer(chost))
	defer C.free(unsafe.Pointer(ctrans))

	v.fs = C.glfs_new(cvolname)

	ret := C.glfs_set_volfile_server(v.fs, ctrans, chost, 24007)

	return int(ret)
}

// Mount() performs the virtual mount.
// The Volume must be initalized before calling Mount().
//
// Return value is 0 for success and non 0 for failure.
func (v *Volume) Mount() int {
	ret := C.glfs_init(v.fs)

	return int(ret)
}

// Unmount() ends the virtual mount.
//
// Return value is 0 for success and non 0 for failure.
//
// BUG : Always returns non-zero presently. Better to ignore the return value for now.
func (v *Volume) Unmount() int {
	ret := C.glfs_fini(v.fs)

	return int(ret)
}

// Create() creates a file with given name on the the Volume v.
// The Volume must be mounted before calling Create().
// Create() is similar to os.Create() in its functioning.
//
// name is the name of the file to be create.
//
// Returns a File object on success and a os.PathError on failure.
func (v *Volume) Create(name string) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_creat(v.fs, cname, C.int(os.O_RDWR|os.O_CREATE|os.O_TRUNC), 0666)

	if cfd == nil {
		return nil, &os.PathError{"create", name, err}
	}

	return &File{cfd}, nil
}

// Open() opens the named file on the the Volume v.
// The Volume must be mounted before calling Open().
// Open() is similar to os.Open() in its functioning.
//
// name is the name of the file to be open.
//
// Returns a File object on success and a os.PathError on failure.
func (v *Volume) Open(name string) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_open(v.fs, cname, C.int(os.O_RDONLY))

	if cfd == nil {
		return nil, &os.PathError{"open", name, err}
	}

	return &File{cfd}, nil
}

// OpenFile() opens the named file on the the Volume v.
// The Volume must be mounted before calling OpenFile().
// OpenFile() is similar to os.OpenFile() in its functioning.
//
// name is the name of the file to be open.
// flags is the access mode of the file.
// perm is the permissions for the opened file.
//
// Returns a File object on success and a os.PathError on failure.
//
// BUG : perm is not used for opening the file.
func (v *Volume) OpenFile(name string, flags int, perm os.FileMode) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_open(v.fs, cname, C.int(flags))

	if cfd == nil {
		return nil, &os.PathError{"open", name, err}
	}

	return &File{cfd}, nil
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
	return ""
}

func (f *File) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (f *File) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, nil
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
	return nil
}

func (f *File) Truncate(size int64) error {
	return nil
}

func (f *File) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (f *File) WriteAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

func (f *File) WriteString(s string) (ret int, err error) {
	return 0, nil
}
