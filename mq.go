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

type imq interface {
	// Open establishes the connection between a process and a message queue
	// with a message queue descriptor. It creates an open message queue
	// description that refers to the new message queue, used by other functions
	// to refer to the message queue
	Open(name *string, oflag int) error // mq_open
	// Close removes the association between the message queue descriptor,
	// `mqdes`, and its message queue.
	Close(mqd_t mqdes) error // mq_close
	// Unlink removes the message queue named by the pathname `name`
	Unlink(name string) error // mq_unlink
	// Send adds the message pointed to by the argument `msg_ptr` to the
	// message queue specified by `mydes`. The `msg_len` argument specifies
	// the length of the message in bytes pointed to by `msg_ptr`. The value
	// of `msg_len` is less than or equal to the `mq_msgsize` or the call fails
	Send(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, msg_prio int) error // mq_send
	// Recieve is used to recieve the oldest of the highest priority message(s)
	// from the message queue specified by `mqdes`. The message is remove from
	// the queue and copied to the buffer pointed to by `msg_ptr`
	Recieve(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, int msg_prio) error // mq_recieve
	// Notify provides an asynchronous mechanism for processes to recieve notice
	// that messages are available in a message queue, rather than synchonously
	// blocking (waiting) in mq_recieve
	Notify(mqd_t mqdes, notification sigevent) error
	// SetAttr is used to set attributes associated with the open message queue description referenced by the message queue descriptor specified by mqdes.
	SetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error
	// GetAttr is used to get attributes associated with the open message queue description referenced by the message queue descriptor specified by mqdes.
	GetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error
}

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
