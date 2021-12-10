package losetup

import (
	"fmt"
	"github.com/daedaluz/golosetup/loopctl"
	"os"
)

type Device struct {
	slot     int64
	f        *os.File
	readOnly bool
}

func NewDevice(slot int64) *Device {
	return &Device{slot: slot}
}

func (d *Device) Path() string {
	return fmt.Sprintf("/dev/loop%d", d.slot)
}

func (d *Device) PartitionPath(n int) string {
	return fmt.Sprintf("/dev/loop%dp%d", d.slot, n)
}

func (d *Device) GetSlot() int64 {
	return d.slot
}

func (d *Device) Open(readOnly bool) (err error) {
	oFlags := os.O_RDWR
	if readOnly {
		oFlags = os.O_RDONLY
	}
	d.f, err = loopctl.OpenLoop(d.slot, oFlags)
	d.readOnly = readOnly
	return err
}

func (d *Device) Close() error {
	return d.f.Close()
}

func (d *Device) GetInfo() (*loopctl.Info, error) {
	return loopctl.GetStatus(d.f)
}

func (d *Device) SetInfo(info *loopctl.Info) error {
	return loopctl.SetStatus(d.f, info)
}

func (d *Device) Detach() error {
	return loopctl.ClrFd(d.f)
}

func (d *Device) Attach(fileName string, offset uint64, flags loopctl.LoopFlag) error {
	flags &= ^(loopctl.ReadOnly)
	if d.readOnly {
		flags |= loopctl.ReadOnly
	}
	oFlags := os.O_RDWR
	if (flags & loopctl.ReadOnly) > 0 {
		oFlags = os.O_RDONLY
	}
	targetFile, err := os.OpenFile(fileName, oFlags, 0660)
	if err != nil {
		return err
	}
	defer targetFile.Close()
	err = loopctl.SetFd(d.f, targetFile)
	if err != nil {
		return err
	}

	info := &loopctl.Info{
		Offset: offset,
		Flags:  flags,
	}
	copy(info.FileName[:], fileName)
	if err := loopctl.SetStatus(d.f, info); err != nil {
		_ = loopctl.ClrFd(d.f)
		return err
	}
	return nil
}

func (d *Device) String() string {
	return fmt.Sprintf("LoopDev(%d)", d.slot)
}

func GetFree() (*Device, error) {
	ctrl, err := loopctl.OpenLoopCTL()
	if err != nil {
		return nil, err
	}
	defer ctrl.Close()
	free, err := loopctl.GetFree(ctrl)
	if err != nil {
		return nil, err
	}
	return &Device{
		slot: free,
	}, nil
}

func Attach(file string, offset uint64, flags loopctl.LoopFlag) (*Device, error) {
	dev, err := GetFree()
	if err != nil {
		return nil, err
	}
	if err := dev.Open((flags & loopctl.ReadOnly) > 0); err != nil {
		return nil, err
	}
	defer dev.Close()
	err = dev.Attach(file, offset, flags)
	if err != nil {
		return nil, err
	}
	return dev, nil
}
