package videoobject

import "github.com/dpb587/go-schemaorg"

var (
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The frame size of the video.
	VideoFrameSize = schemaorg.NewProperty("videoFrameSize")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// The quality of the video.
	VideoQuality = schemaorg.NewProperty("videoQuality")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// If this MediaObject is an AudioObject or VideoObject, the transcript of that
	// object.
	Transcript = schemaorg.NewProperty("transcript")

	// Thumbnail image for an image or video.
	Thumbnail = schemaorg.NewProperty("thumbnail")

	// The caption for this object.
	Caption = schemaorg.NewProperty("caption")

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")
)
