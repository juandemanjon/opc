package opc

import "fmt"

type RelationShipContainer interface {
	AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error)
	// DeleteRelationship()
	// DeleteRelationshipsByTarget()
}

type rsContainer struct {
	rss []*Relationship
}

func (rsc *rsContainer) AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error) {
	if target == nil {
		return nil, fmt.Errorf("Target part cannot be nil")
	}
	if len(rsType) == 0 {
		return nil, fmt.Errorf("Relationship Type cannot be empty")
	}
	rs := &Relationship{Target: target, RelationshipType: rsType, TargetMode: targetMode}
	rsc.rss = append(rsc.rss, rs)
	return rs, nil
}
