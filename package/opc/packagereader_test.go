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

	file, err := reader.OpenEntry("word/webSettings.xml")
	assert.NoError(t, err)
	expected := 80
	buffer := make([]byte, expected)
	read, err := file.Read(buffer)
	assert.NoError(t, err)
	assert.Equal(t, expected, read)
	assert.Contains(t, string(buffer), "w:webSettings")

	p := reader.Package
	reader = nil

	ct := p.ct
	assert.Equal(t, "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml", ct.Override[0].ContentType)
	assert.Equal(t, "/word/document.xml", ct.Override[0].PartName)

	rs := p.RelationshipsByType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument")
	assert.NotEmpty(t, rs)
	assert.Equal(t, "word/document.xml", rs[0].TargetURI)
	document := rs[0].TargetPart
	assert.Equal(t, "word/document.xml", document.Path)

	rs = document.RelationshipsByType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable")
	assert.NotEmpty(t, rs)
	assert.Equal(t, "fontTable.xml", rs[0].TargetURI)
	fontTable := rs[0].TargetPart
	assert.Equal(t, "word/fontTable.xml", fontTable.Path)

}
