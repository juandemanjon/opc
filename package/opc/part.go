package opc

import "io/fs"

type PartContainer struct {
	rsc rsContainer
}

func (p *PartContainer) AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error) {
	return p.rsc.AddRelationship(target, rsType, targetMode)
}

func (p *PartContainer) RelationshipCount() int {
	return p.rsc.RelationshipCount()
}

func (p *PartContainer) RelationshipsByType(type_ string) []*Relationship {
	return p.rsc.RelationshipByType(type_)
}

func (p *PartContainer) readRelationships(file fs.File, partName string) error {
	return p.rsc.decodeRelationships(file, partName)
}

type Part struct {
	PartContainer
	Path string
}

func NewPart(path string) *Part {
	return &Part{
		Path: path,
	}
}
