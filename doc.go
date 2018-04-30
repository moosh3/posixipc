// Package posixipc is a loose implementation of the POSIX IPC interface defined by
// Oracle (https://docs.oracle.com/cd/E19455-01/806-4750/6jdqdfltf/index.html), but modified for Go.
// For some system calls, cgo is used in place of calling os/syscall for speed.
package posixipc
