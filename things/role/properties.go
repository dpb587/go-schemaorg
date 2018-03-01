package role

import "github.com/dpb587/go-schemaorg"

var (
	// The end date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	EndDate = schemaorg.NewProperty("endDate")

	// A position played, performed or filled by a person or organization, as part
	// of an organization. For example, an athlete in a SportsTeam might play in the
	// position named 'Quarterback'.
	NamedPosition = schemaorg.NewProperty("namedPosition")

	// The start date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	StartDate = schemaorg.NewProperty("startDate")

	// A role played, performed or filled by a person or organization. For example,
	// the team of creators for a comic book might fill the roles named 'inker',
	// 'penciller', and 'letterer'; or an athlete in a SportsTeam might play in the
	// position named 'Quarterback'.
	RoleName = schemaorg.NewProperty("roleName")
)
