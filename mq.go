package posixipc

// Implemented in runtime package.
func runtime_BeforeExec()
func runtime_AfterExec()

type mq struct {
	desc mqDesc
	attr mqattr
}

type mqdes struct{}
type mqstat struct{}
type mqattr struct{}

type sigevent string
type key_t string

// Open
func (m *mq) Open(name string, oflag int) error {

}

// O_RDONLY
func (m *mq) openReadOnly() error { return nil }

// O_WRONLY
func (m *mq) openWriteOnly() error { return nil }

// O_RDWR
func (m *mq) openReadWrite() error { return nil }

// O_CREAT
func (m *mq) openCreate() error { return nil }

func (m *mp) Send(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, msg_prio int) error    { return nil }
func (m *mq) Recieve(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, int msg_prio) error { return nil }
func (m *mq) Close(mqd_t mqdes) error                                                  { return nil }
func (m *mq) Unlink(name string) error                                                 { return nil }
func (m *mq) Notify(mqd_t mqdes, notification sigevent) error                          { return nil }
func (m *mq) SetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error               { return nil }
func (m *mq) GetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error               { return nil }
