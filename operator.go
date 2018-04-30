package posixipc

import (
	"sync"
	"time"
)

// Default size of the virtualized heap created by
// the IPC for use with the mq, semaphore, and shared memory.
const (
	dHeapSize = 256
)

// Operator is a POSIX-styled <> heavily inspired by the Oracle documentation
// located at https://docs.oracle.com/cd/E19455-01/806-4750/6jdqdfltf/index.html.
// ipc keeps tracker of the number of registers, and creates a virtualize
// stack and heap to use for allocating/deallocating memory for processes/CPU usage.
type Operator struct {
	lock sync.RWMutex

	semaphore    SemaphoreIPC
	messageQueue MessageQueueIPC
	memory       SharedMemoryIPC
}

// NewOperator creates a new Operator with an initialize Semaphore, MessageQueue and Memory
func NewOperator(sem *Semaphore, mq *MessageQueue, mem *Memory) *Operator {
	return &Operator{
		semaphore:    sem,
		messageQueue: mq,
		memory:       mem,
	}
}

func (o *Operator) Start(msg <-chan struct{}) {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-done:
		case ticker.C:
			go parse(msg)
		}
	}
}

func parse(msg []byte) {
	/* ... */
}
