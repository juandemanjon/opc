package opc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartAddRelationship(t *testing.T) {
	var part1, part2 *Part
	part1 = &Part{}

	part2 = &Part{}
	_, err := part1.AddRelationship(part2, "", InternalTarget)
	assert.Error(t, err)

	part2 = &Part{Name: "name", Path: "/foo"}
	_, err = part1.AddRelationship(part2, "", InternalTarget)
	assert.Error(t, err)

	rs, err := part1.AddRelationship(part2, "type12", InternalTarget)
	assert.NoError(t, err)
	assert.NotNil(t, rs)

}
