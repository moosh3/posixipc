package posixipc

type sem struct {
	name string
}

type isemaphore interface {
	// Open establishes a connection between a name semaphore
	// and a process, LWP or thread
	Open(name string, of oflag) (*sem, error) // sem_open
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
