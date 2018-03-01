package main

import (
	"encoding/json"
	"fmt"

	"github.com/dpb587/go-schemaorg/datatypes/text"
	"github.com/dpb587/go-schemaorg/datatypes/url"
	"github.com/dpb587/go-schemaorg/things/country"
	"github.com/dpb587/go-schemaorg/things/organization"
	"github.com/dpb587/go-schemaorg/things/person"
	"github.com/dpb587/go-schemaorg/things/place"
	"github.com/dpb587/go-schemaorg/things/postaladdress"
	"github.com/dpb587/go-schemaorg/things/thing"
)

func main() {
	subject := person.New().
		Add(thing.Identifier, url.New("https://dpb587.me/#person")).
		Add(thing.Url, url.New("https://dpb587.me/")).
		Add(person.FamilyName, text.New("Berger")).
		Add(person.GivenName, text.New("Danny")).
		Add(person.HomeLocation, place.New().
			Add(place.Address, postaladdress.New().
				Add(postaladdress.AddressCountry, country.New().
					Add(thing.Name, text.New("US")),
				).
				Add(postaladdress.AddressRegion, text.New("California")).
				Add(postaladdress.AddressLocality, text.New("Oakland")),
			),
		).
		Add(person.JobTitle, text.New("Software Engineer")).
		Add(person.Nationality, country.New().
			Add(thing.Name, text.New("US")),
		).
		Add(person.WorksFor, organization.New().
			Add(thing.Name, text.New("The Loopy Ewe")).
			Add(thing.Identifier, url.New("https://www.theloopyewe.com/#organization")),
		).
		Add(person.WorksFor, organization.New().
			Add(thing.Name, text.New("Pivotal Software, Inc.")).
			Add(thing.Identifier, url.New("https://pivotal.io/#organization")),
		)

	bytes, err := json.MarshalIndent(subject, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bytes)
}
