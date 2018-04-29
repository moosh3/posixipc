package posixipc

import (
	"fmt"
	"io"
	"runtime"
	"signal"
)

// SYS calls used in POSIX IPC
const (
	SYS_EXIT_GROUP      = 231
	SYS_EPOLL_WAIT      = 232
	SYS_EPOLL_CTL       = 233
	SYS_TGKILL          = 234
	SYS_UTIMES          = 235
	SYS_VSERVER         = 236
	SYS_MBIND           = 237
	SYS_SET_MEMPOLICY   = 238
	SYS_GET_MEMPOLICY   = 239
	SYS_MQ_OPEN         = 240
	SYS_MQ_UNLINK       = 241
	SYS_MQ_TIMEDSEND    = 242
	SYS_MQ_TIMEDRECEIVE = 243
	SYS_MQ_NOTIFY       = 244
	SYS_MQ_GETSETATTR   = 245
	SYS_KEXEC_LOAD      = 246
	SYS_WAITID          = 247
	SYS_ADD_KEY         = 248
	SYS_REQUEST_KEY     = 249
	SYS_KEYCTL          = 250
)

var (
	units = []string{" bytes", "KB", "MB", "GB", "TB", "PB"}
)

func formatBytes(val uint64) string {
	var i int
	var target uint64
	for i = range units {
		target = 1 << uint(10*(i+1))
		if val < target {
			break
		}
	}
	if i > 0 {
		return fmt.Sprintf("%0.2f%s (%d bytes)", float64(val)/(float64(target)/1024), units[i], val)
	}
	return fmt.Sprintf("%d bytes", val)
}

// HandleSysCall handles syscalls specific to a POSIX IPC implementation.
func HandleSyscall(conn io.ReadWriter, msg []byte) error {
	switch msg[0] {
	case signal.GC:
		runtime.GC()
		_, err := conn.Write([]byte("ok"))
		return err
	case signal.MemStats:
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		fmt.Fprintf(conn, "alloc: %v\n", formatBytes(s.Alloc))
		fmt.Fprintf(conn, "total-alloc: %v\n", formatBytes(s.TotalAlloc))
		fmt.Fprintf(conn, "sys: %v\n", formatBytes(s.Sys))
		fmt.Fprintf(conn, "lookups: %v\n", s.Lookups)
		fmt.Fprintf(conn, "mallocs: %v\n", s.Mallocs)
		fmt.Fprintf(conn, "frees: %v\n", s.Frees)
		fmt.Fprintf(conn, "heap-alloc: %v\n", formatBytes(s.HeapAlloc))
		fmt.Fprintf(conn, "heap-sys: %v\n", formatBytes(s.HeapSys))
		fmt.Fprintf(conn, "heap-idle: %v\n", formatBytes(s.HeapIdle))
		fmt.Fprintf(conn, "heap-in-use: %v\n", formatBytes(s.HeapInuse))
		fmt.Fprintf(conn, "heap-released: %v\n", formatBytes(s.HeapReleased))
		fmt.Fprintf(conn, "heap-objects: %v\n", s.HeapObjects)
		fmt.Fprintf(conn, "stack-in-use: %v\n", formatBytes(s.StackInuse))
		fmt.Fprintf(conn, "stack-sys: %v\n", formatBytes(s.StackSys))
		fmt.Fprintf(conn, "stack-mspan-inuse: %v\n", formatBytes(s.MSpanInuse))
		fmt.Fprintf(conn, "stack-mspan-sys: %v\n", formatBytes(s.MSpanSys))
		fmt.Fprintf(conn, "stack-mcache-inuse: %v\n", formatBytes(s.MCacheInuse))
		fmt.Fprintf(conn, "stack-mcache-sys: %v\n", formatBytes(s.MCacheSys))
		fmt.Fprintf(conn, "other-sys: %v\n", formatBytes(s.OtherSys))
		fmt.Fprintf(conn, "gc-sys: %v\n", formatBytes(s.GCSys))
		fmt.Fprintf(conn, "next-gc: when heap-alloc >= %v\n", formatBytes(s.NextGC))
		lastGC := "-"
		if s.LastGC != 0 {
			lastGC = fmt.Sprint(time.Unix(0, int64(s.LastGC)))
		}
		fmt.Fprintf(conn, "last-gc: %v\n", lastGC)
		fmt.Fprintf(conn, "gc-pause-total: %v\n", time.Duration(s.PauseTotalNs))
		fmt.Fprintf(conn, "gc-pause: %v\n", s.PauseNs[(s.NumGC+255)%256])
		fmt.Fprintf(conn, "num-gc: %v\n", s.NumGC)
		fmt.Fprintf(conn, "enable-gc: %v\n", s.EnableGC)
		fmt.Fprintf(conn, "debug-gc: %v\n", s.DebugGC)
	case signal.Version:
		fmt.Printf(conn, "%v\n", runtime.Version())
	default:
		return nil
	}
}
