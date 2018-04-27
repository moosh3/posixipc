package posixipc // import github.com/aleccunningham/posixipc

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

func MemoryStats() *os.MemStats {
	var memStats *runtime.MemStats
	return runtime.ReadMemStats(memStats)
}
