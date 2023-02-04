package model

import (
	"github.com/voicera/gooseberry/urn"
)

type SharkyDomain struct {
	Name string
	Urn  urn.URN
}

type SharkyDomainBuilder struct {
	name string
	urn  urn.URN
}

func (obj SharkyDomain) Builder() SharkyDomainBuilder {
	return SharkyDomainBuilder{}
}

func (builder SharkyDomainBuilder) SetName(name string) SharkyDomainBuilder {
	builder.name = name
	return builder
}

func (builder SharkyDomainBuilder) SetUrn(urn urn.URN) SharkyDomainBuilder {
	builder.urn = urn
	return builder
}

func (builder SharkyDomainBuilder) Build() SharkyDomain {
	var domain = SharkyDomain{
		Name: builder.name,
		Urn:  builder.urn,
	}
	return domain
}
