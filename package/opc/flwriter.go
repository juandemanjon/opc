package opc

import "io"

type PhysicalLayerWriter interface {
	Open() error
	GetWriter(name string) (io.Writer, error)
	Close() error
}
