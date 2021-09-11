package core_properties

import (
	"encoding/xml"
	"time"

	"github.com/juandemanjon/opc/package/opc/schema"
)

const (
	CP      = "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"
	dc      = "http://purl.org/dc/elements/1.1/"
	dcterms = "http://purl.org/dc/terms/"
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

outerCT_CoreProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: CP, Local: "category"}:
				m.Category = new(string)
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "contentStatus"}:
				m.ContentStatus = new(string)
				if err := d.DecodeElement(m.ContentStatus, &el); err != nil {
					return err
				}
			case xml.Name{Space: dcterms, Local: "created"}:
				m.Created = new(schema.XSDAny)
				if err := d.DecodeElement(m.Created, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "creator"}:
				m.Creator = new(schema.XSDAny)
				if err := d.DecodeElement(m.Creator, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "description"}:
				m.Description = new(schema.XSDAny)
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "identifier"}:
				m.Identifier = new(schema.XSDAny)
				if err := d.DecodeElement(m.Identifier, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "keywords"}:
				m.Keywords = NewCT_Keywords()
				if err := d.DecodeElement(m.Keywords, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "language"}:
				m.Language = new(schema.XSDAny)
				if err := d.DecodeElement(m.Language, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "lastModifiedBy"}:
				m.LastModifiedBy = new(string)
				if err := d.DecodeElement(m.LastModifiedBy, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "lastPrinted"}:
				m.LastPrinted = new(time.Time)
				if err := d.DecodeElement(m.LastPrinted, &el); err != nil {
					return err
				}
			case xml.Name{Space: dcterms, Local: "modified"}:
				m.Modified = new(schema.XSDAny)
				if err := d.DecodeElement(m.Modified, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "revision"}:
				m.Revision = new(string)
				if err := d.DecodeElement(m.Revision, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "subject"}:
				m.Subject = new(schema.XSDAny)
				if err := d.DecodeElement(m.Subject, &el); err != nil {
					return err
				}
			case xml.Name{Space: dc, Local: "title"}:
				m.Title = new(schema.XSDAny)
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: CP, Local: "version"}:
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
