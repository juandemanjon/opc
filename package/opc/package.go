package opc

import (
	"encoding/xml"
	"io/fs"

	"github.com/juandemanjon/opc/package/opc/schema/sso/pkg/metadata/_2006/content_types"
)

type Package struct {
	rsc rsContainer
	ct  *content_types.CT_Types
}

func NewPackage() *Package {
	return &Package{
		ct: content_types.NewCT_Types(),
	}
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

func (p *Package) readRelationships(file fs.File, partName string) error {
	return p.rsc.decodeRelationships(file, partName)
}

func (p *Package) readContentTypes(file fs.File) error {
	p.ct = content_types.NewCT_Types()
	err := xml.NewDecoder(file).Decode(p.ct)
	return err
}
