package content_types

import (
	"encoding/xml"
	"fmt"
)

type CT_Default struct {
	Extension   string
	ContentType string
}

func NewCT_Default() *CT_Default {
	return &CT_Default{
		Extension:   "xml",
		ContentType: "application/xml",
	}
}

func (m *CT_Default) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "Extension" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Extension = parsed
			continue
		}
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentType = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Default: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
