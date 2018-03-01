package spreadsheetdigitaldocument

import "github.com/dpb587/go-schemaorg"

// // A spreadsheet file.
var Type = schemaorg.NewDataType("http://schema.org", "SpreadsheetDigitalDocument")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
