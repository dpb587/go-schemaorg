package organization

import "github.com/dpb587/go-schemaorg"

var (
	// The geographic area where the service is provided.
	ServiceArea = schemaorg.NewProperty("serviceArea")

	// A person or organization that supports (sponsors) something through some kind
	// of financial contribution.
	Funder = schemaorg.NewProperty("funder")

	// The geographic area where a service or offered item is provided.
	AreaServed = schemaorg.NewProperty("areaServed")

	// An Organization (or ProgramMembership) to which this Person or Organization
	// belongs.
	MemberOf = schemaorg.NewProperty("memberOf")

	// Upcoming or past events associated with this place or organization.
	Events = schemaorg.NewProperty("events")

	// A relationship between two organizations where the first includes the second,
	// e.g., as a subsidiary. See also: the more specific 'department' property.
	SubOrganization = schemaorg.NewProperty("subOrganization")

	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	HasOfferCatalog = schemaorg.NewProperty("hasOfferCatalog")

	// The <a href="http://www.gs1.org/gln">Global Location Number</a> (GLN,
	// sometimes also referred to as International Location Number or ILN) of the
	// respective organization, person, or place. The GLN is a 13-digit number used
	// to identify parties and physical locations.
	GlobalLocationNumber = schemaorg.NewProperty("globalLocationNumber")

	// Review of the item.
	Reviews = schemaorg.NewProperty("reviews")

	// A member of this organization.
	Members = schemaorg.NewProperty("members")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// The Dun &amp; Bradstreet DUNS number for identifying an organization or
	// business person.
	Duns = schemaorg.NewProperty("duns")

	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or
	// the CIF/NIF in Spain.
	TaxID = schemaorg.NewProperty("taxID")

	// An award won by or for this item.
	Award = schemaorg.NewProperty("award")

	// A pointer to products or services offered by the organization or person.
	MakesOffer = schemaorg.NewProperty("makesOffer")

	// A contact point for a person or organization.
	ContactPoints = schemaorg.NewProperty("contactPoints")

	// Awards won by or for this item.
	Awards = schemaorg.NewProperty("awards")

	// A pointer to products or services sought by the organization or person
	// (demand).
	Seeks = schemaorg.NewProperty("seeks")

	// A member of an Organization or a ProgramMembership. Organizations can be
	// members of organizations; ProgramMembership is typically for individuals.
	Member = schemaorg.NewProperty("member")

	// A person who founded this organization.
	Founders = schemaorg.NewProperty("founders")

	// Alumni of an organization.
	Alumni = schemaorg.NewProperty("alumni")

	// The date that this organization was dissolved.
	DissolutionDate = schemaorg.NewProperty("dissolutionDate")

	// Physical address of the item.
	Address = schemaorg.NewProperty("address")

	// An associated logo.
	Logo = schemaorg.NewProperty("logo")

	// People working for this organization.
	Employees = schemaorg.NewProperty("employees")

	// The telephone number.
	Telephone = schemaorg.NewProperty("telephone")

	// Email address.
	Email = schemaorg.NewProperty("email")

	// A relationship between an organization and a department of that organization,
	// also described as an organization (allowing different urls, logos, opening
	// hours). For example: a store with a pharmacy, or a bakery with a cafe.
	Department = schemaorg.NewProperty("department")

	// A contact point for a person or organization.
	ContactPoint = schemaorg.NewProperty("contactPoint")

	// The larger organization that this organization is a <a class="localLink"
	// href="http://schema.org/subOrganization">subOrganization</a> of, if any.
	ParentOrganization = schemaorg.NewProperty("parentOrganization")

	// The official name of the organization, e.g. the registered company name.
	LegalName = schemaorg.NewProperty("legalName")

	// The date that this organization was founded.
	FoundingDate = schemaorg.NewProperty("foundingDate")

	// Someone working for this organization.
	Employee = schemaorg.NewProperty("employee")

	// The number of employees in an organization e.g. business.
	NumberOfEmployees = schemaorg.NewProperty("numberOfEmployees")

	// The North American Industry Classification System (NAICS) code for a
	// particular organization or business person.
	Naics = schemaorg.NewProperty("naics")

	// Points-of-Sales operated by the organization or person.
	HasPOS = schemaorg.NewProperty("hasPOS")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// The place where the Organization was founded.
	FoundingLocation = schemaorg.NewProperty("foundingLocation")

	// Products owned by the organization or person.
	Owns = schemaorg.NewProperty("owns")

	// Upcoming or past event associated with this place, organization, or action.
	Event = schemaorg.NewProperty("event")

	// A person who founded this organization.
	Founder = schemaorg.NewProperty("founder")

	// The publishingPrinciples property indicates (typically via <a
	// class="localLink" href="http://schema.org/URL">URL</a>) a document describing
	// the editorial principles of an <a class="localLink"
	// href="http://schema.org/Organization">Organization</a> (or individual e.g. a
	// <a class="localLink" href="http://schema.org/Person">Person</a> writing a
	// blog) that relate to their activities as a publisher, e.g. ethics or
	// diversity policies. When applied to a <a class="localLink"
	// href="http://schema.org/CreativeWork">CreativeWork</a> (e.g. <a
	// class="localLink" href="http://schema.org/NewsArticle">NewsArticle</a>) the
	// principles are those of the party primarily responsible for the creation of
	// the <a class="localLink"
	// href="http://schema.org/CreativeWork">CreativeWork</a>.</p>
	// 
	// <p>While such policies are most typically expressed in natural language,
	// sometimes related information (e.g. indicating a <a class="localLink"
	// href="http://schema.org/funder">funder</a>) can be expressed using schema.org
	// terminology.
	PublishingPrinciples = schemaorg.NewProperty("publishingPrinciples")

	// A person or organization that supports a thing through a pledge, promise, or
	// financial contribution. e.g. a sponsor of a Medical Study or a corporate
	// sponsor of an event.
	Sponsor = schemaorg.NewProperty("sponsor")

	// The International Standard of Industrial Classification of All Economic
	// Activities (ISIC), Revision 4 code for a particular organization, business
	// person, or place.
	IsicV4 = schemaorg.NewProperty("isicV4")

	// The location of for example where the event is happening, an organization is
	// located, or where an action takes place.
	Location = schemaorg.NewProperty("location")

	// The brand(s) associated with a product or service, or the brand(s) maintained
	// by an organization or business person.
	Brand = schemaorg.NewProperty("brand")

	// The Value-added Tax ID of the organization or person.
	VatID = schemaorg.NewProperty("vatID")

	// An organization identifier that uniquely identifies a legal entity as defined
	// in ISO 17442.
	LeiCode = schemaorg.NewProperty("leiCode")

	// The fax number.
	FaxNumber = schemaorg.NewProperty("faxNumber")
)
