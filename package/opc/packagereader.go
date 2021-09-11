package opc

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
)

type PackageReader struct {
	Filename string
	r        *zip.ReadCloser
	Package  *Package
}

func NewPackageReader(filename string) *PackageReader {
	return &PackageReader{Filename: filename}
}

func (pr *PackageReader) Open() error {
	if pr.r != nil {
		err := pr.Close()
		if err != nil {
			return err
		}
	}
	r, err := zip.OpenReader(pr.Filename)
	if err != nil {
		return err
	}
	p := NewPackage()
	p.readRelationships(r)
	pr.Package = p
	pr.r = r
	return nil
}

func (pr *PackageReader) Close() error {
	if pr.r != nil {
		err := pr.r.Close()
		if err != nil {
			return err
		}
		pr.r = nil
	}
	return nil
}

const (
	Relationship_CoreProperties = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
)

func (pr *PackageReader) GetCoreProperties() (*CoreProperties, error) {
	rs := pr.Package.relationshipsByType(Relationship_CoreProperties)
	if len(rs) == 0 {
		return nil, fmt.Errorf("cannot find core properties relationship with type '%s'", Relationship_CoreProperties)
	}

	file, err := pr.r.Open(rs[0].TargetURI)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	core := NewCoreProperties()
	err = xml.NewDecoder(file).Decode(core)

	return core, err
}
