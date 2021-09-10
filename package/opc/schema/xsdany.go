package schema

import "encoding/xml"

type XSDAny struct {
	XMLName xml.Name
	Attrs   []xml.Attr
	Data    []byte
	Nodes   []*XSDAny
}

type any struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Nodes   []*any     `xml:",any"`
	Data    []byte     `xml:",chardata"`
}

func dd(a *any) {
	for _, n := range a.Nodes {
		dd(n)
	}
}

func convertToXNodes(an []*any) []*XSDAny {
	ret := []*XSDAny{}
	for _, a := range an {
		x := &XSDAny{}
		x.XMLName = a.XMLName
		x.Attrs = a.Attrs
		x.Data = a.Data
		x.Nodes = convertToXNodes(a.Nodes)
		ret = append(ret, x)
	}
	return ret
}

func (x *XSDAny) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a := any{}
	if err := d.DecodeElement(&a, &start); err != nil {
		return err
	}
	dd(&a)
	x.XMLName = a.XMLName
	x.Attrs = a.Attrs
	x.Data = a.Data
	x.Nodes = convertToXNodes(a.Nodes)
	return nil
}
