package mediaobject

import "github.com/dpb587/go-schemaorg"

// // A media object, such as an image, video, or audio object embedded in a web
// page or a downloadable dataset i.e. DataDownload. Note that a creative work
// may have many media objects associated with it on the same web page. For
// example, a page about a single song (MusicRecording) may have a music video
// (VideoObject), and a high and low bandwidth audio stream (2 AudioObject's).
var Type = schemaorg.NewDataType("http://schema.org", "MediaObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
