package opc

import (
	"archive/zip"
)

type Package struct {
	rsc rsContainer
}

func NewPackage() *Package {
	return &Package{}
}

func (p *Package) AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error) {
	return p.rsc.AddRelationship(target, rsType, targetMode)
}

func (p *Package) RelationshipCount() int {
	return p.rsc.RelationshipCount()
}

func (p *Package) RelationshipByType(type_ string) []*Relationship {
	return p.rsc.RelationshipByType(type_)
}

func (p *Package) relationshipsByType(type_ string) []*Relationship {
	return p.rsc.RelationshipByType(type_)
}

const (
	URI_PackageRels = "_rels/.rels"
)

func (p *Package) readRelationships(r *zip.ReadCloser) error {
	file, err := r.Open(URI_PackageRels)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.rsc.decodeRelationships(file, URI_PackageRels)
}
