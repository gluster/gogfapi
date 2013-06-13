# gfapi
--
    import "forge.gluster.org/gogfapi/gogfapi.git/gfapi"

Package gfapi provides a wrapper around gfapi, the GlusterFS api, which is used to access files/directories on a Gluster volume.
The design tries to follow the default go file handling functions provided by the os package as much as possible.

## Usage

#### type File

    type File struct {
    }


The gluster file object.

#### func (*File) Close

    func (f *File) Close() error

Close() closes an open File. Close() is similar to os.Close() in its
functioning.

Returns an Error on failure.

#### type Volume

    type Volume struct {
    }


The gluster filesystem object. Represents the virtual filesystem.

#### func (*Volume) Create

    func (v *Volume) Create(name string) (*File, error)

Create() creates a file with given name on the the Volume v. The Volume must be
mounted before calling Create(). Create() is similar to os.Create() in its
functioning.

name is the name of the file to be create.

Returns a File object on success and a os.PathError on failure.

#### func (*Volume) Init

    func (v *Volume) Init(host string, volname string) int

Init() initializes the Volume. This must be performed before calling Mount().

host is the hostname/ip of a gluster server. volname is the name of a volume
that you want to access.

Return value is 0 for success and non 0 for failure.

#### func (*Volume) Mount

    func (v *Volume) Mount() int

Mount() performs the virtual mount. The Volume must be initalized before calling
Mount().

Return value is 0 for success and non 0 for failure.

#### func (*Volume) Open

    func (v *Volume) Open(name string) (*File, error)

Open() opens the named file on the the Volume v. The Volume must be mounted
before calling Open(). Open() is similar to os.Open() in its functioning.

name is the name of the file to be open.

Returns a File object on success and a os.PathError on failure.

#### func (*Volume) OpenFile

    func (v *Volume) OpenFile(name string, flags int, perm os.FileMode) (*File, error)

OpenFile() opens the named file on the the Volume v. The Volume must be mounted
before calling OpenFile(). OpenFile() is similar to os.OpenFile() in its
functioning.

name is the name of the file to be open. flags is the access mode of the file.
perm is the permissions for the opened file.

Returns a File object on success and a os.PathError on failure.

BUG : perm is not used for opening the file.

#### func (*Volume) Unmount

    func (v *Volume) Unmount() int

Unmount() ends the virtual mount.

Return value is 0 for success and non 0 for failure.

BUG : Always returns non-zero presently. Better to ignore the return value for
now.
