package thing

import "github.com/dpb587/go-schemaorg"

var (
	// URL of a reference Web page that unambiguously indicates the item's identity.
	// E.g. the URL of the item's Wikipedia page, Wikidata entry, or official
	// website.
	SameAs = schemaorg.NewProperty("sameAs")

	// URL of the item.
	Url = schemaorg.NewProperty("url")

	// An image of the item. This can be a <a class="localLink"
	// href="http://schema.org/URL">URL</a> or a fully described <a
	// class="localLink" href="http://schema.org/ImageObject">ImageObject</a>.
	Image = schemaorg.NewProperty("image")

	// An additional type for the item, typically used for adding more specific
	// types from external vocabularies in microdata syntax. This is a relationship
	// between something and a class that the thing is in. In RDFa syntax, it is
	// better to use the native RDFa syntax - the 'typeof' attribute - for multiple
	// types. Schema.org tools may have only weaker understanding of extra types, in
	// particular those defined externally.
	AdditionalType = schemaorg.NewProperty("additionalType")

	// The name of the item.
	Name = schemaorg.NewProperty("name")

	// The identifier property represents any kind of identifier for any kind of <a
	// class="localLink" href="http://schema.org/Thing">Thing</a>, such as ISBNs,
	// GTIN codes, UUIDs etc. Schema.org provides dedicated properties for
	// representing many of these, either as textual strings or as URL (URI) links.
	// See <a href="/docs/datamodel.html#identifierBg">background notes</a> for more
	// details.
	Identifier = schemaorg.NewProperty("identifier")

	// Indicates a potential Action, which describes an idealized action in which
	// this thing would play an 'object' role.
	PotentialAction = schemaorg.NewProperty("potentialAction")

	// Indicates a page (or other CreativeWork) for which this thing is the main
	// entity being described. See <a
	// href="/docs/datamodel.html#mainEntityBackground">background notes</a> for
	// details.
	MainEntityOfPage = schemaorg.NewProperty("mainEntityOfPage")

	// A description of the item.
	Description = schemaorg.NewProperty("description")

	// A sub property of description. A short description of the item used to
	// disambiguate from other, similar items. Information from other properties (in
	// particular, name) may be necessary for the description to be useful for
	// disambiguation.
	DisambiguatingDescription = schemaorg.NewProperty("disambiguatingDescription")

	// An alias for the item.
	AlternateName = schemaorg.NewProperty("alternateName")
)
