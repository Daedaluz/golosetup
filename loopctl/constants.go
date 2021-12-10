package loopctl

import "strings"

type LoopFlag uint32

const (
	// ReadOnly
	// The loopback device is read-only.
	ReadOnly = LoopFlag(1)
	// AutoClear
	// The loopback device will autodestruct on last close.
	AutoClear = LoopFlag(4)
	// PartScan
	// Allow automatic partition scanning.
	PartScan = LoopFlag(8)
	// DirectIO
	// Use direct I/O mode to access the backing file.
	DirectIO = LoopFlag(16)
)

func (f LoopFlag) String() string {
	var flags []string
	if (f & ReadOnly) > 0 {
		flags = append(flags, "ReadOnly")
	}
	if (f & AutoClear) > 0 {
		flags = append(flags, "AutoClear")
	}
	if (f & PartScan) > 0 {
		flags = append(flags, "PartScan")
	}
	if (f & DirectIO) > 0 {
		flags = append(flags, "DirectIO")
	}
	return strings.Join(flags, "|")
}

type EncryptionType uint32

const (
	CryptNone = EncryptionType(iota)
	CryptXOR
	CryptDES
	CryptFish2
	CryptBlow
	CryptCast128
	CryptIdea
	CryptDummy
	CryptSkipJack
	CryptCryptoAPI
)
