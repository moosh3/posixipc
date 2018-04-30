package posixipc

import (
	"fmt"
	"os"
	"sync"
)

type Priority int

// Msg's are placed onto MessageQueues
type Msg struct {
	bytes    []byte
	priority Priority
}

const (
	DefaultMaxQueues  = 1
	DefaultMaxMsg     = 2
	DefaultMaxSizeMsg = 1024
)

type MessageQueue struct {
	name string
	desc MqDesc
	attr MqAttr

	mu  sync.Mutex
	buf []byte
}

// NewMessageQueue returns a new instance of a MessageQueue
func NewMessageQueue(name string) *MessageQueue {
	return &MessageQueue{name: name}
}

func (m *MessageQueue) String() string {
	return fmt.Printf("mq: %s", m.name)
}

// MqDesc descriptor used for operations
type MqDesc struct {
	pid int
}

// ProcFromPid returns the process set in any given MqDesc struct,
// which describes a certain MessageQueue.
func (md *MqDesc) ProcFromPid(pid int) *os.Process {
	proc, _ := os.FindProcess(pid)
	return proc
}

// Key in queue for msgs
type Key string

// MqAttr stores queue attributes
type MqAttr struct {
	// opts for the queue; mq_setattr can change it
	OFlags int
	// # messages stored in a queue. 0 == N/A
	MaxMsgs int
	// # messages currently on the given queue
	CurrMsgs int
	// # of processes waiting to send a message
	SendWait int
	// # of process waiting to recieve a message
	RecieveWait int
}

// Open
func (m *MessageQueue) Open(name string, oflag int) error {
	return nil
}

// O_RDONLY
func (m *MessageQueue) openReadOnly() error { return nil }

// O_WRONLY
func (m *MessageQueue) openWriteOnly() error { return nil }

// O_RDWR
func (m *MessageQueue) openReadWrite() error { return nil }

// O_CREAT
func (m *MessageQueue) openCreate() error { return nil }

// Send
func (m *MessageQueue) Send(mq MqDesc, ptr uintptr, len []byte, priority Priority) error {
	return nil
}

// Recieve
func (m *MessageQueue) Recieve(mq MqDesc, ptr uintptr, len []byte, int Priority) error {
	return nil
}

// Close
func (m *MessageQueue) Close(mq MqDesc) error { return nil }

// Unlink
func (m *MessageQueue) Unlink(name string) error { return nil }

// Notify
func (m *MessageQueue) Notify(mq MqDesc, sig os.Signal) error { return nil }

type MqKey int

// SetAttr
func (m *MessageQueue) SetAttr(mq MqDesc, key MqKey, value interface{}) error { return nil }

// GetAttr
func (m *MessageQueue) GetAttr(mq MqDesc, key MqKey) error { return nil }
