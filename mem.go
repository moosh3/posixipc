package posixipc

type mem struct {
	total int64
	free  int64

	// swap
	stotal int64
	sfree  int64
}
