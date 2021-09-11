package opc

import (
	"github.com/juandemanjon/opc/package/opc/schema/sso/pkg/metadata/_2006/core_properties"
)

type CoreProperties struct {
	core_properties.CT_CoreProperties
}

func NewCoreProperties() *CoreProperties {
	return &CoreProperties{}
}
