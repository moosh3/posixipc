# posixpic [![GoDoc](https://godoc.org/github.com/aleccunningham/posixipc?status.svg)](https://godoc.org/github.com/aleccunningham/posixipc)

POSIX pic interfaces implemented in Go

Modified to make use of Go primitives, as POSIC IPC interfaces are multithread safe, and can take advantage of the use of goroutines and channels. `cgo` is used for embedding C where appropriate.

The following is the POSIX IPC interface defined by Oracle. This is not a 1:1 to the API interface exposed by this package. See [types.go](https://github.com/aleccunningham/posixipc/blob/master/types.go) for the Go API equivilant.

### POSIX Messages
```Go
type mq interface {
    mq_open(3RT)    error
    mq_close(3RT)   error
    mq_unlink(3RT)  error
    mq_send(3RT)    error
    mq_recieve(3RT) error
    mq_notify(3RT)  error
    mq_setattr(3RT) error
    mq_getattr(3RT) error
}
```

### POSIX Semaphores
POSIX Semaphores are much lighter weight than System V semaphores. A POSIC semaphore structure defines a single semaphore, not an array of up to twenty five semaphores.
```Go
type sem interface {
    sem_open(3RT)       error
    sem_init(3RT)       error
    sem_close(3RT)      error
    sem_unlink(3RT)     error
    sem_destroy(3RT)    error
    sem_getvalue(3RT)   error
    sem_wait(3RT)       error
    sem_trywait(3RT)    error
    sem_post(3RT)       error
}
```

### POSIX Shared Memory
POSIX shared memory is a variation of mapped memory; the major differences are using `shm_open(3RT)` to open the shared memory object, and using `shm_unlink(3RT)` to close and delete the object (`close(2)` does not remove the object). 

As Oracle does not provide an explicit interface, the following struct and interface will represent shared memory.
```Go
type shmem interface {
    shm_open(3RT)   error
    shm_close(3RT)  error
    shm_unlink(3RT) error
}

type mem struct {
    buf []byte
    lock sync.Mutex
}
```
