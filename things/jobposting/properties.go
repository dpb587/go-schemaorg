package jobposting

import "github.com/dpb587/go-schemaorg"

var (
	// Description of benefits associated with the job.
	Benefits = schemaorg.NewProperty("benefits")

	// Specific qualifications required for this role.
	Qualifications = schemaorg.NewProperty("qualifications")

	// Description of bonus and commission compensation aspects of the job.
	IncentiveCompensation = schemaorg.NewProperty("incentiveCompensation")

	// The typical working hours for this job (e.g. 1st shift, night shift,
	// 8am-5pm).
	WorkHours = schemaorg.NewProperty("workHours")

	// The currency (coded using <a href="http://en.wikipedia.org/wiki/ISO_4217">ISO
	// 4217</a> ) used for the main salary information in this job posting or for
	// this employee.
	SalaryCurrency = schemaorg.NewProperty("salaryCurrency")

	// Description of benefits associated with the job.
	JobBenefits = schemaorg.NewProperty("jobBenefits")

	// Publication date for the job posting.
	DatePosted = schemaorg.NewProperty("datePosted")

	// Skills required to fulfill this role.
	Skills = schemaorg.NewProperty("skills")

	// Description of bonus and commission compensation aspects of the job.
	Incentives = schemaorg.NewProperty("incentives")

	// Educational background needed for the position.
	EducationRequirements = schemaorg.NewProperty("educationRequirements")

	// Responsibilities associated with this role.
	Responsibilities = schemaorg.NewProperty("responsibilities")

	// The base salary of the job or of an employee in an EmployeeRole.
	BaseSalary = schemaorg.NewProperty("baseSalary")

	// The date after when the item is not valid. For example the end of an offer,
	// salary period, or a period of opening hours.
	ValidThrough = schemaorg.NewProperty("validThrough")

	// Organization offering the job position.
	HiringOrganization = schemaorg.NewProperty("hiringOrganization")

	// Any special commitments associated with this job posting. Valid entries
	// include VeteranCommit, MilitarySpouseCommit, etc.
	SpecialCommitments = schemaorg.NewProperty("specialCommitments")

	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy:
	// http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and
	// formal code, with the property repeated for each applicable value.
	OccupationalCategory = schemaorg.NewProperty("occupationalCategory")

	// Description of skills and experience needed for the position.
	ExperienceRequirements = schemaorg.NewProperty("experienceRequirements")

	// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal,
	// internship).
	EmploymentType = schemaorg.NewProperty("employmentType")

	// A (typically single) geographic location associated with the job position.
	JobLocation = schemaorg.NewProperty("jobLocation")

	// The title of the job.
	Title = schemaorg.NewProperty("title")

	// The industry associated with the job position.
	Industry = schemaorg.NewProperty("industry")
)
