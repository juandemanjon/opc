package core_properties

import (
	"encoding/xml"
	"time"

	"github.com/juandemanjon/opc/package/opc/schema"
)

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *schema.XSDAny
	Creator        *schema.XSDAny
	Description    *schema.XSDAny
	Identifier     *schema.XSDAny
	Keywords       *CT_Keywords
	Language       *schema.XSDAny
	LastModifiedBy *string
	LastPrinted    *time.Time
	Modified       *schema.XSDAny
	Revision       *string
	Subject        *schema.XSDAny
	Title          *schema.XSDAny
	Version        *string
}

func NewCT_CoreProperties() *CT_CoreProperties {
	return &CT_CoreProperties{}
}

func (m *CT_CoreProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
outerCT_CoreProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "category"}:
				m.Category = new(string)
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "contentStatus"}:
				m.ContentStatus = new(string)
				if err := d.DecodeElement(m.ContentStatus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "created"}:
				m.Created = new(schema.XSDAny)
				if err := d.DecodeElement(m.Created, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "creator"}:
				m.Creator = new(schema.XSDAny)
				if err := d.DecodeElement(m.Creator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "description"}:
				m.Description = new(schema.XSDAny)
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "identifier"}:
				m.Identifier = new(schema.XSDAny)
				if err := d.DecodeElement(m.Identifier, &el); err != nil {
					return err
				}
			// case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "keywords"}:
			// 	m.Keywords = NewCT_Keywords()
			// 	if err := d.DecodeElement(m.Keywords, &el); err != nil {
			// 		return err
			// 	}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "language"}:
				m.Language = new(schema.XSDAny)
				if err := d.DecodeElement(m.Language, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastModifiedBy"}:
				m.LastModifiedBy = new(string)
				if err := d.DecodeElement(m.LastModifiedBy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastPrinted"}:
				m.LastPrinted = new(time.Time)
				if err := d.DecodeElement(m.LastPrinted, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "modified"}:
				m.Modified = new(schema.XSDAny)
				if err := d.DecodeElement(m.Modified, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "revision"}:
				m.Revision = new(string)
				if err := d.DecodeElement(m.Revision, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "subject"}:
				m.Subject = new(schema.XSDAny)
				if err := d.DecodeElement(m.Subject, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "title"}:
				m.Title = new(schema.XSDAny)
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "version"}:
				m.Version = new(string)
				if err := d.DecodeElement(m.Version, &el); err != nil {
					return err
				}
			default:
			}
		case xml.EndElement:
			break outerCT_CoreProperties
		case xml.CharData:
		}
	}
	return nil
}
