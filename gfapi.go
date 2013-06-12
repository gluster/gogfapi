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

type Volume struct {
	fs *C.glfs_t
}

type File struct {
	fd *C.glfs_fd_t
}

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

func (v *Volume) Mount() int {
	ret := C.glfs_init(v.fs)

	return int(ret)
}

func (v *Volume) Unmount() int {
	ret := C.glfs_fini(v.fs)

	return int(ret)
}

func (v *Volume) Create(name string) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_creat(v.fs, cname, C.int(os.O_RDWR|os.O_CREATE|os.O_TRUNC), 0666)

	if cfd == nil {
		return nil, &os.PathError{"create", name, err}
	}

	return &File{cfd}, nil
}

func (v *Volume) Open(name string) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_open(v.fs, cname, C.int(os.O_RDONLY))

	if cfd == nil {
		return nil, &os.PathError{"open", name, err}
	}

	return &File{cfd}, nil
}
// TODO: Add support for 'perm FileMode'
func (v *Volume) OpenFile(name string, flags int, perm os.FileMode) (*File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cfd, err := C.glfs_open(v.fs, cname, C.int(flags))

	if cfd == nil {
		return nil, &os.PathError{"open", name, err}
	}

	return &File{cfd}, nil
}

func (f *File) Close() error {
	_, err := C.glfs_close(f.fd)

	return err
}
