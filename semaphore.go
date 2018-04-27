package posixipc

type sem struct {
	name string
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
