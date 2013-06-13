package gogfapi

import "testing"

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

func TestMount(t *testing.T) {
	ret := vol.Mount()
	if ret != 0 {
		t.Fatalf("Failed to mount volume. Ret = %d", ret)
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

func TestUnmount(t *testing.T) {
	ret := vol.Unmount()
	if ret != 0 {
		t.Logf("Failed to unmount volume. Ret = %d", ret)
	}
}
