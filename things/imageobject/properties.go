package imageobject

import "github.com/dpb587/go-schemaorg"

var (
	// exif data for this object.
	ExifData = schemaorg.NewProperty("exifData")

	// Indicates whether this image is representative of the content of the page.
	RepresentativeOfPage = schemaorg.NewProperty("representativeOfPage")

	// Thumbnail image for an image or video.
	Thumbnail = schemaorg.NewProperty("thumbnail")

	// The caption for this object.
	Caption = schemaorg.NewProperty("caption")
)
