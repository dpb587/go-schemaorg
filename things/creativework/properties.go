package creativework

import "github.com/dpb587/go-schemaorg"

var (
	// The subject matter of the content.
	About = schemaorg.NewProperty("about")

	// An alignment to an established educational framework.
	EducationalAlignment = schemaorg.NewProperty("educationalAlignment")

	// A media object that encodes this CreativeWork. This property is a synonym for
	// encoding.
	AssociatedMedia = schemaorg.NewProperty("associatedMedia")

	// A person or organization that supports (sponsors) something through some kind
	// of financial contribution.
	Funder = schemaorg.NewProperty("funder")

	// The position of an item in a series or sequence of items.
	Position = schemaorg.NewProperty("position")

	// An embedded audio object.
	Audio = schemaorg.NewProperty("audio")

	// Example/instance/realization/derivation of the concept of this creative work.
	// eg. The paperback edition, first edition, or eBook.
	WorkExample = schemaorg.NewProperty("workExample")

	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// A media object that encodes this CreativeWork. This property is a synonym for
	// associatedMedia.
	Encoding = schemaorg.NewProperty("encoding")

	// The predominant mode of learning supported by the learning resource.
	// Acceptable values are 'active', 'expositive', or 'mixed'.
	InteractivityType = schemaorg.NewProperty("interactivityType")

	// A human-readable summary of specific accessibility features or deficiencies,
	// consistent with the other accessibility metadata but expressing subtleties
	// such as "short descriptions are present but long descriptions will be needed
	// for non-visual users" or "short descriptions are present and no long
	// descriptions are needed."
	AccessibilitySummary = schemaorg.NewProperty("accessibilitySummary")

	// Fictional person connected with a creative work.
	Character = schemaorg.NewProperty("character")

	// An intended audience, i.e. a group for whom something was created.
	Audience = schemaorg.NewProperty("audience")

	// The Organization on whose behalf the creator was working.
	SourceOrganization = schemaorg.NewProperty("sourceOrganization")

	// Indicates a CreativeWork that this CreativeWork is (in some sense) part of.
	IsPartOf = schemaorg.NewProperty("isPartOf")

	// An embedded video object.
	Video = schemaorg.NewProperty("video")

	// The publisher of the creative work.
	Publisher = schemaorg.NewProperty("publisher")

	// A publication event associated with the item.
	Publication = schemaorg.NewProperty("publication")

	// The textual content of this CreativeWork.
	Text = schemaorg.NewProperty("text")

	// Date the content expires and is no longer useful or available. For example a
	// <a class="localLink" href="http://schema.org/VideoObject">VideoObject</a> or
	// <a class="localLink" href="http://schema.org/NewsArticle">NewsArticle</a>
	// whose availability or relevance is time-limited, or a <a class="localLink"
	// href="http://schema.org/ClaimReview">ClaimReview</a> fact check whose
	// publisher wants to indicate that it may no longer be relevant (or helpful to
	// highlight) after some date.
	Expires = schemaorg.NewProperty("expires")

	// A secondary contributor to the CreativeWork or Event.
	Contributor = schemaorg.NewProperty("contributor")

	// Review of the item.
	Reviews = schemaorg.NewProperty("reviews")

	// The typical expected age range, e.g. '7-9', '11-'.
	TypicalAgeRange = schemaorg.NewProperty("typicalAgeRange")

	// The place and time the release was issued, expressed as a PublicationEvent.
	ReleasedEvent = schemaorg.NewProperty("releasedEvent")

	// The purpose of a work in the context of education; for example, 'assignment',
	// 'group work'.
	EducationalUse = schemaorg.NewProperty("educationalUse")

	// The location depicted or described in the content. For example, the location
	// in a photograph or painting.
	ContentLocation = schemaorg.NewProperty("contentLocation")

	// Indicates (by URL or string) a particular version of a schema used in some
	// CreativeWork. For example, a document could declare a schemaVersion using an
	// URL such as http://schema.org/version/2.0/ if precise indication of schema
	// version was required by some application.
	SchemaVersion = schemaorg.NewProperty("schemaVersion")

	// Content features of the resource, such as accessible media, alternatives and
	// supported enhancements for accessibility (<a
	// href="http://www.w3.org/wiki/WebSchemas/Accessibility">WebSchemas wiki lists
	// possible values</a>).
	AccessibilityFeature = schemaorg.NewProperty("accessibilityFeature")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// A secondary title of the CreativeWork.
	AlternativeHeadline = schemaorg.NewProperty("alternativeHeadline")

	// The location where the CreativeWork was created, which may not be the same as
	// the location depicted in the CreativeWork.
	LocationCreated = schemaorg.NewProperty("locationCreated")

	// A list of single or combined accessModes that are sufficient to understand
	// all the intellectual content of a resource. Expected values include:
	// auditory, tactile, textual, visual.
	AccessModeSufficient = schemaorg.NewProperty("accessModeSufficient")

	// The temporalCoverage of a CreativeWork indicates the period that the content
	// applies to, i.e. that it describes, either as a DateTime or as a textual
	// string indicating a time period in <a
	// href="https://en.wikipedia.org/wiki/ISO_8601#Time_intervals">ISO 8601 time
	// interval format</a>. In
	//       the case of a Dataset it will typically indicate the relevant time
	// period in a precise notation (e.g. for a 2011 census dataset, the year 2011
	// would be written "2011/2012"). Other forms of content e.g. ScholarlyArticle,
	// Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader
	// terms - textually or via well-known URL.
	//       Written works such as books may sometimes have precise temporal
	// coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601
	// interval format format via "1939/1945".
	TemporalCoverage = schemaorg.NewProperty("temporalCoverage")

	// Specifies the Person that is legally accountable for the CreativeWork.
	AccountablePerson = schemaorg.NewProperty("accountablePerson")

	// The spatialCoverage of a CreativeWork indicates the place(s) which are the
	// focus of the content. It is a subproperty of
	//       contentLocation intended primarily for more technical and detailed
	// materials. For example with a Dataset, it indicates
	//       areas that the dataset describes: a dataset of New York weather would
	// have spatialCoverage which was the place: the state of New York.
	SpatialCoverage = schemaorg.NewProperty("spatialCoverage")

	// An offer to provide this item&#x2014;for example, an offer to sell a product,
	// rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers = schemaorg.NewProperty("offers")

	// Specifies the Person who edited the CreativeWork.
	Editor = schemaorg.NewProperty("editor")

	// A link to the page containing the comments of the CreativeWork.
	DiscussionUrl = schemaorg.NewProperty("discussionUrl")

	// An award won by or for this item.
	Award = schemaorg.NewProperty("award")

	// The party holding the legal copyright to the CreativeWork.
	CopyrightHolder = schemaorg.NewProperty("copyrightHolder")

	// A characteristic of the described resource that is physiologically dangerous
	// to some users. Related to WCAG 2.0 guideline 2.3 (<a
	// href="http://www.w3.org/wiki/WebSchemas/Accessibility">WebSchemas wiki lists
	// possible values</a>).
	AccessibilityHazard = schemaorg.NewProperty("accessibilityHazard")

	// The year during which the claimed copyright for the CreativeWork was first
	// asserted.
	CopyrightYear = schemaorg.NewProperty("copyrightYear")

	// Awards won by or for this item.
	Awards = schemaorg.NewProperty("awards")

	// The Event where the CreativeWork was recorded. The CreativeWork may capture
	// all or part of the event.
	RecordedAt = schemaorg.NewProperty("recordedAt")

	// The number of comments this CreativeWork (e.g. Article, Question or Answer)
	// has received. This is most applicable to works published in Web sites with
	// commenting system; additional comments may exist elsewhere.
	CommentCount = schemaorg.NewProperty("commentCount")

	// Media type, typically MIME format (see <a
	// href="http://www.iana.org/assignments/media-types/media-types.xhtml">IANA
	// site</a>) of the content e.g. application/zip of a SoftwareApplication
	// binary. In cases where a CreativeWork has several media type representations,
	// 'encoding' can be used to indicate each MediaObject alongside particular
	// fileFormat information. Unregistered or niche file formats can be indicated
	// instead via the most appropriate URL, e.g. defining Web page or a Wikipedia
	// entry.
	FileFormat = schemaorg.NewProperty("fileFormat")

	// The language of the content or performance or used in an action. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink"
	// href="http://schema.org/availableLanguage">availableLanguage</a>.
	InLanguage = schemaorg.NewProperty("inLanguage")

	// Indicates that the resource is compatible with the referenced accessibility
	// API (<a href="http://www.w3.org/wiki/WebSchemas/Accessibility">WebSchemas
	// wiki lists possible values</a>).
	AccessibilityAPI = schemaorg.NewProperty("accessibilityAPI")

	// The number of interactions for the CreativeWork using the WebSite or
	// SoftwareApplication. The most specific child type of InteractionCounter
	// should be used.
	InteractionStatistic = schemaorg.NewProperty("interactionStatistic")

	// Official rating of a piece of content&#x2014;for example,'MPAA PG-13'.
	ContentRating = schemaorg.NewProperty("contentRating")

	// The predominant type or kind characterizing the learning resource. For
	// example, 'presentation', 'handout'.
	LearningResourceType = schemaorg.NewProperty("learningResourceType")

	// The human sensory perceptual system or cognitive faculty through which a
	// person may process or perceive information. Expected values include:
	// auditory, tactile, textual, visual, colorDependent, chartOnVisual,
	// chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
	AccessMode = schemaorg.NewProperty("accessMode")

	// A material that something is made from, e.g. leather, wool, cotton, paper.
	Material = schemaorg.NewProperty("material")

	// Indicates whether this content is family friendly.
	IsFamilyFriendly = schemaorg.NewProperty("isFamilyFriendly")

	// A creative work that this work is an example/instance/realization/derivation
	// of.
	ExampleOfWork = schemaorg.NewProperty("exampleOfWork")

	// The version of the CreativeWork embodied by a specified resource.
	Version = schemaorg.NewProperty("version")

	// The date on which the CreativeWork was most recently modified or when the
	// item's entry was modified within a DataFeed.
	DateModified = schemaorg.NewProperty("dateModified")

	// Indicates the primary entity described in some page or other CreativeWork.
	MainEntity = schemaorg.NewProperty("mainEntity")

	// Genre of the creative work, broadcast channel or group.
	Genre = schemaorg.NewProperty("genre")

	// Keywords or tags used to describe this content. Multiple entries in a
	// keywords list are typically delimited by commas.
	Keywords = schemaorg.NewProperty("keywords")

	// The author of this content or rating. Please note that author is special in
	// that HTML 5 provides a special mechanism for indicating authorship via the
	// rel tag. That is equivalent to this and may be used interchangeably.
	Author = schemaorg.NewProperty("author")

	// A resource that was used in the creation of this resource. This term can be
	// repeated for multiple sources. For example,
	// http://example.com/great-multiplication-intro.html.
	IsBasedOnUrl = schemaorg.NewProperty("isBasedOnUrl")

	// Approximate or typical time it takes to work with or through this learning
	// resource for the typical intended target audience, e.g. 'P30M', 'P1H25M'.
	TimeRequired = schemaorg.NewProperty("timeRequired")

	// Organization or person who adapts a creative work to different languages,
	// regional differences and technical requirements of a target market, or that
	// translates during some event.
	Translator = schemaorg.NewProperty("translator")

	// A thumbnail image relevant to the Thing.
	ThumbnailUrl = schemaorg.NewProperty("thumbnailUrl")

	// Indicates a CreativeWork that is (in some sense) a part of this CreativeWork.
	HasPart = schemaorg.NewProperty("hasPart")

	// Comments, typically from users.
	Comment = schemaorg.NewProperty("comment")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// A license document that applies to this content, typically indicated by URL.
	License = schemaorg.NewProperty("license")

	// Identifies input methods that are sufficient to fully control the described
	// resource (<a
	// href="http://www.w3.org/wiki/WebSchemas/Accessibility">WebSchemas wiki lists
	// possible values</a>).
	AccessibilityControl = schemaorg.NewProperty("accessibilityControl")

	// A media object that encodes this CreativeWork.
	Encodings = schemaorg.NewProperty("encodings")

	// A resource that was used in the creation of this resource. This term can be
	// repeated for multiple sources. For example,
	// http://example.com/great-multiplication-intro.html.
	IsBasedOn = schemaorg.NewProperty("isBasedOn")

	// The creator/author of this CreativeWork. This is the same as the Author
	// property for CreativeWork.
	Creator = schemaorg.NewProperty("creator")

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

	// The person or organization who produced the work (e.g. music album, movie,
	// tv/radio series etc.).
	Producer = schemaorg.NewProperty("producer")

	// Indicates that the CreativeWork contains a reference to, but is not
	// necessarily about a concept.
	Mentions = schemaorg.NewProperty("mentions")

	// The date on which the CreativeWork was created or the item was added to a
	// DataFeed.
	DateCreated = schemaorg.NewProperty("dateCreated")

	// Date of first broadcast/publication.
	DatePublished = schemaorg.NewProperty("datePublished")

	// A flag to signal that the item, event, or place is accessible for free.
	IsAccessibleForFree = schemaorg.NewProperty("isAccessibleForFree")

	// Headline of the article.
	Headline = schemaorg.NewProperty("headline")

	// A citation or reference to another creative work, such as another
	// publication, web page, scholarly article, etc.
	Citation = schemaorg.NewProperty("citation")
)
