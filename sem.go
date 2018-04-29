package posixipc

import (
	"errors"
	"sync"
)

var ErrNoTickets = errors.New("could not acquire semaphore ticket")

type Semaphore struct {
	sync.RWMutex
	name string

	buf SemBuff
}

type SemBuff struct {
	index     int // # of semaphore element
	operation int // operation to perform
	oflags    int // operations specific options
}

// Open
func (s *Semaphore) Open(name string, oflag int) error {

}

// O_EXCL
func (s *Semaphore) openReadWrite() error { return nil }

// O_CREAT
func (s *Semaphore) openCreate() error { return nil }

// Init
func (s *Semaphore) Init(sem_t *sem, pshared int, value int) error { return nil }

// Close
func (s *Semaphore) Close(sem_t *sem) error { return nil }

// Unlink
func (s *Semaphore) Unlink(name *string) error { return nil }

// Destroy
func (s *Semaphore) Destroy(sem_t *sem) error { return nil }

// GetValue
func (s *Semaphore) GetValue(sem_t *sem, sval *int) error { return nil }
