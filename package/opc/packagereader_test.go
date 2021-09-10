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

}
