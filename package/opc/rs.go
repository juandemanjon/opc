package opc

type TargetMode int

const (
	InternalTarget TargetMode = iota
	ExternalTarget
)

type Relationship struct {
	Target           *Part
	RelationshipType string
	TargetMode       TargetMode
	Id               string
}
