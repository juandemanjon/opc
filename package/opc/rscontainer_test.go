package opc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelationshipIterator(t *testing.T) {
	reader := NewPackageReader("testfiles/empty.docx")
	err := reader.Open()
	assert.NoError(t, err)
	defer reader.Close()

	p := reader.Package

	it := p.Relationships()
	got := []string{}
	expected := []string{"docProps/app.xml", "docProps/core.xml", "word/document.xml"}
	for {
		if !it.HasNext() {
			break
		}
		got = append(got, it.GetNext().TargetURI)
	}
	assert.Equal(t, expected, got)

	got = []string{}
	parts := p.InternalParts()
	for _, part := range parts {
		got = append(got, part.Path)
	}
	assert.Equal(t, expected, got)
}
