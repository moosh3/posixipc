package posixipc

import (
	"sync"
)

// SharedMem provides a segment of memory that can be shared
// between goroutines, passed via a mutex locker
type SharedMem struct {
	sync.Mutex

	frames int64
	max    []byte

	free int64
}
