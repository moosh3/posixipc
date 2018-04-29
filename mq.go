package posixipc

/*
#include <stdio.h>
#include <errno.h>
#include <signal.h>
#include <pthread.h>
#include <mqueue.h>
*/
import "C"

import (
	"fmt"
	"sync"
)

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

func (m *MessageQueue) String() string {
	return fmt.Printf("Message Queue name: %s", mq.name)
}

// descriptor used for operations
type MqDesc struct {
	pid int
}

// ProcFromPid returns the process set in any given MqDesc struct,
// which describes a certain MessageQueue.
func (md *MyDesc) ProcFromPid(pid int) *os.Process {
	return os.FindProcess(pid)
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
func (m *mq) Open(name string, oflag int) error {

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
func (m *MessageQueue) Send(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, msg_prio int) error {
	return nil
}

// Recieve
func (m *MessageQueue) Recieve(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, int msg_prio) error {
	return nil
}

// Close
func (m *MessageQueue) Close(mqd_t mqdes) error { return nil }

// Unlink
func (m *MessageQueue) Unlink(name string) error { return nil }

// Notify
func (m *MessageQueue) Notify(mqd_t mqdes, notification sigevent) error { return nil }

// SetAttr
func (m *MessageQueue) SetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error { return nil }

// GetAttr
func (m *MessageQueue) GetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error { return nil }
