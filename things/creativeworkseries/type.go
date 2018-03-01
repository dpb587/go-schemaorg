package creativeworkseries

import "github.com/dpb587/go-schemaorg"

// // A CreativeWorkSeries in schema.org is a group of related items, typically but
// not necessarily of the same kind. CreativeWorkSeries are usually organized
// into some order, often chronological. Unlike <a class="localLink"
// href="http://schema.org/ItemList">ItemList</a> which is a general purpose
// data structure for lists of things, the emphasis with CreativeWorkSeries is
// on published materials (written e.g. books and periodicals, or media such as
// tv, radio and games).</p>
// 
// <p>Specific subtypes are available for describing <a class="localLink"
// href="http://schema.org/TVSeries">TVSeries</a>, <a class="localLink"
// href="http://schema.org/RadioSeries">RadioSeries</a>, <a class="localLink"
// href="http://schema.org/MovieSeries">MovieSeries</a>, <a class="localLink"
// href="http://schema.org/BookSeries">BookSeries</a>, <a class="localLink"
// href="http://schema.org/Periodical">Periodical</a> and <a class="localLink"
// href="http://schema.org/VideoGameSeries">VideoGameSeries</a>. In each case,
// the <a class="localLink" href="http://schema.org/hasPart">hasPart</a> / <a
// class="localLink" href="http://schema.org/isPartOf">isPartOf</a> properties
// can be used to relate the CreativeWorkSeries to its parts. The general
// CreativeWorkSeries type serves largely just to organize these more specific
// and practical subtypes.</p>
// 
// <p>It is common for properties applicable to an item from the series to be
// usefully applied to the containing group. Schema.org attempts to anticipate
// some of these cases, but publishers should be free to apply properties of the
// series parts to the series as a whole wherever they seem appropriate.
var Type = schemaorg.NewDataType("http://schema.org", "CreativeWorkSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
