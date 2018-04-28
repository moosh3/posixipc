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

type mq struct {
	name string
	desc mg_des
	attr mq_attr

	mu  sync.Mutex
	buf []byte
}

func (m *mq) String() string {
	return fmt.Printf("mq name: %s", mq.name)
}

// descriptor used for operations
type mqdes struct {
	pid int
}

// key in queue for msgs
type key_t string

// store queue attributes
type mq_attr struct {
	// opts for the queue; mq_setattr can change it
	mq_flags int
	// # messages stored in a queue. 0 == N/A
	mq_maxmsg int
	// # messages currently on the given queue
	mq_curmsgs int
	// # of processes waiting to send a message
	mq_sendwait int
	// # of process waiting to recieve a message
	mq_recvwait int
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
