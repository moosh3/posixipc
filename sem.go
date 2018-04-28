package posixipc

import (
	"errors"
	"sync"
)

var ErrNoTickets = errors.New("could not acquire semaphore ticket")

type sem struct {
	sync.RWMutex
	name string
}

type sembuf struct {
	sem_num int // # of semaphore element
	sem_op  int // operation to perform
	sem_flq int // operations specific options
}

func (s *sem) Open(name string, oflag int) error {

}

// O_EXCL
func (s *sem) openReadWrite() error { return nil }

// O_CREAT
func (s *sem) openCreate() error { return nil }

func (s *sem) Init(sem_t *sem, pshared int, value int) error { return nil }
func (s *sem) Close(sem_t *sem) error                        { return nil }
func (s *sem) Unlink(name *string) error                     { return nil }
func (s *sem) Destroy(sem_t *sem) error                      { return nil }
func (s *sem) GetValue(sem_t *sem, sval *int) error          { return nil }
