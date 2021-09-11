package content_types

import (
	"encoding/xml"
	"fmt"
)

type CT_Override struct {
	ContentType string
	PartName    string
}

func NewCT_Override() *CT_Override {
	return &CT_Override{
		ContentType: "application/xml",
	}
}

func (m *CT_Override) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentType = parsed
			continue
		}
		if attr.Name.Local == "PartName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PartName = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Override: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
