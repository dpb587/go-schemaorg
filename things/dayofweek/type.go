package dayofweek

import "github.com/dpb587/go-schemaorg"

// // The day of the week, e.g. used to specify to which day the opening hours of
// an OpeningHoursSpecification refer.</p>
// 
// <p>Originally, URLs from <a
// href="http://purl.org/goodrelations/v1">GoodRelations</a> were used (for <a
// class="localLink" href="http://schema.org/Monday">Monday</a>, <a
// class="localLink" href="http://schema.org/Tuesday">Tuesday</a>, <a
// class="localLink" href="http://schema.org/Wednesday">Wednesday</a>, <a
// class="localLink" href="http://schema.org/Thursday">Thursday</a>, <a
// class="localLink" href="http://schema.org/Friday">Friday</a>, <a
// class="localLink" href="http://schema.org/Saturday">Saturday</a>, <a
// class="localLink" href="http://schema.org/Sunday">Sunday</a> plus a special
// entry for <a class="localLink"
// href="http://schema.org/PublicHolidays">PublicHolidays</a>); these have now
// been integrated directly into schema.org.
var Type = schemaorg.NewDataType("http://schema.org", "DayOfWeek")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}