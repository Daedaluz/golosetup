package loopctl

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
