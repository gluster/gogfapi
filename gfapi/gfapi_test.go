package gfapi

import (
	"os"
	"path/filepath"
	"testing"
)

/* The testcases assume that it is being run on a peer in a gluster cluster,
 * and that the cluster has a volume named "test"
 */

var (
	vol  *Volume
	file *File
)

func TestInit(t *testing.T) {

	vol = new(Volume)

	if vol == nil {
		t.Fatalf("Failed to allocate variable")
	}

	ret := vol.Init("localhost", "test")
	if ret != 0 {
		t.Fatalf("Failed to initialize volume. Ret = %d", ret)
	}
}

func TestSetLogging(t *testing.T) {
	ret, err := vol.SetLogging("/var/log/glusterfs/test.log", LogDebug)
	if ret != 0 && err != nil {
		t.Fatalf("Unable to set Logging ret = %d; error = %v", ret, err)
	}
}

func TestMount(t *testing.T) {
	ret := vol.Mount()
	if ret != 0 {
		t.Fatalf("Failed to mount volume. Ret = %d", ret)
	}
}

func TestMkdirAll(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestMkdirAll_/dir/./dir2"
	err := vol.MkdirAll(path, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", path, err)
	}

	// Already exists, should succeed.
	err = vol.MkdirAll(path, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q (second time): %s", path, err)
	}

	// Make file.
	fpath := path + "/file"
	f, err := vol.Create(fpath)
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	defer f.Close()

	// Can't make directory named after file.
	err = vol.MkdirAll(fpath, 0777)
	if err == nil {
		t.Fatalf("MkdirAll %q: no error", fpath)
	}
	perr, ok := err.(*os.PathError)
	if !ok {
		t.Fatalf("MkdirAll %q returned %T, not *PathError", fpath, err)
	}
	if filepath.Clean(perr.Path) != filepath.Clean(fpath) {
		t.Fatalf("MkdirAll %q returned wrong error path: %q not %q", fpath, filepath.Clean(perr.Path), filepath.Clean(fpath))
	}

	// Can't make subdirectory of file.
	ffpath := fpath + "/subdir"
	err = vol.MkdirAll(ffpath, 0777)
	if err == nil {
		t.Fatalf("MkdirAll %q: no error", ffpath)
	}

	perr, ok = err.(*os.PathError)
	if !ok {
		t.Fatalf("MkdirAll %q returned %T, not *PathError", ffpath, err)
	}
	if filepath.Clean(perr.Path) != filepath.Clean(fpath) {
		t.Fatalf("MkdirAll %q returned wrong error path: %q not %q", ffpath, filepath.Clean(perr.Path), filepath.Clean(fpath))
	}
}

func TestCreate(t *testing.T) {
	f, err := vol.Create("test")

	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}
	file = f
}

func TestClose1(t *testing.T) {
	err := file.Close()
	if err != nil {
		t.Errorf("Failed to close file. Error = %v", err)
	}
	file = nil
}

func TestOpen(t *testing.T) {
	f, err := vol.Open("test")

	if err != nil {
		t.Fatalf("Failed to open file. Error = %v", err)
	}
	file = f
}

func TestClose2(t *testing.T) {
	err := file.Close()
	if err != nil {
		t.Errorf("Failed to close file. Error = %v", err)
	}
	file = nil
}

func TestUnlink(t *testing.T) {
	f, err := vol.Create("/TestUnlink")
	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}
	f.Close()

	err = vol.Unlink("/TestUnlink")
	if err != nil {
		t.Errorf("vol.Unlink failed . Error = %v", err)
	}
}

func TestRmdir(t *testing.T) {
	err := vol.Mkdir("/TestRmdir", 0755)
	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}

	err = vol.Rmdir("/TestRmdir")
	if err != nil {
		t.Errorf("vol.Rmdir failed . Error = %v", err)
	}
}

func TestRename(t *testing.T) {
	f, err := vol.Create("TestRename")
	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}
	f.Close()

	err = vol.Rename("TestRename", "TestRenameNew")
	if err != nil {
		t.Errorf("vol.Rename failed . Error = %v", err)
	}
}

func TestFxattrs(t *testing.T) {

	f, err := vol.Create("/testFxattrs")
	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}
	defer f.Close()

	err = f.Setxattr("user.glusterfs", []byte("Gluster is awesome!"), 0)
	if err != nil {
		t.Errorf("f.Setxattr() failed. Error = %v", err)
	}

	size, err := f.Getxattr("user.glusterfs", nil)
	if err != nil {
		t.Errorf("f.Getxattr() failed. Error = %v", err)
	}
	buf := make([]byte, size)
	size, err = f.Getxattr("user.glusterfs", buf)
	if err != nil {
		t.Errorf("f.Getxattr() failed. Error = %v", err)
	}

	if "Gluster is awesome!" != string(buf[:size]) {
		t.Errorf("xattrs do not match")
	}

	err = f.Removexattr("user.glusterfs")
	if err != nil {
		t.Errorf("f.Removexattr() failed. Error = %v", err)
	}
}

func TestXattrs(t *testing.T) {

	path := "/testXattrs"
	f, err := vol.Create(path)
	if err != nil {
		t.Fatalf("Failed to create file. Error = %v", err)
	}
	f.Close()

	err = vol.Setxattr(path, "user.glusterfs", []byte("Gluster is awesome!"), 0)
	if err != nil {
		t.Errorf("vol.Setxattr() failed. Error = %v", err)
	}

	size, err := vol.Getxattr(path, "user.glusterfs", nil)
	if err != nil {
		t.Errorf("vol.Getxattr() failed. Error = %v", err)
	}
	buf := make([]byte, size)
	size, err = vol.Getxattr(path, "user.glusterfs", buf)
	if err != nil {
		t.Errorf("vol.Getxattr() failed. Error = %v", err)
	}

	if "Gluster is awesome!" != string(buf[:size]) {
		t.Errorf("xattrs do not match")
	}

	err = vol.Removexattr(path, "user.glusterfs")
	if err != nil {
		t.Errorf("vol.Removexattr() failed. Error = %v", err)
	}
}

func TestUnmount(t *testing.T) {
	ret := vol.Unmount()
	if ret != 0 {
		t.Logf("Failed to unmount volume. Ret = %d", ret)
	}
}
