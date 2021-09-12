package opc

type Part struct {
	RelationshipContainer
	Path string
}

func NewPart(path string) *Part {
	return &Part{
		Path: path,
	}
}
