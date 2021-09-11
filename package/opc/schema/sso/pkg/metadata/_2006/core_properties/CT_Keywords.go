package core_properties

import "encoding/xml"

const (
	w3_namespace = "http://www.w3.org/XML/1998/namespace"
)

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
	Tags     *string
}

func NewCT_Keywords() *CT_Keywords {
	return &CT_Keywords{}
}

func (m *CT_Keywords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Space == w3_namespace && attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
			continue
		}
	}
outerCT_Keywords:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: CP, Local: "value"}:
				ct := NewCT_Keyword()
				if err := d.DecodeElement(ct, &el); err != nil {
					return err
				}
				m.Value = append(m.Value, ct)
			default:
			}
		case xml.EndElement:
			break outerCT_Keywords
		case xml.CharData:
			if m.Tags == nil {
				m.Tags = new(string)
			}
			value := *m.Tags + string(el)
			m.Tags = &value
		}
	}
	return nil
}
