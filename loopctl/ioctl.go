package loopctl

import (
	"fmt"
	ioctl "github.com/daedaluz/goioctl"
	"os"
	"unsafe"
)

const (
	// Previous ioctls, no longer used.
	//	ctlSetStatus    = 0x4C01
	//	ctlGetStatus    = 0x4C01

	ctlSetFD        = 0x4C00
	ctlClrFD        = 0x4C01
	ctlSetStatus64  = 0x4C01
	ctlGetStatus64  = 0x4C01
	ctlChangeFD     = 0x4C01
	ctlSetCapacity  = 0x4C01
	ctlSetDirectIO  = 0x4C01
	ctlSetBlockSize = 0x4C01
	ctlConfigure    = 0x4C01
)

const (
	ctlAdd     = 0x4C80
	ctlRemove  = 0x4C81
	ctlGetFree = 0x4C82
)

// SetFd
// Associate the loop device with the open file descriptor
func SetFd(loopFd, fileFd *os.File) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlSetFD, fileFd.Fd())
}

// ClrFd
// Disassociate the loop device from any file descriptor.
func ClrFd(loopFd *os.File) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlClrFD, 0)
}

// SetStatus
// Set the status of the loop device.
// Settable flags: AutoClear | PartScan
// Clearable flags: AutoClear
func SetStatus(loopFd *os.File, info *Info) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlSetStatus64, uintptr(unsafe.Pointer(info)))
}

// GetStatus
// Get the status of the loop device.
func GetStatus(loopFd *os.File) (*Info, error) {
	info := &Info{}
	err := ioctl.Ioctl(loopFd.Fd(), ctlGetStatus64, uintptr(unsafe.Pointer(info)))
	return info, err
}

// ChangeFd
// Switch  the  backing  store  of  the  loop  device  to the new file identified file
// descriptor specified by fileFd.
// This operation is possible only if the loop device is read-only and the new backing
// store is the same size and type as the old backing store.
func ChangeFd(loopFd, fileFd *os.File) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlChangeFD, fileFd.Fd())
}

// SetCapacity
// Resize a live loop device. One can change the size of the underlying backing store
// and then use this operation so that the loop driver learns about the new size.
func SetCapacity(loopFd *os.File) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlSetCapacity, 0)
}

// SetDirectIO
// Set DIRECT I/O mode on the loop device, so that it can be used to open backing file.
func SetDirectIO(loopFd *os.File, dio bool) error {
	param := uintptr(0)
	if dio {
		param = 1
	}
	return ioctl.Ioctl(loopFd.Fd(), ctlSetDirectIO, param)
}

// SetBlockSize
// Set the block size of the loop device.
// This value must be a power of two in the range [512,pagesize];
// otherwise, an EINVAL error results.
func SetBlockSize(loopFd *os.File, size uint64) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlSetBlockSize, uintptr(size))
}

// Configure
// Setup and configure all loop device parameters in a single step.
func Configure(loopFd *os.File, cfg *Config) error {
	return ioctl.Ioctl(loopFd.Fd(), ctlConfigure, uintptr(unsafe.Pointer(cfg)))
}

// GetFree
// Allocate or find a free loop device for use.
func GetFree(loopCtlFd *os.File) (int64, error) {
	return ioctl.IoctlX(loopCtlFd.Fd(), ctlGetFree, 0)
}

// Add
// Add the new loop device whose device number is n.
// If the device is already  allocated, the call fails with the error EEXIST
func Add(loopCtlFd *os.File, n uint64) (int64, error) {
	return ioctl.IoctlX(loopCtlFd.Fd(), ctlAdd, uintptr(n))
}

// Remove
// Remove  the  loop  device whose device number is n.
// If the device is in use, the call fails with the error EBUSY.
func Remove(loopCtlFd *os.File, n uint64) error {
	return ioctl.Ioctl(loopCtlFd.Fd(), ctlRemove, uintptr(n))
}

// OpenLoopCTL
// Opens the loop control device.
func OpenLoopCTL() (*os.File, error) {
	return os.OpenFile("/dev/loop-control", os.O_RDWR, 0660)
}

// OpenLoop
// Opens a loop device by number.
func OpenLoop(n int64, flags int) (*os.File, error) {
	devName := fmt.Sprintf("/dev/loop%d", n)
	return os.OpenFile(devName, flags, 0660)
}
