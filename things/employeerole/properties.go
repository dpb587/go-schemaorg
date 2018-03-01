package employeerole

import "github.com/dpb587/go-schemaorg"

var (
	// The currency (coded using <a href="http://en.wikipedia.org/wiki/ISO_4217">ISO
	// 4217</a> ) used for the main salary information in this job posting or for
	// this employee.
	SalaryCurrency = schemaorg.NewProperty("salaryCurrency")

	// The base salary of the job or of an employee in an EmployeeRole.
	BaseSalary = schemaorg.NewProperty("baseSalary")
)
