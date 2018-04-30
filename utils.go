package posixipc

import (
	"runtime"
)

//NumCPU returns the number of CPUs
func NumCPU() int {
	return runtime.NumCPU()
}

// MaxProcs sets GOMAXPROCS to the number
// of CPUs reported by NumCPU() and returns the number
func MaxProcs() int {
	num := NumCpu()
	return runtime.GOMAXPROCS(num)
}

func MemoryStats() *runtime.MemStats {
	memStats := new(runtime.MemStats)
	return runtime.ReadMemStats(memStats)
}

// SlicePtrFromStrings converts a slice of strings to a slice of
// pointers to NUL-terminated byte arrays. If any string contains
// a NUL byte, it returns (nil, EINVAL).
func SlicePtrFromStrings(ss []string) ([]*byte, error) {
	var err error
	bb := make([]*byte, len(ss)+1)
	for i := 0; i < len(ss); i++ {
		bb[i], err = BytePtrFromString(ss[i])
		if err != nil {
			return nil, err
		}
	}
	bb[len(ss)] = nil
	return bb, nil
}
