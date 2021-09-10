package opc

type TargetMode int

const (
	InternalTarget TargetMode = iota
	ExternalTarget
)

type Relationship struct {
	Target     *Part
	TargetURI  string
	Type       string
	TargetMode TargetMode
	ID         string
}
