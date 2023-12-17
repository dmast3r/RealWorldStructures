package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Generator defines the structure for our UUID generator.
type Generator struct {
	mx         sync.Mutex // Protects access to the below fields to ensure thread safety.
	lastEpoch  int64      // Stores the last epoch timestamp when a UUID was generated.
	machineId  int64      // Unique identifier of the machine where this generator is running.
	seq        int64      // Counter for UUIDs generated within the same millisecond.
}

// NewGenerator initializes and returns a new Generator.
func NewGenerator() *Generator {
	// In a production setting, the machine ID should be assigned in a manner that ensures uniqueness across all machines.
	// Using a service like Zookeeper or etcd can help achieve this.
	return &Generator{
		mx:         sync.Mutex{},
		lastEpoch:  time.Now().UnixMilli(),
		machineId:  rand.Int63n(MaxMachineId), // Randomly assigned for demonstration; replace with a stable machine ID in production.
		seq:        0,                          // Sequence starts at 0 for each new Generator instance.
	}
}

// nextId generates the next UUID.
func (generator *Generator) nextId() (int64, error) {
	currentEpoch := time.Now().UnixMilli()

	// If the current epoch is less than the last recorded epoch, it indicates a system clock issue.
	if currentEpoch < generator.lastEpoch {
		return 0, errors.New("current time is earlier than the last generated time, indicating a possible clock issue")
	}

	generator.mx.Lock() // Ensure that sequence number changes are atomic.
	defer generator.mx.Unlock()

	if currentEpoch == generator.lastEpoch {
		// If the UUID is generated in the same millisecond as the last UUID, increment the sequence.
		generator.seq = (generator.seq + 1) & MaxSeq
		if generator.seq == 0 {
			// If the sequence number overflows, wait until the next millisecond.
			for currentEpoch <= generator.lastEpoch {
				currentEpoch = time.Now().UnixMilli()
			}
		}
	} else {
		// If the UUID is generated in a new millisecond, reset the sequence to 0.
		generator.seq = 0
	}

	// Update lastEpoch to the current time.
	generator.lastEpoch = currentEpoch

	// Construct the next UUID using bit shifting to combine the timestamp, machine ID, and sequence number.
	nextId := ((currentEpoch - BaseEpoch) << TimestampPad) | (generator.machineId << MachineIdPad) | generator.seq

	return nextId, nil
}

// GetNext provides a thread-safe way to obtain the next UUID.
func (generator *Generator) GetNext() (int64, error) {
	return generator.nextId()
}