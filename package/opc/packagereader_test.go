package opc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageReader(t *testing.T) {
	reader := NewPackageReader("testfiles/empty.docx")

	err := reader.Open()
	assert.NoError(t, err)
	defer reader.Close()

	assert.Equal(t, 3, reader.Package.RelationshipCount())

	core, err := reader.GetCoreProperties()
	assert.NoError(t, err)
	assert.NotNil(t, core)
	assert.Equal(t, "Manjon, Juande", string(core.Creator.Data))
	assert.NotNil(t, core.Keywords)
	assert.NotNil(t, core.Keywords.Tags)
	assert.Equal(t, "Tags1; Tags2", *core.Keywords.Tags)

	p := reader.Package
	reader = nil

	ct := p.ct
	assert.Equal(t, "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml", ct.Override[0].ContentType)
	assert.Equal(t, "/word/document.xml", ct.Override[0].PartName)

}
