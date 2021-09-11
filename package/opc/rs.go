package opc

type TargetMode int

const (
	InternalTarget TargetMode = iota
	ExternalTarget
)

type Relationship struct {
	TargetPart *Part
	TargetURI  string
	Type       string
	TargetMode TargetMode
	ID         string
}
