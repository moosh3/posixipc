package posixipc

type lwp struct {
	pid int
}

type ilwp interface{}

var _ ilwp = &lwp
