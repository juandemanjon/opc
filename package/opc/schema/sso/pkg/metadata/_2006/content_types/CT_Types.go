package content_types

import "encoding/xml"

const (
	ct = "http://schemas.openxmlformats.org/package/2006/content-types"
)

type CT_Types struct {
	Default  []*CT_Default
	Override []*CT_Override
}

func NewCT_Types() *CT_Types {
	return &CT_Types{}
}

func (m *CT_Types) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
outerTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:

			switch el.Name {
			case xml.Name{Space: ct, Local: "Default"}:
				tmp := NewCT_Default()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Default = append(m.Default, tmp)
			case xml.Name{Space: ct, Local: "Override"}:
				tmp := NewCT_Override()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Override = append(m.Override, tmp)
			default:
			}
		case xml.EndElement:
			break outerTypes
		default:
		}
	}
	return nil
}
