package programmembership

import "github.com/dpb587/go-schemaorg"

var (
	// A unique identifier for the membership.
	MembershipNumber = schemaorg.NewProperty("membershipNumber")

	// A member of this organization.
	Members = schemaorg.NewProperty("members")

	// A member of an Organization or a ProgramMembership. Organizations can be
	// members of organizations; ProgramMembership is typically for individuals.
	Member = schemaorg.NewProperty("member")

	// The organization (airline, travelers' club, etc.) the membership is made
	// with.
	HostingOrganization = schemaorg.NewProperty("hostingOrganization")

	// The program providing the membership.
	ProgramName = schemaorg.NewProperty("programName")
)
