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
}

func (obj SharkyObject) Builder() *SharkyObjectBuilder {
	return &SharkyObjectBuilder{}
}

func (builder SharkyObjectBuilder) SetName(name string) SharkyObjectBuilder {
	builder.name = name
	return builder
}

func (builder SharkyObjectBuilder) Build() SharkyObject {
	var obj = SharkyObject{
		Name: builder.name,
	}
	return obj
}
