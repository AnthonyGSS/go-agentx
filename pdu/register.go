package pdu

import (
	"github.com/juju/errgo"
	"github.com/posteo/go-agentx/marshaler"
)

// Register defines the pdu register packet.
type Register struct {
	Timeout Timeout
	Subtree ObjectIdentifier
}

// Type returns the pdu packet type.
func (r *Register) Type() Type {
	return TypeRegister
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (r *Register) MarshalBinary() ([]byte, error) {
	combined := marshaler.NewMulti(&r.Timeout, &r.Subtree)

	combinedBytes, err := combined.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return combinedBytes, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Register) UnmarshalBinary(data []byte) error {
	return nil
}