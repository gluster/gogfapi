package gfapi

// This file includes lower level operations on fd like the ones in the 'syscall' package

// #cgo pkg-config: glusterfs-api
// #include "glusterfs/api/glfs.h"
// #include <stdlib.h>
// #include <sys/stat.h>
import "C"
import (
	"os"
	"syscall"
	"unsafe"
)

// Fd is the glusterfs fd type
type Fd struct {
	fd *C.glfs_fd_t
}

var _zero uintptr

// Fchmod changes the mode of the Fd to the given mode
//
// Returns error on failure
func (fd *Fd) Fchmod(mode uint32) error {
	_, err := C.glfs_fchmod(fd.fd, C.mode_t(mode))

	return err
}

// Fstat performs an fstat call on the Fd and saves stat details in the passed stat structure
//
// Returns error on failure
func (fd *Fd) Fstat(stat *syscall.Stat_t) error {

	ret, err := C.glfs_fstat(fd.fd, (*C.struct_stat)(unsafe.Pointer(stat)))
	if int(ret) < 0 {
		return err
	}
	return nil
}

// Read reads at most len(b) bytes into b from Fd
//
// Returns number of bytes read on success and error on failure
func (fd *Fd) Read(b []byte) (n int, err error) {
	var p0 unsafe.Pointer

	if len(b) > 0 {
		p0 = unsafe.Pointer(&b[0])
	} else {
		p0 = unsafe.Pointer(&_zero)
	}

	// glfs_read returns a ssize_t. The value of which is the number of bytes written.
	// Unless, ret is -1, an error, implying to check errno. cgo collects errno as the
	// functions error return value.
	ret, e1 := C.glfs_read(fd.fd, p0, C.size_t(len(b)), 0)
	n = int(ret)
	if n < 0 {
		err = e1
	}

	return n, err
}

// Write writes len(b) bytes from b into the Fd
//
// Returns number of bytes written on success and error on failure
func (fd *Fd) Write(b []byte) (n int, err error) {
	var p0 unsafe.Pointer

	if len(b) > 0 {
		p0 = unsafe.Pointer(&b[0])
	} else {
		p0 = unsafe.Pointer(&_zero)
	}

	// glfs_write returns a ssize_t. The value of which is the number of bytes written.
	// Unless, ret is -1, an error, implying to check errno. cgo collects errno as the
	// functions error return value.
	ret, e1 := C.glfs_write(fd.fd, p0, C.size_t(len(b)), 0)
	n = int(ret)
	if n < 0 {
		err = e1
	}

	return n, err
}

func (fd *Fd) lseek(offset int64, whence int) (int64, error) {
	ret, err := C.glfs_lseek(fd.fd, C.off_t(offset), C.int(whence))

	return int64(ret), err
}

func (fd *Fd) Fallocate(mode int, offset int64, len int64) error {
	ret, err := C.glfs_fallocate(fd.fd, C.int(mode),
		C.off_t(offset), C.size_t(len))

	if ret == 0 {
		err = nil
	}
	return err
}

func (fd *Fd) Fgetxattr(attr string, dest []byte) (int64, error) {
	var ret C.ssize_t
	var err error

	cattr := C.CString(attr)
	defer C.free(unsafe.Pointer(cattr))

	if len(dest) <= 0 {
		ret, err = C.glfs_fgetxattr(fd.fd, cattr, nil, 0)
	} else {
		ret, err = C.glfs_fgetxattr(fd.fd, cattr,
			unsafe.Pointer(&dest[0]), C.size_t(len(dest)))
	}

	if ret >= 0 {
		return int64(ret), nil
	} else {
		return int64(ret), err
	}
}

func (fd *Fd) Fsetxattr(attr string, data []byte, flags int) error {

	cattr := C.CString(attr)
	defer C.free(unsafe.Pointer(cattr))

	ret, err := C.glfs_fsetxattr(fd.fd, cattr,
		unsafe.Pointer(&data[0]), C.size_t(len(data)),
		C.int(flags))

	if ret == 0 {
		err = nil
	}
	return err
}

func (fd *Fd) Fremovexattr(attr string) error {

	cattr := C.CString(attr)
	defer C.free(unsafe.Pointer(cattr))

	ret, err := C.glfs_fremovexattr(fd.fd, cattr)

	if ret == 0 {
		err = nil
	}
	return err
}

func direntName(dirent *syscall.Dirent) string {
	name := make([]byte, 0, len(dirent.Name))
	for i, c := range dirent.Name {
		if c == 0 || i > 255 {
			break
		}

		name = append(name, byte(c))
	}

	return string(name)
}

// Readdir returns the information of files in a directory.
//
// n is the maximum number of items to return. If there are more items than
// the maximum they can be obtained in successive calls. If maximum is 0
// then all the items will be returned.
func (fd *Fd) Readdir(n int) ([]os.FileInfo, error) {
	var (
		stat  syscall.Stat_t
		files []os.FileInfo
		statP = (*C.struct_stat)(unsafe.Pointer(&stat))
	)

	for i := 0; n == 0 || i < n; i++ {
		d, err := C.glfs_readdirplus(fd.fd, statP)
		if err != nil {
			return nil, err
		}

		dirent := (*syscall.Dirent)(unsafe.Pointer(d))
		if dirent == nil {
			break
		}

		name := direntName(dirent)
		file := fileInfoFromStat(&stat, name)
		files = append(files, file)
	}

	return files, nil
}

// Readdirnames returns the names of files in a directory.
//
// n is the maximum number of items to return and works the same way as Readdir.
func (fd *Fd) Readdirnames(n int) ([]string, error) {
	var names []string

	for i := 0; n == 0 || i < n; i++ {
		d, err := C.glfs_readdir(fd.fd)
		if err != nil {
			return nil, err
		}

		dirent := (*syscall.Dirent)(unsafe.Pointer(d))
		if dirent == nil {
			break
		}

		name := direntName(dirent)
		names = append(names, name)
	}

	return names, nil
}
