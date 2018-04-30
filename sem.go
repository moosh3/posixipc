package posixipc

import (
	"errors"
	"sync"
)

// ErrNoTickets
var ErrNoTickets = errors.New("could not acquire semaphore ticket")

// Semaphore
type Semaphore struct {
	sync.RWMutex
	name string

	buf SemBuff
}

// SemBuff
type SemBuff struct {
	index     int // # of semaphore element
	operation int // operation to perform
	oflags    int // operations specific options
}

// NewSemaphore returns a new instance of a Semaphore locking mechanism
func NewSemaphore(name string) *Semaphore {
	return &Semaphore{name: name}
}

// Open
func (s *Semaphore) Open(name *string, oflag int) error {

}

// O_EXCL
func (s *Semaphore) openReadWrite() error { return nil }

// O_CREAT
func (s *Semaphore) openCreate() error { return nil }

// Init
func (s *Semaphore) Init(sem *Semaphore, pshared int, value int) error { return nil }

// Close
func (s *Semaphore) Close(sem *Semaphore) error { return nil }

// Unlink
func (s *Semaphore) Unlink(name *string) error { return nil }

// Destroy
func (s *Semaphore) Destroy(sem *Semaphore) error { return nil }

// GetValue
func (s *Semaphore) GetValue(sem *Semaphore, sval *int) error { return nil }
