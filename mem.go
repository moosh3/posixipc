package posixipc

import (
	"io"
)

type mem struct {
	total int64
	free  int64

	// swap
	stotal int64
	sfree  int64
}

type imem interface {
	Open(io.File)
	mmap()
}
