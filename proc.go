package posixipc

import "os"

// Proc is a wrapper around a generic process
type Proc struct {
	pid  int
	proc *os.Process
}

// ProcSlice is a slice of Lwps that can be sorted by pid
type ProcSlice []Proc

// Len returns the length of the ProcSlice
func (l ProcSlice) Len() int { return len(l) }

// Less returns either True or False, in regards to the size of a
// Lwp instance in a ProcSlice
func (l ProcSlice) Less(i, j int) bool { return l[i] < l[j] }

// Swap swaps the order of two Lwp objects in a ProcSlice
func (l ProcSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
