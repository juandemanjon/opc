package opc

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"path/filepath"
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
	err = p.readRelationships(file, URI_PackageRels)
	if err != nil {
		return err
	}

	file, err = r.Open(URI_ContentTypes)
	if err != nil {
		return err
	}
	defer file.Close()
	p.readContentTypes(file)

	pr.r = r

	pr.readRelationships("", p.rsc.rss)

	pr.Package = p
	return nil
}

func (pr *PackageReader) readRelationships(parent string, rss []*Relationship) {
	for _, rs := range rss {
		if rs.TargetMode == InternalTarget {
			if rs.TargetPart == nil {
				uri := parent + rs.TargetURI
				rs.TargetPart = NewPart(uri)
				dir, filename := filepath.Split(uri)
				rel := dir + "_rels/" + filename + ".rels"
				file, err := pr.r.Open(rel)
				if err == nil {
					rs.TargetPart.readRelationships(file, rs.TargetURI)
					pr.readRelationships(dir, rs.TargetPart.rsc.rss)
				}
			}
		}
	}
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
	rs := pr.Package.RelationshipsByType(Relationship_CoreProperties)
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
