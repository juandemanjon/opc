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

	var core *CoreProperties

	decoder := xml.NewDecoder(file)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch tt := t.(type) {
		case xml.ProcInst:
			fmt.Printf("ProcInst %v", tt)
		case xml.StartElement:
			switch tt.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "coreProperties"}:
				core = &CoreProperties{} //  schema.NewCT_CoreProperties()
				core.UnmarshalXML(decoder, tt)
				fmt.Printf("CP %v", core)
			default:
				fmt.Printf("StartElement %v", tt)
			}
		case xml.CharData:
			fmt.Printf("CharData %v", tt)
		case xml.EndElement:
			fmt.Printf("EndElement %v", tt)
		default:
			fmt.Printf("default %v", tt)
		}
	}

	return core, nil
}
