package main

const (
	BaseEpoch    = 1288834974657  // Twitter's Epoch value 
	MaxSeq       = 1 << 12 // the max possible value for sequence within a milliseconds on the same machine. I takes at most 12 bits.
	MaxMachineId = 1 << 10 // the max machine id that is possible, the machine id id takes at most 22 bits
	MachineIdPad = 12 // number of bits to the right of machine ID
	TimestampPad = 22      // number of bits to the right of timestamp
)
