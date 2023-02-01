package model

import (
	"github.com/voicera/gooseberry/urn"
)

type SharkyObject struct {
	Name string
	Urn  urn.URN
}

type SharkyObjectBuilder struct {
	name string
	urn  urn.URN
}

func (obj SharkyObject) Builder() SharkyObjectBuilder {
	return SharkyObjectBuilder{}
}

func (builder SharkyObjectBuilder) SetName(name string) SharkyObjectBuilder {
	builder.name = name
	return builder
}

func (builder SharkyObjectBuilder) SetUrn(urn urn.URN) SharkyObjectBuilder {
	builder.urn = urn
	return builder
}

func (builder SharkyObjectBuilder) Build() SharkyObject {
	var obj = SharkyObject{
		Name: builder.name,
		Urn:  builder.urn,
	}
	return obj
}
