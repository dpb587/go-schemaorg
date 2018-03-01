package datadownload

import "github.com/dpb587/go-schemaorg"

// // A dataset in downloadable form.
var Type = schemaorg.NewDataType("http://schema.org", "DataDownload")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
