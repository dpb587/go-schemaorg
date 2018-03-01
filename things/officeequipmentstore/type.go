package officeequipmentstore

import "github.com/dpb587/go-schemaorg"

// // An office equipment store.
var Type = schemaorg.NewDataType("http://schema.org", "OfficeEquipmentStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
