package posixipc

const megabyte uint64 = 1024 * 1024

type mem struct {
	total int64
	free  int64

	// swap
	stotal int64
	sfree  int64
}
