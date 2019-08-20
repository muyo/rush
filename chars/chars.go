package chars

const (
	uint64Digits = 20
	uint64Max    = 1<<64 - 1
	uint64Cutoff = uint64Max / 10

	uint32Digits = 10
	uint32Max    = 1<<32 - 1
	uint32Cutoff = uint32Max / 10

	uint16Digits = 5
	uint16Max    = 1<<16 - 1
	uint16Cutoff = uint16Max / 10

	uint8Max    = 1<<8 - 1
	uint8Digits = 3
)
