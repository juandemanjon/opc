package opc

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
)

const (
	URI_PackageRels  = "_rels/.rels"
	URI_ContentTypes = "[Content_Types].xml"
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

	file, err := r.Open(URI_PackageRels)
	if err != nil {
		return err
	}
	defer file.Close()
	p.readRelationships(file, URI_PackageRels)

	file, err = r.Open(URI_ContentTypes)
	if err != nil {
		return err
	}
	defer file.Close()
	p.readContentTypes(file)

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
