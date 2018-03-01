package mediaobject

import "github.com/dpb587/go-schemaorg"

var (
	// Date when this media object was uploaded to this site.
	UploadDate = schemaorg.NewProperty("uploadDate")

	// Player type required&#x2014;for example, Flash or Silverlight.
	PlayerType = schemaorg.NewProperty("playerType")

	// The height of the item.
	Height = schemaorg.NewProperty("height")

	// The bitrate of the media object.
	Bitrate = schemaorg.NewProperty("bitrate")

	// Indicates if use of the media require a subscription  (either paid or free).
	// Allowed values are <code>true</code> or <code>false</code> (note that an
	// earlier version had 'yes', 'no').
	RequiresSubscription = schemaorg.NewProperty("requiresSubscription")

	// The regions where the media is allowed. If not specified, then it's assumed
	// to be allowed everywhere. Specify the countries in <a
	// href="http://en.wikipedia.org/wiki/ISO_3166">ISO 3166 format</a>.
	RegionsAllowed = schemaorg.NewProperty("regionsAllowed")

	// File size in (mega/kilo) bytes.
	ContentSize = schemaorg.NewProperty("contentSize")

	// A URL pointing to a player for a specific video. In general, this is the
	// information in the <code>src</code> element of an <code>embed</code> tag and
	// should not be the same as the content of the <code>loc</code> tag.
	EmbedUrl = schemaorg.NewProperty("embedUrl")

	// The width of the item.
	Width = schemaorg.NewProperty("width")

	// Actual bytes of the media object, for example the image file or video file.
	ContentUrl = schemaorg.NewProperty("contentUrl")

	// A NewsArticle associated with the Media Object.
	AssociatedArticle = schemaorg.NewProperty("associatedArticle")

	// The production company or studio responsible for the item e.g. series, video
	// game, episode etc.
	ProductionCompany = schemaorg.NewProperty("productionCompany")

	// The duration of the item (movie, audio recording, event, etc.) in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>.
	Duration = schemaorg.NewProperty("duration")

	// The CreativeWork encoded by this media object.
	EncodesCreativeWork = schemaorg.NewProperty("encodesCreativeWork")

	// mp3, mpeg4, etc.
	EncodingFormat = schemaorg.NewProperty("encodingFormat")
)
