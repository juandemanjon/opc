package opc

import (
	"encoding/xml"
	"io/fs"

	"github.com/juandemanjon/opc/package/opc/schema/sso/pkg/metadata/_2006/content_types"
)

type Package struct {
	PartContainer
	ct *content_types.CT_Types
}

func NewPackage() *Package {
	return &Package{
		ct: content_types.NewCT_Types(),
	}
}

func (p *Package) readContentTypes(file fs.File) error {
	p.ct = content_types.NewCT_Types()
	err := xml.NewDecoder(file).Decode(p.ct)
	return err
}
