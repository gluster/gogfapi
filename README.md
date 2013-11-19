# gfapi
--
    import "forge.gluster.org/gogfapi/gogfapi.git/gfapi"

Package gfapi provides a wrapper around gfapi, the GlusterFS api, which is used
to access files/directories on a Gluster volume. The design tries to follow the
default go file handling functions provided by the os package as much as
possible.

Package gfapi provides a wrapper around gfapi, the GlusterFS api, which is used
to access files/directories on a Gluster volume. The design tries to follow the
default go file handling functions provided by the os package as much as
possible.

Package gfapi provides a wrapper around gfapi, the GlusterFS api, which is used
to access files/directories on a Gluster volume. The design tries to follow the
default go file handling functions provided by the os package as much as
possible.

## Usage

#### type Fd

```go
type Fd struct {
}
```


#### func (*Fd) Fsync

```go
func (fd *Fd) Fsync() error
```

#### func (*Fd) Pread

```go
func (fd *Fd) Pread(b []byte, off int64) (int, error)
```

#### func (*Fd) Pwrite

```go
func (fd *Fd) Pwrite(b []byte, off int64) (int, error)
```

#### func (*Fd) Read

```go
func (fd *Fd) Read(b []byte) (int, error)
```

#### func (*Fd) Write

```go
func (fd *Fd) Write(b []byte) (int, error)
```

#### type File

```go
type File struct {
	Fd
}
```

The gluster file object.

#### func (*File) Chdir

```go
func (f *File) Chdir() error
```

#### func (*File) Chmod

```go
func (f *File) Chmod(mode os.FileMode) error
```

#### func (*File) Chown

```go
func (f *File) Chown(uid, gid int) error
```

#### func (*File) Close

```go
func (f *File) Close() error
```
Close() closes an open File. Close() is similar to os.Close() in its
functioning.

Returns an Error on failure.

#### func (*File) Name

```go
func (f *File) Name() string
```
Name() returns the name of the opened file

#### func (*File) Read

```go
func (f *File) Read(b []byte) (int, error)
```
Read() reads atmost len(b) bytes into b

Returns number of bytes read and an error if any

#### func (*File) ReadAt

```go
func (f *File) ReadAt(b []byte, off int64) (int, error)
```
ReadAt() reads atmost len(b) bytes into b starting from offset off

Returns number of bytes read and an error if any

#### func (*File) Readdir

```go
func (f *File) Readdir(n int) ([]os.FileInfo, error)
```

#### func (*File) Readdirnames

```go
func (f *File) Readdirnames(n int) ([]string, error)
```

#### func (*File) Seek

```go
func (f *File) Seek(offset int64, whence int) (int64, error)
```

#### func (*File) Stat

```go
func (f *File) Stat() (os.FileInfo, error)
```

#### func (*File) Sync

```go
func (f *File) Sync() error
```
Sync() commits the file to the storage

Returns error on failure

#### func (*File) Truncate

```go
func (f *File) Truncate(size int64) error
```

#### func (*File) Write

```go
func (f *File) Write(b []byte) (int, error)
```
Write() writes len(b) bytes to the file

Returns number of bytes written and an error if any

#### func (*File) WriteAt

```go
func (f *File) WriteAt(b []byte, off int64) (int, error)
```
Write() writes len(b) bytes to the file starting at offset off

Returns number of bytes written and an error if any

#### func (*File) WriteString

```go
func (f *File) WriteString(s string) (int, error)
```
WriteString() writes the contents of string s to the file

Returns number of bytes written and an error if any

#### type Volume

```go
type Volume struct {
}
```

The gluster filesystem object. Represents the virtual filesystem.

#### func (*Volume) Create

```go
func (v *Volume) Create(name string) (*File, error)
```
Create() creates a file with given name on the the Volume v. The Volume must be
mounted before calling Create(). Create() is similar to os.Create() in its
functioning.

name is the name of the file to be create.

Returns a File object on success and a os.PathError on failure.

#### func (*Volume) Init

```go
func (v *Volume) Init(host string, volname string) int
```
Init() initializes the Volume. This must be performed before calling Mount().

host is the hostname/ip of a gluster server. volname is the name of a volume
that you want to access.

Return value is 0 for success and non 0 for failure.

#### func (*Volume) Mount

```go
func (v *Volume) Mount() int
```
Mount() performs the virtual mount. The Volume must be initalized before calling
Mount().

Return value is 0 for success and non 0 for failure.

#### func (*Volume) Open

```go
func (v *Volume) Open(name string) (*File, error)
```
Open() opens the named file on the the Volume v. The Volume must be mounted
before calling Open(). Open() is similar to os.Open() in its functioning.

name is the name of the file to be open.

Returns a File object on success and a os.PathError on failure.

#### func (*Volume) OpenFile

```go
func (v *Volume) OpenFile(name string, flags int, perm os.FileMode) (*File, error)
```
OpenFile() opens the named file on the the Volume v. The Volume must be mounted
before calling OpenFile(). OpenFile() is similar to os.OpenFile() in its
functioning.

name is the name of the file to be open. flags is the access mode of the file.
perm is the permissions for the opened file.

Returns a File object on success and a os.PathError on failure.

BUG : perm is not used for opening the file.

#### func (*Volume) Unmount

```go
func (v *Volume) Unmount() int
```
Unmount() ends the virtual mount.

Return value is 0 for success and non 0 for failure.

BUG : Always returns non-zero presently. Better to ignore the return value for
now.
