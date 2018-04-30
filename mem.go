package posixipc

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// DefaultHeapSize is the set size of a Memory object's heap,
// referenced when using AllocDefaultHeap().
const DefaultHeapSize = 256

// Memory provides a segment of memory that can be shared
// between goroutines, passed via a mutex locker.
type Memory struct {
	mid int        // id
	mu  sync.Mutex // memory lock for sync

	desc MemDesc
	heap []byte
}

// MemDesc describes a shared piece of memory and is used
// as an identifier in the SharedMemory interface.
type MemDesc struct {
	mid int // id

	free   int // unused memory
	frames int // pages
}

// NewMemoryDefaultHeap creates a new Memory object that has a heap
// size of 256Kb, which is set as the default size when creating a
// new Memory object
func (m Memory) NewMemoryDefaultHeap() *Memory {
	heap := alloc(DefaultHeapSize)
	mid := m.NewMid()
	return &Memory{
		mid:  mid,
		heap: heap,
	}
}

// NewMemory creates a Memory object with a heap of size N
func (m Memory) NewMemory(size int) *Memory {
	heap := alloc(size)
	mid := m.generateMid()
	return &Memory{
		mid:  mid,
		heap: heap,
	}
}

func alloc(size int) []byte {
	heap := make([]byte, size)
	return heap
}

func fillSlice(mem []byte) {
	_, err := rand.Read(mem)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

func (m Memory) Dealloc(mid int) error {
	m.destroy(mid)
	return nil
}

func (m Memory) destroy(mid int) {
	// Release bytes created via Alloc()
}

// Mid is a unique id for a Memory type
type Mid struct {
	mid int
}

// NewMid creates a random integer that represents a Mid type
func (m Memory) NewMid() *Mid {
	mid := m.generateMid()
	return &Mid{mid: mid}
}

func (m Memory) generateMid() int {
	midInt := randomInt()
	return midInt
}

// Default min and max values for new Mids
const (
	// Smallest allowed integer representing a Mid
	MidMin = 1
	// Larget allowed integer representing a Mid
	MidMax = 999
)

// Returns an int >= min, < max
func randomInt() int {
	return MidMin + rand.Int(MidMax-MidMin)
}
