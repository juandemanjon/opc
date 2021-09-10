package opc

type Part struct {
	Path string
	Name string
	rsc  rsContainer
}

func (p *Part) AddRelationship(target *Part, rsType string, targetMode TargetMode) (*Relationship, error) {
	return p.rsc.AddRelationship(target, rsType, targetMode)
}

func (p *Part) RelationshipCount() int {
	return p.rsc.RelationshipCount()
}

func (p *Part) RelationshipByType(type_ string) []*Relationship {
	return p.rsc.RelationshipByType(type_)
}
