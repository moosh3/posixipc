package posixipc

import (
	"sync"
)

// Default size of the virtualized heap and stack created by
// the IPC for use with the mq, semaphore, and shared memory.
const (
	dStackSize = 64
	dHeapSize  = 256
)

// ipc is a POSIX-styled <> heavily inspired by the Oracle documentation
// located at https://docs.oracle.com/cd/E19455-01/806-4750/6jdqdfltf/index.html.
// ipc keeps tracker of the number of registers, and creates a virtualize
// stack and heap to use for allocating/deallocating memory for processes/CPU usage.
type ipc struct {
	sync.RWMutex
	isem
	imq
	imem

	registers int64

	stack []byte // 64kb
	heap  []byte // 256kb
}

// NewIPC creates a new IPC with a fixed sized virtualized stack and
// heap, handling one register at a time.
func NewIPC() *ipc {
	return &ipc{
		registers: 1,
		stack:     make([]byte, dStackSize),
		heap:      make([]byte, dHeapSize),
	}
}

// Semaphore
type Semaphore interface {
	// Open establishes a connection between a name semaphore
	// and a process, LWP or thread
	Open(name string, oflag int) (*sem, error) // sem_open
	// Init is used to initialize the unnamed semaphores referred
	// too by *sem. The value of the initialized semaphore is `value`.
	// If `pshared` has a non-zero value, then the semaphore is shared
	// between processes; in this case any process that can
	// access the semaphore `sem` for performing
	// Wait(), TryWait(), Post(), and Destroy() operations
	Init(sem_t *sem, pshared int, value int) error // sem_init
	// Close is used to indicate that the calling process is finished
	// used the named semaphore indicated by `sem`. The sem_close() function
	// deallocates any system resources allocated by the system for use
	// by this process for this semaphore.
	Close(sem_t *sem) error // sem_close
	// Unlink removes the semaphore named by the string `name`. If the
	// semaphore named by `name` is currently referenced by other
	// processes, then sem_unlink has no effect on the state of the semaphore.
	Unlink(name *string) error // sem_unlink
	// Destory is used to destroy the unnamed semaphore indicated by `sem`.
	// Only a semaphore that was create using sem_init() may be destroyed via sem_destroy().
	Destroy(sem_t *sem) error // sem_destroy
	// GetValue updates the location referenced by the `*sval` argument
	// to have the value of the semaphore referenced by `sem`
	// without affecting the state of the semaphore.
	// The value `sval` may be 0 or positive. If `sval` is 0, there
	// may be other processes (or LWPs or threads) waiting for the
	// semaphore; if `sval` is positive, no processed is waiting.
	GetValue(sem_t *sem, sval *int) error
	// Wait locks the semaphore referenced by `sem` by performing a semaphore
	// lock operation on that semaphore. If the semaphore value is
	// currently zero, then the calling thread will not return from the call
	// to sem_wait() until it either locks the semaphore or the call is
	// interrupted by a signal. Upon success, the state of the semaphore
	// remains locked until sem_post() is executed successfully
	Wait(sem_t *sem)
	TryWait(sem_t *sem)
	// Post unlocks the semaphore referenced by `sem` by performing
	// a semaphore unlock operation on that semaphore
	Post(*sem) error
}

// MessageQueue is an interface that defines an operator for a given
// message queue, created via a SYS_CALL (O_RDONLY, O_WRONLY, O_RDWR, O_CREAT,
// O_EXCL, and O_NONBLOCK. More information can be found in the Oracle
// documentation: https://docs.oracle.com/cd/E19455-01/806-0630/6j9vkb8dq/index.html
type MessageQueue interface {
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

// SharedMemory is an interface for interacting with memory that is accessed
// by more than one process at any given time; the implementation uses Go's philosphy:
// `Do not communicate by sharing memory; instead, share memory by communicating`
type SharedMemory interface {
	Open(io.File)
	mmap()
}
