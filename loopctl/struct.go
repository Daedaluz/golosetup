package loopctl

type Info struct {
	Device  uint64
	Inode   uint64
	RDevice uint64
	Offset  uint64
	// bytes; 0 == max available
	SizeLimit uint64

	Number         uint32
	EncryptType    EncryptionType
	EncryptKeySize uint32
	Flags          LoopFlag

	FileName   [64]byte
	CryptName  [64]byte
	EncryptKey [32]byte
	Init       [2]uint64
}

type Config struct {
	Fd        uint32
	BlockSize uint32
	Info      Info
	Reserved  [8]int64
}
