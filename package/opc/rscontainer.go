package opc

import (
	"encoding/xml"
	"fmt"
	"io"
)

type RelationshipContainer struct {
	rss []*Relationship
}

func (rsc *RelationshipContainer) AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error) {
	if target == nil {
		return nil, fmt.Errorf("target part cannot be nil")
	}
	if len(rsType) == 0 {
		return nil, fmt.Errorf("relationship Type cannot be empty")
	}
	rs := &Relationship{TargetPart: target, Type: rsType, TargetMode: targetMode}
	if target != nil {
		rs.TargetURI = target.Path
	}
	rsc.rss = append(rsc.rss, rs)
	return rs, nil
}

func (rsc *RelationshipContainer) RelationshipCount() int {
	return len(rsc.rss)
}

func (rsc *RelationshipContainer) RelationshipsByType(type_ string) []*Relationship {
	var rs []*Relationship
	for _, r := range rsc.rss {
		if r.Type == type_ {
			rs = append(rs, r)
		}
	}
	return rs
}

type relationshipsXML struct {
	XMLName xml.Name           `xml:"Relationships"`
	XML     string             `xml:"xmlns,attr"`
	RelsXML []*relationshipXML `xml:"Relationship"`
}

type relationshipXML struct {
	ID        string `xml:"Id,attr"`
	RelType   string `xml:"Type,attr"`
	TargetURI string `xml:"Target,attr"`
	Mode      string `xml:"TargetMode,attr,omitempty"`
}

func (rsc *RelationshipContainer) decodeRelationships(r io.Reader, partName string) error {
	relDecode := new(relationshipsXML)
	if err := xml.NewDecoder(r).Decode(relDecode); err != nil {
		return fmt.Errorf("opc: %s: cannot be decoded: %v", partName, err)
	}
	rel := make([]*Relationship, len(relDecode.RelsXML))
	for i, rl := range relDecode.RelsXML {
		newRel := &Relationship{ID: rl.ID, TargetURI: rl.TargetURI, Type: rl.RelType}
		if rl.Mode == "" || rl.Mode == "Internal" {
			newRel.TargetMode = InternalTarget
		} else {
			newRel.TargetMode = ExternalTarget
		}
		rel[i] = newRel
	}
	rsc.rss = rel
	return nil
}
