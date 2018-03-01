package courseinstance

import "github.com/dpb587/go-schemaorg"

var (
	// A person assigned to instruct or provide instructional assistance for the <a
	// class="localLink" href="http://schema.org/CourseInstance">CourseInstance</a>.
	Instructor = schemaorg.NewProperty("instructor")

	// The medium or means of delivery of the course instance or the mode of study,
	// either as a text label (e.g. "online", "onsite" or "blended"; "synchronous"
	// or "asynchronous"; "full-time" or "part-time") or as a URL reference to a
	// term from a controlled vocabulary (e.g.
	// https://ceds.ed.gov/element/001311#Asynchronous ).
	CourseMode = schemaorg.NewProperty("courseMode")
)
