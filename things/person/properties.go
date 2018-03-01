package person

import "github.com/dpb587/go-schemaorg"

var (
	// The person's spouse.
	Spouse = schemaorg.NewProperty("spouse")

	// A person or organization that supports (sponsors) something through some kind
	// of financial contribution.
	Funder = schemaorg.NewProperty("funder")

	// A colleague of the person.
	Colleagues = schemaorg.NewProperty("colleagues")

	// Date of death.
	DeathDate = schemaorg.NewProperty("deathDate")

	// An Organization (or ProgramMembership) to which this Person or Organization
	// belongs.
	MemberOf = schemaorg.NewProperty("memberOf")

	// The height of the item.
	Height = schemaorg.NewProperty("height")

	// A contact location for a person's place of work.
	WorkLocation = schemaorg.NewProperty("workLocation")

	// The total financial value of the person as calculated by subtracting assets
	// from liabilities.
	NetWorth = schemaorg.NewProperty("netWorth")

	// A child of the person.
	Children = schemaorg.NewProperty("children")

	// The job title of the person (for example, Financial Manager).
	JobTitle = schemaorg.NewProperty("jobTitle")

	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	HasOfferCatalog = schemaorg.NewProperty("hasOfferCatalog")

	// The place where the person died.
	DeathPlace = schemaorg.NewProperty("deathPlace")

	// The <a href="http://www.gs1.org/gln">Global Location Number</a> (GLN,
	// sometimes also referred to as International Location Number or ILN) of the
	// respective organization, person, or place. The GLN is a 13-digit number used
	// to identify parties and physical locations.
	GlobalLocationNumber = schemaorg.NewProperty("globalLocationNumber")

	// The place where the person was born.
	BirthPlace = schemaorg.NewProperty("birthPlace")

	// Gender of the person. While http://schema.org/Male and
	// http://schema.org/Female may be used, text strings are also acceptable for
	// people who do not identify as a binary gender.
	Gender = schemaorg.NewProperty("gender")

	// An organization that the person is an alumni of.
	AlumniOf = schemaorg.NewProperty("alumniOf")

	// A contact location for a person's residence.
	HomeLocation = schemaorg.NewProperty("homeLocation")

	// The Dun &amp; Bradstreet DUNS number for identifying an organization or
	// business person.
	Duns = schemaorg.NewProperty("duns")

	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or
	// the CIF/NIF in Spain.
	TaxID = schemaorg.NewProperty("taxID")

	// An award won by or for this item.
	Award = schemaorg.NewProperty("award")

	// Date of birth.
	BirthDate = schemaorg.NewProperty("birthDate")

	// A pointer to products or services offered by the organization or person.
	MakesOffer = schemaorg.NewProperty("makesOffer")

	// Given name. In the U.S., the first name of a Person. This can be used along
	// with familyName instead of the name property.
	GivenName = schemaorg.NewProperty("givenName")

	// A contact point for a person or organization.
	ContactPoints = schemaorg.NewProperty("contactPoints")

	// Awards won by or for this item.
	Awards = schemaorg.NewProperty("awards")

	// Family name. In the U.S., the last name of an Person. This can be used along
	// with givenName instead of the name property.
	FamilyName = schemaorg.NewProperty("familyName")

	// A pointer to products or services sought by the organization or person
	// (demand).
	Seeks = schemaorg.NewProperty("seeks")

	// A sibling of the person.
	Sibling = schemaorg.NewProperty("sibling")

	// Physical address of the item.
	Address = schemaorg.NewProperty("address")

	// Event that this person is a performer or participant in.
	PerformerIn = schemaorg.NewProperty("performerIn")

	// An honorific prefix preceding a Person's name such as Dr/Mrs/Mr.
	HonorificPrefix = schemaorg.NewProperty("honorificPrefix")

	// An additional name for a Person, can be used for a middle name.
	AdditionalName = schemaorg.NewProperty("additionalName")

	// A sibling of the person.
	Siblings = schemaorg.NewProperty("siblings")

	// The telephone number.
	Telephone = schemaorg.NewProperty("telephone")

	// Email address.
	Email = schemaorg.NewProperty("email")

	// The weight of the product or person.
	Weight = schemaorg.NewProperty("weight")

	// A contact point for a person or organization.
	ContactPoint = schemaorg.NewProperty("contactPoint")

	// A colleague of the person.
	Colleague = schemaorg.NewProperty("colleague")

	// A parents of the person.
	Parents = schemaorg.NewProperty("parents")

	// The North American Industry Classification System (NAICS) code for a
	// particular organization or business person.
	Naics = schemaorg.NewProperty("naics")

	// Points-of-Sales operated by the organization or person.
	HasPOS = schemaorg.NewProperty("hasPOS")

	// A parent of this person.
	Parent = schemaorg.NewProperty("parent")

	// Products owned by the organization or person.
	Owns = schemaorg.NewProperty("owns")

	// An organization that this person is affiliated with. For example, a
	// school/university, a club, or a team.
	Affiliation = schemaorg.NewProperty("affiliation")

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

	// The brand(s) associated with a product or service, or the brand(s) maintained
	// by an organization or business person.
	Brand = schemaorg.NewProperty("brand")

	// An honorific suffix preceding a Person's name such as M.D. /PhD/MSCSW.
	HonorificSuffix = schemaorg.NewProperty("honorificSuffix")

	// The Value-added Tax ID of the organization or person.
	VatID = schemaorg.NewProperty("vatID")

	// Nationality of the person.
	Nationality = schemaorg.NewProperty("nationality")

	// The fax number.
	FaxNumber = schemaorg.NewProperty("faxNumber")

	// The most generic familial relation.
	RelatedTo = schemaorg.NewProperty("relatedTo")

	// The most generic uni-directional social relation.
	Follows = schemaorg.NewProperty("follows")

	// The most generic bi-directional social/work relation.
	Knows = schemaorg.NewProperty("knows")

	// Organizations that the person works for.
	WorksFor = schemaorg.NewProperty("worksFor")
)
