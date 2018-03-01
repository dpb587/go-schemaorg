package course

import "github.com/dpb587/go-schemaorg"

var (
	// Requirements for taking the Course. May be completion of another <a
	// class="localLink" href="http://schema.org/Course">Course</a> or a textual
	// description like "permission of instructor". Requirements may be a
	// pre-requisite competency, referenced using <a class="localLink"
	// href="http://schema.org/AlignmentObject">AlignmentObject</a>.
	CoursePrerequisites = schemaorg.NewProperty("coursePrerequisites")

	// The identifier for the <a class="localLink"
	// href="http://schema.org/Course">Course</a> used by the course <a
	// class="localLink" href="http://schema.org/provider">provider</a> (e.g. CS101
	// or 6.001).
	CourseCode = schemaorg.NewProperty("courseCode")

	// An offering of the course at a specific time and place or through specific
	// media or mode of study or to a specific section of students.
	HasCourseInstance = schemaorg.NewProperty("hasCourseInstance")
)
