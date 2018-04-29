package posixipc

// Lwp is a wrapper around a generic process
type Lwp struct {
	pid int
}

// LwpSlice is a slice of Lwps that can be sorted
// by name
type LwpSlice []Lwp

// Len returns the length of the LwpSlice
func (l LwpSlice) Len() int { return len(l) }

// Less returns either True or False, in regards to the size of a
// Lwp instance in a LwpSlice
func (l LwpSlice) Less(i, j int) bool { return p[i] < p[j] }

// Swap swaps the order of two Lwp objects in a LwpSlice
func (l LwpSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
